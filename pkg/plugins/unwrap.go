package plugins

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
  "github.com/davecgh/go-spew/spew"
  "github.com/pkg/errors"

  log "github.com/sirupsen/logrus"

  "fmt"
  "reflect"
)

func unwrapAttributes(value interface{}) (types.Attributes, error) {
  log.Trace("Unwrapping a Attributes")

  // if it's nil, just return a nil of its type
  if value == nil || reflect.ValueOf(value).IsNil() {
    log.Trace("Attributes is nil")
    return (types.Attributes)(nil), nil
  }

  // make sure the type is correct
  attr, ok := value.(map[string]interface{})
  if ! ok {
    return nil, errors.New("'Attributes' type not map[string]interface{}")
  }

  // recurse through each k,v and build Attributes
  obj := types.Attributes{}
  for k,v := range(attr) {
    var result interface{}
    log.Tracef("Unwrapping value at key %s", k)
    result, err := unwrap(v)
    if err != nil {
      return nil, err
    }
    obj[k] = result
  }

  return obj, nil
}

func unwrapDocument(value interface{}) (*types.Document, error) {
  log.Trace("Unwrapping a Document")

  // make sure all parts are present
  doc, ok := value.(map[string]interface{})
  if ! ok {
    return nil, errors.New("'Document' type not map[string]interface{}")
  }
  elements, ok := doc["elements"]
  if ! ok {
    return nil, errors.New("Document does not contain 'elements'")
  }
  elementReferences, ok := doc["elementReferences"]
  if ! ok {
    return nil, errors.New("Document does not contain 'elementReferences'")
  }
  footnotes, ok := doc["footnotes"]
  if ! ok {
    return nil, errors.New("Document does not contain 'footnotes'")
  }
  toc, ok := doc["tableOfContents"]
  if ! ok {
    return nil, errors.New("Document does not contain 'tableOfContents'")
  }

  // unwrap non-atomic parts
  log.Trace("Unwrapping Document Elements")
  elements, err := unwrap(elements)
  if err != nil {
    return nil, err
  }
  log.Trace("Unwrapping Document ElementReferences")
  elementReferences, err = unwrap(elementReferences)
  if err != nil {
    return nil, err
  }
  log.Trace("Unwrapping Document Footnotes")
  footnotes, err = unwrap(footnotes)
  if err != nil {
    return nil, err
  }
  log.Trace("Unwrapping Document TableOfContents")
  toc, err = unwrap(toc)
  if err != nil {
    return nil, err
  }

  // assert the types of the parts
  assertedElements, ok := elements.([]interface{})
  if ! ok {
    return nil, errors.New("Document Elements is not type []interface{}")
  }
  assertedElementReferences, ok := elementReferences.(types.ElementReferences)
  if ! ok {
    return nil, errors.New("Document ElementReferences is not type ElementReferences")
  }
  assertedFootnotes, ok := footnotes.([]*types.Footnote)
  if ! ok {
    return nil, errors.New("Document Footnotes is not type []*types.Footnote")
  }
  assertedToc, ok := toc.(*types.TableOfContents)
  if ! ok {
    return nil, errors.New("Document TableOfContents is not type *TableOfContents")
  }

  // build object
  return &types.Document{
    Elements: assertedElements,
    ElementReferences: assertedElementReferences,
    Footnotes: assertedFootnotes,
    TableOfContents: assertedToc,
  }, nil
}

