 
package plugins

// this file is automatically generated

import (
  "github.com/pkg/errors"
	"github.com/bytesparadise/libasciidoc/pkg/types"

  log "github.com/sirupsen/logrus"
)

func wrapDelimitedBlock(obj *types.DelimitedBlock) Wrap {
  log.Trace("Wrapping a DelimitedBlock")

  // add type
  wrapped := Wrap {
    Type: "DelimitedBlock",
    Value: nil,
  }

  // return nil values
  if obj == nil {
    return wrapped
  }

  // wrap non-atomic parts and store atoms
  wrapped.Value = map[string]interface{}{
  
    "kind": obj.Kind,
  
    "attributes": wrap(obj.Attributes),
  
    "elements": wrap(obj.Elements),
  
  }

  return wrapped
}

func unwrapDelimitedBlock(value interface{}) (*types.DelimitedBlock, error) {
  log.Trace("Unwrapping a DelimitedBlock")

  // if it's nil, just return a nil of its type
  if value == nil {
    log.Trace("DelimitedBlock is nil")
    return (*types.DelimitedBlock)(nil), nil
  }

  // make sure all parts are present
  objMap, ok := value.(map[string]interface{})
  if ! ok {
    return nil, errors.New("'DelimitedBlock' type not map[string]interface{}")
  }
  
    kind, ok := objMap["kind"]
    if ! ok {
      return nil, errors.New("DelimitedBlock does not contain 'kind'")
    }
  
    attributes, ok := objMap["attributes"]
    if ! ok {
      return nil, errors.New("DelimitedBlock does not contain 'attributes'")
    }
  
    elements, ok := objMap["elements"]
    if ! ok {
      return nil, errors.New("DelimitedBlock does not contain 'elements'")
    }
  

  // unwrap non-atomic parts
  var err error
  
    
  
    
      log.Trace("Unwrapping DelimitedBlock attributes")
      attributes, err = unwrap(attributes)
      if err != nil {
        return nil, err
      }
    
  
    
      log.Trace("Unwrapping DelimitedBlock elements")
      elements, err = unwrap(elements)
      if err != nil {
        return nil, err
      }
    
  

  // assert the types of the parts
  
    
      assertedKind, ok := kind.(string)
      if ! ok {
        return nil, errors.New("DelimitedBlock Kind is not type string")
      }
    
  
    
      assertedAttributes, ok := attributes.(types.Attributes)
      if ! ok {
        return nil, errors.New("DelimitedBlock Attributes is not type types.Attributes")
      }
    
  
    
      assertedElements, ok := elements.([]interface{})
      if ! ok {
        return nil, errors.New("DelimitedBlock Elements is not type []interface{}")
      }
    
  

  // build object
  return &types.DelimitedBlock{
    Kind: assertedKind,
    Attributes: assertedAttributes,
    Elements: assertedElements,
  }, nil
}
