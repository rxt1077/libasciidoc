package plugins

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
  "github.com/davecgh/go-spew/spew"
  "github.com/pkg/errors"

  "encoding/json"
  "fmt"
  "reflect"
)

type ElementWithType struct {
  Element interface{}
  Type string
}

// recursively walks through all elements, wrapping them in an ElementWithType
// should this return type be ElementWithType?
func wrap(element interface{}) interface{} {
  fmt.Printf("=== Wrapping ===\n%s", spew.Sdump(element))

  reflectType := reflect.TypeOf(element)
  name := reflectType.String()
  value := reflect.ValueOf(element)
  kind := value.Type().Kind()
  fmt.Printf("Type: %s Kind: %s\n", name, kind.String())

  // derefence pointers that aren't nil
  if kind == reflect.Ptr {
    if !value.IsNil() {
      fmt.Println("Dereferencing pointer")
      value = value.Elem()
      kind = value.Type().Kind()
    }
  }

  var elementWithType ElementWithType
  switch kind {
  // if this is a slice, walk though each index and return a new slice
  // wrapped in an ElementWithType
  case reflect.Slice:
    fmt.Println("Element is a slice")
    var newSlice []interface{}
    for i := 0; i < value.Len(); i++ {
      fmt.Printf("Wrapping index %d\n", i)
      newSlice = append(newSlice, wrap(value.Index(i).Interface()))
    }
    elementWithType = ElementWithType{
      Type: name,
      Element: newSlice,
    }
  // if this is a struct, walk through each field and return a
  // map[string]interface{} wrapped in an ElementWithType
  case reflect.Struct:
    fmt.Println("Element is a struct")
    newMap := map[string]interface{}{}
    for i := 0; i < value.NumField(); i++ {
      name := value.Type().Field(i).Name
      fmt.Printf("Wrapping field %s\n", name)
      newMap[name] = wrap(value.Field(i).Interface())
    }
    elementWithType = ElementWithType{
      Type: name,
      Element: newMap,
    }
  // if this is a map, go through all keys and wrap each value in ElementWithType
  // this assumes that keys are strings
  case reflect.Map:
    fmt.Println("Element is a map")
    newMap := map[string]interface{}{}
    for _, e := range value.MapKeys() {
      v := value.MapIndex(e)
      fmt.Printf("Wrapping value at key %s\n", e)
      newMap[e.String()] = wrap(v)
    }
    elementWithType = ElementWithType {
      Type: name,
      Element: newMap,
    }
  // if this is anything else, we assume its atomic. Return it wrapped in an
  // ElementWithType
  default:
    fmt.Println("Element is an atom")
    elementWithType = ElementWithType{
      Type: name,
      Element: element,
    }
  }

  fmt.Printf("=== Returning ===\n%s", spew.Sdump(elementWithType))
  return elementWithType
}

// Source: https://stackoverflow.com/questions/26744873/converting-map-to-struct
// the mapstructure lib seemed like overkill
func SetField(obj interface{}, name string, value interface{}) error {
    fmt.Printf("SetField obj: %s name: %s value: %s\n", spew.Sdump(obj), name, spew.Sdump(value))
    structValue := reflect.ValueOf(obj).Elem()
    structFieldValue := structValue.FieldByName(name)

    if !structFieldValue.IsValid() {
        return fmt.Errorf("No such field: %s in obj", name)
    }

    if !structFieldValue.CanSet() {
        return fmt.Errorf("Cannot set %s field value", name)
    }

    structFieldType := structFieldValue.Type()
    val := reflect.ValueOf(value)
    if structFieldType != val.Type() {
        return errors.New("Provided value type didn't match obj field type")
    }

    structFieldValue.Set(val)
    return nil
}