func unwrapDocumentHeader(value interface{}) (*types.DocumentHeader, error) {
  log.Trace("Unwrapping DocumentHeader")

  // if it's nil, just return a nil of its type
  if value == nil || reflect.ValueOf(value).IsNil() {
    log.Trace("DocumentHeader is nil")
    return (*types.DocumentHeader)(nil), nil
  }

  // make sure all parts are present
  docHead, ok := value.(map[string]interface{})
  if ! ok {
    return nil, errors.New("'DocumentHeader' type not map[string]interface{}")
  }
  title, ok := docHead["title"]
  if ! ok {
    return nil, errors.New("DocumentHeader does not contain 'title'")
  }
  attributes, ok := docHead["attributes"]
  if ! ok {
    return nil, errors.New("DocumentHeader does not contain 'attributes'")
  }
  elements, ok := docHead["elements"]
  if ! ok {
    return nil, errors.New("DocumentHeader does not contain 'elements'")
  }

  // unwrap non-atomic parts
  log.Trace("Unwrapping DocumentHeader title")
  title, err := unwrap(title)
  if err != nil {
    return nil, err
  }
  log.Trace("Unwrapping DocumentHeader attributes")
  attributes, err = unwrap(attributes)
  if err != nil {
    return nil, err
  }
  log.Trace("Unwrapping DocumentHeader elements")
  elements, err = unwrap(elements)
  if err != nil {
    return nil, err
  }

  // assert the types of the parts
  assertedTitle, ok := title.([]interface{})
  if ! ok {
    return nil, errors.New("DocumentHeader Title is not type []interface{}")
  }
  assertedAttributes, ok := attributes.(types.Attributes)
  if ! ok {
    return nil, errors.New("DocumentHeader Attributes is not type Attributes")
  }
  assertedElements, ok := elements.([]interface{})
  if ! ok {
    return nil, errors.New("DocumentHeader Elements is not type []interface{}")
  }

  // build object
  return &types.DocumentHeader{
    Title: assertedTitle,
    Attributes: assertedAttributes,
    Elements: assertedElements,
  }, nil
}

func unwrapElementReferences(value interface{}) (types.ElementReferences, error) {
  log.Trace("Unwrapping ElementReferences")

  // if it's nil, just return a nil of its type
  if value == nil || reflect.ValueOf(value).IsNil() {
    log.Trace("ElementReferences is nil")
    return types.ElementReferences(nil), nil
  }

  // make sure the type is actually correct
  srcMap, ok := value.(map[string]interface{})
  if ! ok {
    return nil, errors.New("'ElementReferences' type not map[string]interface{}")
  }

  // recurse through each k,v and build object
  obj := types.ElementReferences{}
  for k,v := range(srcMap) {
    var result interface{}
    log.Tracef("Unwrapping value at key %s", k)
    result, err := unwrap(v)
    if err != nil {
      return nil, err
    }
    obj[k] = result
  }
  return obj, nil
}

func unwrapFootnoteSlice(value interface{}) ([]*types.Footnote, error) {
  log.Trace("Unwrapping a []Footnote")

  // if it's nil, just return a nil of its type
  if value == nil || reflect.ValueOf(value).IsNil() {
    log.Trace("[]Footnote is nil")
    return []*types.Footnote(nil), nil
  }

  // make sure the type is actually correct
  slice, ok := value.([]interface{})
  if ! ok {
    return nil, errors.New("'[]Footnote' type not []interface{}")
  }

  // unwrap each index and build object
  newSlice := []*types.Footnote{}
  for i, val := range slice {
    var result interface{}
    log.Tracef("Unwrapping index %d", i)
    result, err := unwrap(val)
    if err != nil {
      return nil, err
    }
    assertedFootnote, ok := result.(*types.Footnote)
    if ! ok {
      return nil, fmt.Errorf("index %d is not type Footnote", i)
    }
    newSlice = append(newSlice, assertedFootnote)
  }
  return newSlice, nil
}

func unwrapInterfaceSlice(value interface{}) ([]interface{}, error) {
  log.Trace("Unwrapping a []interface{}")

  // if it's nil, just return a nil of its type
  if value == nil || reflect.ValueOf(value).IsNil() {
    log.Trace("[]interface{} is nil")
    return ([]interface{})(nil), nil
  }

  // make sure the type is actually correct
  slice, ok := value.([]interface{})
  if ! ok {
    return nil, errors.New("'[]interface{}' type not []interface{}")
  }

  // unwrap each index and build object
  newSlice := []interface{}{}
  for i, val := range slice {
    var result interface{}
    log.Tracef("Unwrapping index %d", i)
    result, err := unwrap(val)
    if err != nil {
      return nil, err
    }
    newSlice = append(newSlice, result)
  }
  return newSlice, nil
}

