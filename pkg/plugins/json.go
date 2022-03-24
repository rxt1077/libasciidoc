package plugins

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
  "github.com/davecgh/go-spew/spew"
  "github.com/pkg/errors"

  log "github.com/sirupsen/logrus"

  "encoding/json"
  "fmt"
  "reflect"
)

// basic struct for storing a value and a type
// uses json field tags to make it easier to look at
type Wrap struct {
  Type string `json:"type"`
  Value interface{} `json:"value"`
}

// recursively walks through all elements, wrapping them in an ElementWithType
func wrap(element interface{}) Wrap {
  log.Tracef("=== Wrapping ===\n%s", spew.Sdump(element))

  reflectType := reflect.TypeOf(element)
  name := reflectType.String()
  value := reflect.ValueOf(element)
  kind := value.Type().Kind()
  log.Tracef("Type: %s Kind: %s\n", name, kind.String())

  // derefence pointers that aren't nil
  if kind == reflect.Ptr {
    if !value.IsNil() {
      log.Trace("Dereferencing pointer")
      value = value.Elem()
      kind = value.Type().Kind()
    }
  }

  wrapped := Wrap{
    Type: name,
    Value: nil,
  }
  switch kind {
  // if this is a slice, walk though each index and return a new slice
  // wrapped in an ElementWithType
  case reflect.Slice:
    log.Trace("Element is a slice")
    var newSlice []interface{}
    for i := 0; i < value.Len(); i++ {
      log.Tracef("Wrapping index %d\n", i)
      newSlice = append(newSlice, wrap(value.Index(i).Interface()))
    }
    wrapped.Value = newSlice
  // if this is a struct, walk through each field and return a
  // map[string]interface{} wrapped in an ElementWithType
  case reflect.Struct:
    log.Trace("Element is a struct")
    newMap := map[string]interface{}{}
    for i := 0; i < value.NumField(); i++ {
      name := value.Type().Field(i).Name
      log.Tracef("Wrapping field %s\n", name)
      newMap[name] = wrap(value.Field(i).Interface())
    }
    wrapped.Value = newMap
  // if this is a map, go through all keys and wrap each value in ElementWithType
  // this assumes that keys are strings
  case reflect.Map:
    log.Trace("Element is a map")
    if ! value.IsNil() { // this prevents nils from ending up as empty maps
      newMap := map[string]interface{}{}
      for _, e := range value.MapKeys() {
        v := value.MapIndex(e)
        log.Tracef("Wrapping value at key %s\n", e)
        newMap[e.String()] = wrap(v.Interface())
      }
      wrapped.Value = newMap
    }
  // if this is anything else, we assume its atomic. Return it wrapped in an
  // ElementWithType
  default:
    log.Trace("Element is an atom")
    wrapped.Value = element
  }

  log.Tracef("%s=== Returning ===", spew.Sdump(wrapped))
  return wrapped
}