func FillStruct(m map[string]interface{}, s interface{}) error {
  for k, v := range m {
      err := SetField(s, k, v)
      if err != nil {
          return err
      }
  }
  return nil
}
/*
// structs are map[string]interface{} in the JSON representation
// recursively unwrap each element and then use reflect to set the
// field values of the struct
func unwrapStruct(dstStruct interface{}, obj interface{}) (interface{}, error) {
  srcMap, ok := obj.(map[string]interface{})
  if ! ok {
    return nil, errors.New("unwrapStruct obj is not map[string]interface{}")
  }
  value := reflect.ValueOf(dstStruct)
  for k, v := range srcMap {
    //unwrap the fields in the map
    fmt.Printf("unwrapStruct unwrapping field: %s\n", k)
    result, err := unwrap(v)
    if err != nil {
      return nil, err
    }

    //set the field in the destination struct
    structFieldValue := value.FieldByName(k)
    if !structFieldValue.IsValid() {
        return nil, fmt.Errorf("unwrapStruct no such field: %s", k)
    }
    if !structFieldValue.CanSet() {
        return nil, fmt.Errorf("unwrapStruct can't set %s field value", k)
    }
    if structFieldValue.Type() != reflect.TypeOf(result) {
        return nil, errors.New("unwrapStruct value type didn't match field type")
    }
    structFieldValue.Set(reflect.ValueOf(result))
  }
  return dstStruct, nil
}

func unwrapMap(dst interface{}, obj interface{}) (interface{}, error) {
  srcMap, ok := obj.(map[string]interface{})
  if ! ok {
    return nil, errors.New("unwrapMap obj is not map[string]interface{}")
  }

  dstValue := reflect.ValueOf(dst)
  for k, v := range srcMap {
    result, err := unwrap(v)
    if err != nil {
      return nil, err
    }
    //TODO: Add more checks like in struct above
    dstValue.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(result))
  }
  return dst, nil
}

func unwrapSlice(obj interface{}) ([]interface{}, error) {
  srcSlice, ok := obj.([]interface{})
  if ! ok {
    return nil, errors.New("unwrapSlice obj is not []interface{}")
  }
  dstSlice := []interface{}{}
  for i, v := range srcSlice {
    //unwrap the indexes in the slice
    fmt.Printf("unwrapSlice unwrapping index: %d\n", i)
    result, err := unwrap(v)
    if err != nil {
      return nil, err
    }

    //append the unwrapped result to the new slice
    dstSlice = append(dstSlice, result)
  }
  return dstSlice, nil
}
*/

