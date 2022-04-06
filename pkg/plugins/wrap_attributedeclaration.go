package plugins

// AttributeDeclarations need to be handled differently due to the non-exported
// rawText field and the possiblity of getting a interface{}(nil) Value

import (
  "github.com/pkg/errors"
	"github.com/bytesparadise/libasciidoc/pkg/types"

  log "github.com/sirupsen/logrus"
)

func wrapAttributeDeclaration(obj *types.AttributeDeclaration) Wrap {
  log.Trace("Wrapping a AttributeDeclaration")

  // add type
  wrapped := Wrap {
    Type: "AttributeDeclaration",
    Value: nil,
  }

  // return nil values
  if obj == nil {
    return wrapped
  }

  // pull out rawText
  rawText, _ := obj.RawText()

  // handle interface{}(nil)
  var val interface{}
  if obj.Value == nil {
    val = nil
  } else {
    val = wrap(obj.Value)
  }

  // wrap non-atomic parts and store atoms
  wrapped.Value = map[string]interface{}{
    "name": obj.Name,
    "value": val,
    "rawText": rawText,
  }

  return wrapped
}

func unwrapAttributeDeclaration(value interface{}) (*types.AttributeDeclaration, error) {
  log.Trace("Unwrapping a AttributeDeclaration")

  // if it's nil, just return a nil of its type
  if value == nil {
    log.Trace("AttributeDeclaration is nil")
    return (*types.AttributeDeclaration)(nil), nil
  }

  // make sure all parts are present
  objMap, ok := value.(map[string]interface{})
  if ! ok {
    return nil, errors.New("'AttributeDeclaration' type not map[string]interface{}")
  }
    name, ok := objMap["name"]
    if ! ok {
      return nil, errors.New("AttributeDeclaration does not contain 'name'")
    }
    val, ok := objMap["value"]
    if ! ok {
      return nil, errors.New("AttributeDeclaration does not contain 'value'")
    }
    rawText, ok := objMap["rawText"]
    if ! ok {
      return nil, errors.New("AttributeDeclaration does not contain 'rawText'")
    }

  // unwrap non-atomic parts
  var err error
  log.Trace("Unwrapping AttributeDeclaration value")
  val, err = unwrap(val)
  if err != nil {
    return nil, err
  }

  // assert the types of the parts
  assertedName, ok := name.(string)
  if ! ok {
    return nil, errors.New("AttributeDeclaration Name is not type string")
  }
  assertedValue, ok := val.(interface{})
  if ! ok {
    return nil, errors.New("AttributeDeclaration Value is not type interface{}")
  }
  assertedRawText, ok := rawText.(string)
  if ! ok {
    return nil, errors.New("AttributeDeclaration rawText is not type string")
  }

  // build object
  return types.NewAttributeDeclaration(assertedName, assertedValue, assertedRawText)
}