// recursively walks through a map, unwrapping each value and creating its
// corresponding libasciidoc type
func unwrap(src interface{}) (interface{}, error) {
  log.Tracef("=== Unwrapping ===\n%s", spew.Sdump(src))

  // take off the outer wrapping of the element
  elem, ok := src.(map[string]interface{})
  if ! ok {
    return nil, errors.New("src is not a map[string]interface{}")
  }
  t, ok := elem["type"]
  if ! ok {
    return nil, errors.New("'type' not found in map")
  }
  ourType, ok := t.(string)
  if ! ok {
    return nil, errors.New("'type' is not a string")
  }
  e, ok := elem["value"]
  if ! ok {
    return nil, errors.New("'value' not found in map")
  }

  log.Tracef("Type: %s\n", ourType)

  // if it's nil, return immediately. It won't be set in the struct.
  if e == nil {
    log.Tracef("%s=== Returning ===\n", spew.Sdump(nil))
    return nil, nil
  }

  // create an object for the element to be put into
  // or return immediately if it's an atom
  var obj interface{}
  switch ourType {
  case "string": // atom
    log.Tracef("%s=== Returning ===\n", spew.Sdump(e))
    return e, nil
  case "int": // atom
    // JSON decodes all numbers to float64 so we must cast it
    floatNum, ok := e.(float64)
    if ! ok {
      return nil, errors.New("'int' type is not actually float64)")
    }
    intNum := int(floatNum)
    log.Tracef("%s=== Returning ===\n", spew.Sdump(intNum))
    return intNum, nil
  case "types.UnorderedListElementBulletStyle": // atom
    // this is a string at its base and JSON will store it as such
    str, ok := e.(string)
    if ! ok {
      return nil, errors.New("'types.UnorderedListElementBulletStyle' type is not actually string)")
    }
    ulebs := types.UnorderedListElementBulletStyle(str)
    log.Tracef("%s=== Returning ===\n", spew.Sdump(ulebs))
    return ulebs, nil
  case "types.UnorderedListElementCheckStyle": // atom
    // this is a string at its base and JSON will store it as such
    str, ok := e.(string)
    if ! ok {
      return nil, errors.New("'types.UnorderedListElementCheckStyle' type is not actually string)")
    }
    ulecs := types.UnorderedListElementCheckStyle(str)
    log.Tracef("%s=== Returning ===\n", spew.Sdump(ulecs))
    return ulecs, nil
  case "[]interface {}":
    obj = []interface{}{}
  case "[]*types.Footnote":
    obj = []*types.Footnote{}
  case "[]types.ListElement":
    obj = []types.ListElement{}
  case "[]*types.ToCSection":
    obj = []*types.ToCSection{}
  case "types.Attributes":
    obj = types.Attributes{}
  case "*types.Document":
    obj = &types.Document{}
  case "*types.DocumentHeader":
    obj = &types.DocumentHeader{}
  case "types.ElementReferences":
    obj = types.ElementReferences{}
  case "*types.InlineImage":
    obj = &types.InlineImage{}
  case "*types.List":
    obj = &types.List{}
  case "*types.Location":
    obj = &types.Location{}
  case "*types.Paragraph":
    obj = &types.Paragraph{}
  case "*types.Preamble":
    obj = &types.Preamble{}
  case "*types.Section":
    obj = &types.Section{}
  case "*types.StringElement":
    obj = &types.StringElement{}
  case "*types.TableOfContents":
    obj = &types.TableOfContents{}
  case "*types.ToCSection":
    obj = &types.ToCSection{}
  case "*types.UnorderedListElement":
    obj = &types.UnorderedListElement{}
  default:
    return nil, fmt.Errorf("unknown type: %s", ourType)
  }

  objValue := reflect.ValueOf(obj)
  kind := objValue.Type().Kind()
  log.Tracef("Kind: %s\n", kind.String())

  if kind == reflect.Ptr {
    if !objValue.IsNil() {
      log.Trace("Dereferencing pointer")
      objValue = objValue.Elem()
      kind = objValue.Type().Kind()
      log.Tracef("Kind: %s\n", kind.String())
    }
  }

  // use reflection to fill the object
  switch kind {
  case reflect.Struct:
    // structs are represented as map[string]interface{} in JSON
    assertedElement, ok := e.(map[string]interface{})
    if ! ok {
      return nil, errors.New("unwrapping struct but src element is not map[string]interface{}")
    }
    // recursively go through all fields of the struct and unwrap them
    for k, v := range assertedElement {
      // unwrap the fields in the map
      log.Tracef("unwrapping field: %s\n", k)
      result, err := unwrap(v)
      if err != nil {
        return nil, err
      }
      // set the field in the destination struct if it's not nil
      if result != nil {
        structFieldValue := objValue.FieldByName(k)
        if !structFieldValue.IsValid() {
            return nil, fmt.Errorf("no such field: %s", k)
        }
        if !structFieldValue.CanSet() {
            return nil, fmt.Errorf("can't set %s field value", k)
        }
        // make sure the types match or the destination is an interface{}
        if (structFieldValue.Type() != reflect.TypeOf(result)) &&
           (structFieldValue.Type() != reflect.TypeOf(new(interface{})).Elem()) {
            return nil, fmt.Errorf("value type %s didn't match field type %s", reflect.TypeOf(result).String(), structFieldValue.Type().String())
        }
        structFieldValue.Set(reflect.ValueOf(result))
      }
    }
    log.Tracef("%s=== Returning ===\n", spew.Sdump(obj))
    return obj, nil
  case reflect.Map:
    // maps are represented as map[string]interface{} in JSON
    assertedElement, ok := e.(map[string]interface{})
    if ! ok {
      return nil, errors.New("unwrapping struct but src element is not map[string]interface{}")
    }
    // recursively go through all fields of the map and unwrap them
    for k, v := range assertedElement {
      // unwrap the fields in the map
      log.Tracef("unwrapping value at key: %s\n", k)
      result, err := unwrap(v)
      if err != nil {
        return nil, err
      }

      // set the value in obj
      objValue.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(result))
    }
    log.Tracef("%s=== Returning ===\n", spew.Sdump(obj))
    return obj, nil
  case reflect.Slice:
    // slices are represented as []interface{} in JSON
    assertedElement, ok := e.([]interface{})
    if ! ok {
      return nil, errors.New("unwrapping slice but src element is not []interface{}")
    }
    // recursively go through all indexes and unwrap them
    for i, v := range assertedElement {
      log.Tracef("unwrapping index: %d\n", i)
      result, err := unwrap(v)
      if err != nil {
        return nil, err
      }
      // append the result to the slice
      objValue = reflect.Append(objValue, reflect.ValueOf(result))
    }
    log.Tracef("%s=== Returning ===\n", spew.Sdump(obj))
    return objValue.Interface(), nil
  default:
    return nil, fmt.Errorf("unknown kind: %s", kind.String())
  }

  return nil, nil
}

func MarshalJSON(doc *types.Document) ([]byte, error) {
  return json.Marshal(wrap(doc))
}

func UnmarshalJSON(b []byte) (*types.Document, error) {
  var result map[string]interface{}
  err := json.Unmarshal(b, &result)
  if err != nil {
    return nil, err
  }
  doc, err := unwrap(result)
  if err != nil {
    return nil, err
  }
  return doc.(*types.Document), nil
}