func unwrap(src interface{}) (interface{}, error) {
  fmt.Printf("=== Unwrapping ===\n%s", spew.Sdump(src))

  // take off the outer wrapping of the element
  elem, ok := src.(map[string]interface{})
  if ! ok {
    return nil, errors.New("plugins.unwrap src is not a map[string]interface{}")
  }
  t, ok := elem["Type"]
  if ! ok {
    return nil, errors.New("plugins.unwrap 'Type' not found in map")
  }
  ourType, ok := t.(string)
  if ! ok {
    return nil, errors.New("plugins.unwrap 'Type' is not a string")
  }
  e, ok := elem["Element"]
  if ! ok {
    return nil, errors.New("plugins.unwrap 'Element' not found in map")
  }

  fmt.Printf("Type: %s\n", ourType)

  // create an object for the element to be put into
  // or the correct nil object if needed
  // I can't seem to get a reflect Convert statement that will build
  // these nils automatically
  var obj interface{}
  var nilObj interface{}
  switch ourType {
  case "[]interface {}":
    obj = []interface{}{}
    nilObj = []interface{}(nil)
  case "[]*types.Footnote":
    obj = []*types.Footnote{}
    nilObj = []*types.Footnote(nil)
  case "string": // atom
    fmt.Printf("%s\n=== Returning ===\n", spew.Sdump(e))
    return e, nil
  case "types.Attributes":
    obj = types.Attributes{}
  case "*types.Document":
    obj = &types.Document{}
  case "*types.DocumentHeader":
    obj = &types.DocumentHeader{}
  case "types.ElementReferences":
    obj = types.ElementReferences{}
  case "*types.StringElement":
    obj = &types.StringElement{}
  case "*types.TableOfContents":
    obj = &types.TableOfContents{}
    nilObj = (*types.TableOfContents)(nil)
  default:
    return nil, fmt.Errorf("plugins.unwrap unknown type: %s", ourType)
  }

  spew.Dump(obj)
  spew.Dump(nilObj)

  if e == nil {
    fmt.Printf("%s\n=== Returning ===\n", spew.Sdump(nilObj))
    return nilObj, nil
  }


  objValue := reflect.ValueOf(obj)
  kind := objValue.Type().Kind()
  fmt.Printf("Kind: %s\n", kind.String())

  if kind == reflect.Ptr {
    if !objValue.IsNil() {
      fmt.Println("Dereferencing pointer")
      objValue = objValue.Elem()
      kind = objValue.Type().Kind()
      fmt.Printf("Kind: %s\n", kind.String())
    }
  }

  // use reflection to fill the object
  switch kind {
  case reflect.Struct:
    // structs are represented as map[string]interface{} in JSON
    assertedElement, ok := e.(map[string]interface{})
    if ! ok {
      return nil, errors.New("plugins.unwrap unwrapping struct but src element is not map[string]interface{}")
    }
    // recursively go through all fields of the struct and unwrap them
    for k, v := range assertedElement {
      // unwrap the fields in the map
      fmt.Printf("plugins.unwrap unwrapping field: %s\n", k)
      result, err := unwrap(v)
      if err != nil {
        return nil, err
      }

      // set the field in the destination struct
      structFieldValue := objValue.FieldByName(k)
      if !structFieldValue.IsValid() {
          return nil, fmt.Errorf("plugins.unwrap no such field: %s", k)
      }
      if !structFieldValue.CanSet() {
          return nil, fmt.Errorf("plugins.unwrap can't set %s field value", k)
      }
      if structFieldValue.Type() != reflect.TypeOf(result) {
          return nil, errors.New("plugins.unwrap value type didn't match field type")
      }
      structFieldValue.Set(reflect.ValueOf(result))
      spew.Dump(obj)
    }
    fmt.Printf("%s\n=== Returning ===\n", spew.Sdump(obj))
    return obj, nil
  case reflect.Map:
    // maps are represented as map[string]interface{} in JSON
    assertedElement, ok := e.(map[string]interface{})
    if ! ok {
      return nil, errors.New("plugins.unwrap unwrapping struct but src element is not map[string]interface{}")
    }
    // recursively go through all fields of the map and unwrap them
    for k, v := range assertedElement {
      // unwrap the fields in the map
      fmt.Printf("plugins.unwrap unwrapping value at key: %s\n", k)
      result, err := unwrap(v)
      if err != nil {
        return nil, err
      }

      // set the value in obj
      objValue.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(result))
    }
    fmt.Printf("%s\n=== Returning ===\n", spew.Sdump(obj))
    return obj, nil
  case reflect.Slice:
    // slices are represented as []interface{} in JSON
    assertedElement, ok := e.([]interface{})
    if ! ok {
      return nil, errors.New("plugins.unwrap unwrapping slice but src element is not []interface{}")
    }
    // recursively go through all indexes and unwrap them
    for i, v := range assertedElement {
      fmt.Printf("plugins.unwrap unwrapping index: %d\n", i)
      result, err := unwrap(v)
      if err != nil {
        return nil, err
      }
      // append the result to the slice
      objValue = reflect.Append(objValue, reflect.ValueOf(result))
    }
    fmt.Printf("%s\n=== Returning ===\n", spew.Sdump(obj))
    return objValue.Interface(), nil
  default:
    return nil, fmt.Errorf("plugins.unwrap unknown kind: %s", kind.String())
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
    fmt.Println(err)
  }
  fmt.Println("HERE HERE HERE")
  spew.Dump(doc)
  return nil, nil
}