func unwrapParagraph(value interface{}) (*types.Paragraph, error) {
  log.Trace("Unwrapping a Paragraph")

  // if it's nil, just return a nil of its type
  if value == nil {
    log.Trace("Paragraph is nil")
    return (*types.Paragraph)(nil), nil
  }

  // make sure all parts are present
  paragraph, ok := value.(map[string]interface{})
  if ! ok {
    return nil, errors.New("'Paragraph' type not map[string]interface{}")
  }
  attributes, ok := paragraph["attributes"]
  if ! ok {
    return nil, errors.New("Paragraph does not contain 'attributes'")
  }
  elements, ok := paragraph["elements"]
  if ! ok {
    return nil, errors.New("Paragraph does not contain 'elements'")
  }

  // unwrap non-atomic parts
  log.Trace("Unwrapping Paragraph attributes")
  attributes, err := unwrap(attributes)
  if err != nil {
    return nil, err
  }
  log.Trace("Unwrapping Paragraph elements")
  elements, err = unwrap(elements)
  if err != nil {
    return nil, err
  }

  // assert the types of the parts
  assertedAttributes, ok := attributes.(types.Attributes)
  if ! ok {
    return nil, errors.New("Paragraph Attributes is not type Attributes")
  }
  assertedElements, ok := elements.([]interface{})
  if ! ok {
    return nil, errors.New("Paragraph Elements is not type []interface{}")
  }

  // build object
  return &types.Paragraph{
    Attributes: assertedAttributes,
    Elements: assertedElements,
  }, nil
}

func unwrapSection(value interface{}) (*types.Section, error) {
  log.Trace("Unwrapping a Section")

  // if it's nil, just return a nil of its type
  if value == nil {
    log.Trace("Section is nil")
    return (*types.Section)(nil), nil
  }

  // make sure all parts are present
  section, ok := value.(map[string]interface{})
  if ! ok {
    return nil, errors.New("'Section' type not map[string]interface{}")
  }
  level, ok := section["level"]
  if ! ok {
    return nil, errors.New("Section does not contain 'level'")
  }
  attributes, ok := section["attributes"]
  if ! ok {
    return nil, errors.New("Section does not contain 'attributes'")
  }
  title, ok := section["title"]
  if ! ok {
    return nil, errors.New("Section does not contain 'title'")
  }
  elements, ok := section["elements"]
  if ! ok {
    return nil, errors.New("Section does not contain 'elements'")
  }

  // unwrap non-atomic parts
  log.Trace("Unwrapping Section attributes")
  attributes, err := unwrap(attributes)
  if err != nil {
    return nil, err
  }
  log.Trace("Unwrapping Section title")
  title, err = unwrap(title)
  if err != nil {
    return nil, err
  }
  log.Trace("Unwrapping Section elements")
  elements, err = unwrap(elements)
  if err != nil {
    return nil, err
  }

  // assert the types of the parts
  assertedLevel, ok := level.(int)
  if ! ok {
    return nil, errors.New("Section Level is not type int")
  }
  assertedAttributes, ok := attributes.(types.Attributes)
  if ! ok {
    return nil, errors.New("Section Attributes is not type Attributes")
  }
  assertedTitle, ok := title.([]interface{})
  if ! ok {
    return nil, errors.New("Section Title is not type []interface{}")
  }
  assertedElements, ok := elements.([]interface{})
  if ! ok {
    return nil, errors.New("Section Elements is not type []interface{}")
  }

  // build object
  return &types.Section{
    Level: assertedLevel,
    Attributes: assertedAttributes,
    Title: assertedTitle,
    Elements: assertedElements,
  }, nil
}

// StringElements are atomic single-valued structs so we just store Content as
// a string instead of inside a separate map[string]inteface{}
func unwrapStringElement(value interface{}) (*types.StringElement, error) {
  log.Trace("Unwrapping a StringElement")

  // if it's nil, just return a nil of its type
  if value == nil {
    log.Trace("StringElement is nil")
    return (*types.StringElement)(nil), nil
  }

  // make sure content is present
  content, ok := value.(string)
  if ! ok {
    return nil, errors.New("'StringElement' type not string")
  }

  return types.NewStringElement(content)
}

func unwrapTableOfContents(value interface{}) (*types.TableOfContents, error) {
  log.Trace("Unwrapping a TableOfContents")

  // if it's nil, just return a nil of its type
  if value == nil || reflect.ValueOf(value).IsNil() {
    log.Trace("TableOfContents is nil")
    return (*types.TableOfContents)(nil), nil
  }

  // make sure all parts are present
  doc, ok := value.(map[string]interface{})
  if ! ok {
    return nil, errors.New("'TableOfContents' type not map[string]interface{}")
  }
  maxDepth, ok := doc["maxDepth"]
  if ! ok {
    return nil, errors.New("TableOfContents does not contain 'maxDepth'")
  }
  sections, ok := doc["sections"]
  if ! ok {
    return nil, errors.New("TableOfContents does not contain 'sections'")
  }

  // recurse into non-atomic parts
  log.Trace("Unwrapping TableOfContents Sections")
  sections, err := unwrap(sections)
  if err != nil {
    return nil, err
  }

  // assert the types of the parts
  assertedMaxDepth, ok := maxDepth.(int)
  if ! ok {
    return nil, errors.New("TableOfContents MaxDepth is not type int")
  }
  assertedSections, ok := sections.([]*types.ToCSection)
  if ! ok {
    return nil, errors.New("TableOfContents Sections is not type []*ToCSection")
  }

  // build object
  // NOTE: NewTableOfContents function is meant for adding sections as the
  // Document is processed
  return &types.TableOfContents{
    MaxDepth: assertedMaxDepth,
    Sections: assertedSections,
  }, nil
}

func unwrap(src interface{}) (interface{}, error) {
  log.Tracef("unwrap:\nsrc = %s", spew.Sdump(src))

  // take off the outer wrapping of the element
  element, ok := src.(map[string]interface{})
  if ! ok {
    return nil, errors.New("src is not a map[string]interface{}")
  }
  t, ok := element["type"]
  if ! ok {
    return nil, errors.New("'type' not found in map")
  }
  srcType, ok := t.(string)
  if ! ok {
    return nil, errors.New("'type' is not a string")
  }
  value, ok := element["value"]
  if ! ok {
    return nil, errors.New("'value' not found in map")
  }

  var returnValue interface{}
  var err error

  switch srcType {
  case "Attributes":
    returnValue, err = unwrapAttributes(value)
  case "Document":
    returnValue, err = unwrapDocument(value)
  case "DocumentHeader":
    returnValue, err = unwrapDocumentHeader(value)
  case "ElementReferences":
    returnValue, err = unwrapElementReferences(value)
  case "[]Footnote":
    returnValue, err = unwrapFootnoteSlice(value)
  case "[]interface{}":
    returnValue, err = unwrapInterfaceSlice(value)
  case "Paragraph":
    returnValue, err = unwrapParagraph(value)
  case "StringElement":
    returnValue, err = unwrapStringElement(value)
  case "TableOfContents":
    returnValue, err = unwrapTableOfContents(value)
  default:
    err = fmt.Errorf("unwrap: unknown type %s", srcType)
  }

  log.Tracef("unwrap returning:\nreturnValue = %serr = %s", spew.Sdump(returnValue), spew.Sdump(err))
  return returnValue, err
}
