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

func unwrapInlineLink(value interface{}) (*types.InlineLink, error) {
  log.Trace("Unwrapping a InlineLink")

  // if it's nil, just return a nil of its type
  if value == nil || reflect.ValueOf(value).IsNil() {
    return (*types.InlineLink)(nil), nil
  }

  // make sure all parts are present
  inlineLink, ok := value.(map[string]interface{})
  if ! ok {
    return nil, errors.New("'InlineLink' type not map[string]interface{}")
  }
  attributes, ok := inlineLink["attributes"]
  if ! ok {
    return nil, errors.New("InlineLink does not contain 'attributes'")
  }
  location, ok := inlineLink["location"]
  if ! ok {
    return nil, errors.New("InlineLink does not contain 'location'")
  }

  // unwrap non-atomic parts
  log.Trace("Unwrapping InlineLink attributes")
  attributes, err := unwrap(attributes)
  if err != nil {
    return nil, err
  }
  log.Trace("Unwrapping InlineLink location")
  location, err = unwrap(location)
  if err != nil {
    return nil, err
  }

  // assert the types of the parts
  assertedAttributes, ok := attributes.(types.Attributes)
  if ! ok {
    return nil, errors.New("InlineLink Attributes is not type Attributes")
  }
  assertedLocation, ok := location.(*types.Location)
  if ! ok {
    return nil, errors.New("InlineLink Location is not type *Location")
  }

  // build object
  return &types.InlineLink{
    Attributes: assertedAttributes,
    Location: assertedLocation,
  }, nil

}

func unwrapInlinePassthrough(value interface{}) (*types.InlinePassthrough, error) {
  log.Trace("Unwrapping a InlinePassthrough")

  // if it's nil, just return a nil of its type
  if value == nil || reflect.ValueOf(value).IsNil() {
    return (*types.InlinePassthrough)(nil), nil
  }

  // make sure all parts are present
  inlinePassthrough, ok := value.(map[string]interface{})
  if ! ok {
    return nil, errors.New("'InlinePassthrough' type not map[string]interface{}")
  }
  kind, ok := inlinePassthrough["kind"]
  if ! ok {
    return nil, errors.New("InlinePassthrough does not contain 'kind'")
  }
  elements, ok := inlinePassthrough["elements"]
  if ! ok {
    return nil, errors.New("InlinePassthrough does not contain 'elements'")
  }

  // unwrap non-atomic parts
  log.Trace("Unwrapping InlinePassthrough kind")
  kind, err := unwrap(kind)
  if err != nil {
    return nil, err
  }
  log.Trace("Unwrapping InlinePassthrough elements")
  elements, err = unwrap(elements)
  if err != nil {
    return nil, err
  }

  // assert the types of the parts
  assertedKind, ok := kind.(types.PassthroughKind)
  if ! ok {
    return nil, errors.New("InlinePassthrough Kind is not type PassthroughKind")
  }
  assertedElements, ok := elements.([]interface{})
  if ! ok {
    return nil, errors.New("InlinePassthrough Elements is not type []interface{}")
  }

  // build object
  return &types.InlinePassthrough{
    Kind: assertedKind,
    Elements: assertedElements,
  }, nil
}

func unwrapInterface(value interface{}) (interface{}, error) {
  log.Trace("Unwrapping a interface{}")

  // if it's nil, just return a nil of its type
  if value == nil || reflect.ValueOf(value).IsNil() {
    log.Trace("interface{} is nil")
    return (interface{})(nil), nil
  }

  // make sure the type is actually correct
  inter, ok := value.(interface{})
  if ! ok {
    return nil, errors.New("'interface{}' type not interface{}")
  }

  return unwrap(inter)
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

func unwrapList(value interface{}) (*types.List, error) {
  log.Trace("Unwrapping a List")

  // if it's nil, just return a nil of its type
  if value == nil {
    log.Trace("List is nil")
    return (*types.List)(nil), nil
  }

  // make sure all parts are present
  list, ok := value.(map[string]interface{})
  if ! ok {
    return nil, errors.New("'List' type not map[string]interface{}")
  }
  kind, ok := list["kind"]
  if ! ok {
    return nil, errors.New("List does not contain 'kind'")
  }
  attributes, ok := list["attributes"]
  if ! ok {
    return nil, errors.New("List does not contain 'attributes'")
  }
  elements, ok := list["elements"]
  if ! ok {
    return nil, errors.New("List does not contain 'elements'")
  }

  // unwrap non-atomic parts
  log.Trace("Unwrapping List kind")
  kind, err := unwrap(kind)
  if err != nil {
    return nil, err
  }
  log.Trace("Unwrapping List attributes")
  attributes, err = unwrap(attributes)
  if err != nil {
    return nil, err
  }
  log.Trace("Unwrapping List elements")
  elements, err = unwrap(elements)
  if err != nil {
    return nil, err
  }

  // assert the types of the parts
  assertedKind, ok := kind.(types.ListKind)
  if ! ok {
    return nil, errors.New("List Kind is not type ListKind")
  }
  assertedAttributes, ok := attributes.(types.Attributes)
  if ! ok {
    return nil, errors.New("List Attributes is not type ListKind")
  }
  assertedElements, ok := elements.([]types.ListElement)
  if ! ok {
    return nil, errors.New("List Elements is not type []ListElement")
  }

  // build object
  return &types.List{
    Kind: assertedKind,
    Attributes: assertedAttributes,
    Elements: assertedElements,
  }, nil
}

func unwrapListElementSlice(value interface{}) ([]types.ListElement, error) {
  log.Trace("Unwrapping a []ListElement")

  // this should be a slice of interfaces
  listElements, ok := value.([]interface{})
  if ! ok {
    return nil, errors.New("'[]ListElement' is not type []interface{}")
  }

  // unwrap each part of the slice and build a new slice
  newSlice := []types.ListElement{}
  for i, val := range listElements {
    log.Tracef("Unwrapping []ListElement index %d", i)
    obj, err := unwrap(val)
    if err != nil {
      return nil, err
    }
    listElement, ok := obj.(types.ListElement)
    if ! ok {
      return nil, errors.New("[]ListElement element is not of type ListElement")
    }
    newSlice = append(newSlice, listElement)
  }
  return newSlice, nil
}

func unwrapListKind(value interface{}) (types.ListKind, error) {
  log.Trace("Unwrapping a ListKind")

  // ListKinds are atomic strings
  kind, ok := value.(string)
  if ! ok {
    return "", errors.New("'ListKind' is not type string")
  }

  return types.ListKind(kind), nil
}

func unwrapLocation(value interface{}) (*types.Location, error) {
  log.Trace("Unwrapping a Location")

  // if it's nil, just return a nil of its type
  if value == nil {
    log.Trace("Location is nil")
    return (*types.Location)(nil), nil
  }

  // make sure all parts are present
  location, ok := value.(map[string]interface{})
  if ! ok {
    return nil, errors.New("'Location' type not map[string]interface{}")
  }
  scheme, ok := location["scheme"]
  if ! ok {
    return nil, errors.New("Location does not contain 'scheme'")
  }
  path, ok := location["path"]
  if ! ok {
    return nil, errors.New("Location does not contain 'path'")
  }

  // unwrap non-atomic parts
  log.Trace("Unwrapping Location path")
  path, err := unwrap(path)
  if err != nil {
    return nil, err
  }

  // assert the types of the parts
  assertedScheme, ok := scheme.(string)
  if ! ok {
    return nil, errors.New("Location Scheme is not type string")
  }
  assertedPath, ok := path.(interface{})
  if ! ok {
    return nil, errors.New("Location Path is not type interface{}")
  }

  // build object
  return &types.Location{
    Scheme: assertedScheme,
    Path: assertedPath,
  }, nil
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

func unwrapPassthroughKind(value interface{}) (types.PassthroughKind, error) {
  log.Trace("Unwrapping a PassthroughKind")

  kind, ok := value.(string)
  if ! ok {
    return "", errors.New("PassthroughKind is not type string")
  }

  return types.PassthroughKind(kind), nil
}

func unwrapQuotedString(value interface{}) (*types.QuotedString, error) {
  log.Trace("Unwrapping a QuotedString")

  // if it's nil, just return a nil of its type
  if value == nil {
    log.Trace("QuotedString is nil")
    return (*types.QuotedString)(nil), nil
  }

  // make sure all parts are present
  objMap, ok := value.(map[string]interface{})
  if ! ok {
    return nil, errors.New("'QuotedString' type not map[string]interface{}")
  }
  kind, ok := objMap["kind"]
  if ! ok {
    return nil, errors.New("QuotedString does not contain 'kind'")
  }
  elements, ok := objMap["elements"]
  if ! ok {
    return nil, errors.New("QuotedString does not contain 'elements'")
  }

  // unwrap non-atomic parts
  var err error
  log.Trace("Unwrapping QuotedString kind")
  kind, err = unwrap(kind)
  if err != nil {
    return nil, err
  }
  log.Trace("Unwrapping QuotedString elements")
  elements, err = unwrap(elements)
  if err != nil {
    return nil, err
  }

  // assert the types of the parts
  assertedKind, ok := kind.(types.QuotedStringKind)
  if ! ok {
    return nil, errors.New("QuotedString Kind is not type types.QuotedStringKind")
  }
  assertedElements, ok := elements.([]interface{})
  if ! ok {
    return nil, errors.New("QuotedString Elements is not type []interface{}")
  }

  // build object
  return &types.QuotedString{
    Kind: assertedKind,
    Elements: assertedElements,
  }, nil
}

func unwrapQuotedStringKind(value interface{}) (types.QuotedStringKind, error) {
  log.Trace("Unwrapping a QuotedStringKind")

  // QuotedStringKinds are atomic strings
  kind, ok := value.(string)
  if ! ok {
    return "", errors.New("'QuotedStringKind' is not type string")
  }

  return types.QuotedStringKind(kind), nil
}

func unwrapQuotedText(value interface{}) (*types.QuotedText, error) {
  log.Trace("Unwrapping a QuotedText")

  // if it's nil, just return a nil of its type
  if value == nil {
    log.Trace("QuotedText is nil")
    return (*types.QuotedText)(nil), nil
  }

  // make sure all parts are present
  quotedText, ok := value.(map[string]interface{})
  if ! ok {
    return nil, errors.New("'QuotedText' type not map[string]interface{}")
  }
  kind, ok := quotedText["kind"]
  if ! ok {
    return nil, errors.New("QuotedText does not contain 'kind'")
  }
  elements, ok := quotedText["elements"]
  if ! ok {
    return nil, errors.New("QuotedText does not contain 'elements'")
  }
  attributes, ok := quotedText["attributes"]
  if ! ok {
    return nil, errors.New("QuotedText does not contain 'attributes'")
  }

  // unwrap non-atomic parts
  log.Trace("Unwrapping QuotedText kind")
  kind, err := unwrap(kind)
  if err != nil {
    return nil, err
  }
  log.Trace("Unwrapping QuotedText elements")
  elements, err = unwrap(elements)
  if err != nil {
    return nil, err
  }
  log.Trace("Unwrapping QuotedText attributes")
  attributes, err = unwrap(attributes)
  if err != nil {
    return nil, err
  }

  // assert the types of the parts
  assertedKind, ok := kind.(types.QuotedTextKind)
  if ! ok {
    return nil, errors.New("QuotedText Kind is not type QuotedTextKind")
  }
  assertedElements, ok := elements.([]interface{})
  if ! ok {
    return nil, errors.New("QuotedText Elements is not type []interface{}")
  }
  assertedAttributes, ok := attributes.(types.Attributes)
  if ! ok {
    return nil, errors.New("QuotedText Attributes is not type Attributes")
  }

  // build object
  return &types.QuotedText{
    Kind: assertedKind,
    Elements: assertedElements,
    Attributes: assertedAttributes,
  }, nil
}

func unwrapQuotedTextKind(value interface{}) (types.QuotedTextKind, error) {
  log.Trace("Unwrapping a QuotedTextKind")

  // QuotedTextKinds are atomic strings
  kind, ok := value.(string)
  if ! ok {
    return "", errors.New("'QuotedTextKind' is not type string")
  }

  return types.QuotedTextKind(kind), nil
}

func unwrapSpecialCharacter(value interface{}) (*types.SpecialCharacter, error) {
  log.Trace("Unwrapping a SpecialCharacter")

  // if it's nil, just return a nil of its type
  if value == nil {
    log.Trace("SpecialCharacter is nil")
    return (*types.SpecialCharacter)(nil), nil
  }

  // make sure all parts are present
  char, ok := value.(map[string]interface{})
  if ! ok {
    return nil, errors.New("'SpecialCharacter' type not map[string]interface{}")
  }
  name, ok := char["name"]
  if ! ok {
    return nil, errors.New("SpecialCharacter does not contain 'name'")
  }

  // assert the types of the parts
  assertedName, ok := name.(string)
  if ! ok {
    return nil, errors.New("SpecialCharacter Name is not type string")
  }

  // build object
  return &types.SpecialCharacter{
    Name: assertedName,
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
  assertedLevel, ok := level.(float64) // this JSON lib makes all nums float64
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
    Level: int(assertedLevel), // convert from float64 to int
    Attributes: assertedAttributes,
    Title: assertedTitle,
    Elements: assertedElements,
  }, nil
}

// strings will be wrapped if they are stored in an interface{} and their type
// info is needed
func unwrapString(value interface{}) (string, error) {
  log.Trace("Unwrapping a string")

  str, ok := value.(string)
  if ! ok {
    return "", errors.New("'string' type not string")
  }

  return str, nil
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

func unwrapSymbol(value interface{}) (*types.Symbol, error) {
  log.Trace("Unwrapping a Symbol")

  // if it's nil, just return a nil of its type
  if value == nil || reflect.ValueOf(value).IsNil() {
    log.Trace("Symbol is nil")
    return (*types.Symbol)(nil), nil
  }

  // make sure all parts are present
  symbol, ok := value.(map[string]interface{})
  if ! ok {
    return nil, errors.New("'Symbol' type not map[string]interface{}")
  }
  prefix, ok := symbol["prefix"]
  if ! ok {
    return nil, errors.New("Symbol does not contain 'prefix'")
  }
  name, ok := symbol["name"]
  if ! ok {
    return nil, errors.New("Symbol does not contain 'name'")
  }

  // assert the types of the parts
  assertedPrefix, ok := prefix.(string)
  if ! ok {
    return nil, errors.New("Symbol Prefix is not type string")
  }
  assertedName, ok := name.(string)
  if ! ok {
    return nil, errors.New("Symbol Name is not type string")
  }

  return &types.Symbol{
    Prefix: assertedPrefix,
    Name: assertedName,
  }, nil
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
  assertedMaxDepth, ok := maxDepth.(float64) // JSON makes all nums float64
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
    MaxDepth: int(assertedMaxDepth), // convert from float64 to int
    Sections: assertedSections,
  }, nil
}

func unwrapToCSectionSlice(value interface{}) ([]*types.ToCSection, error) {
  log.Trace("Unwrapping a []ToCSection")

  // if it's nil, just return a nil of its type
  if value == nil || reflect.ValueOf(value).IsNil() {
    log.Trace("[]ToCSection is nil")
    return ([]*types.ToCSection)(nil), nil
  }

  slice, ok := value.([]interface{})
  if ! ok {
    return nil, errors.New("'[]ToCSection' type not []interface{}")
  }

  newSlice := []*types.ToCSection{}
  for i, elem := range slice {
    log.Tracef("Unwrapping []ToCSection index %d", i)
    toCSection, err := unwrap(elem)
    if err != nil {
      return nil, err
    }
    assertedToCSection, ok := toCSection.(*types.ToCSection)
    if ! ok {
      return nil, errors.New("[]ToCSection element is not type *ToCSection")
    }
    newSlice = append(newSlice, assertedToCSection)
  }
  return newSlice, nil
}

func unwrapUnorderedListElement(value interface{}) (*types.UnorderedListElement, error) {
  log.Trace("Unwrapping a UnorderedListElement")

  // if it's nil, just return a nil of its type
  if value == nil || reflect.ValueOf(value).IsNil() {
    log.Trace("UnorderedListElement is nil")
    return (*types.UnorderedListElement)(nil), nil
  }

  // make sure all parts are present
  ule, ok := value.(map[string]interface{})
  if ! ok {
    return nil, errors.New("'UnorderedListElement' type not map[string]interface{}")
  }
  bulletStyle, ok := ule["bulletStyle"]
  if ! ok {
    return nil, errors.New("UnorderedListElement does not contain 'bulletStyle'")
  }
  checkStyle, ok := ule["checkStyle"]
  if ! ok {
    return nil, errors.New("UnorderedListElement does not contain 'checkStyle'")
  }
  attributes, ok := ule["attributes"]
  if ! ok {
    return nil, errors.New("UnorderedListElement does not contain 'attributes'")
  }
  elements, ok := ule["elements"]
  if ! ok {
    return nil, errors.New("UnorderedListElement does not contain 'elements'")
  }

  // recurse into non-atomic parts
  log.Trace("Unwrapping UnorderedListElement BulletStyle")
  bulletStyle, err := unwrap(bulletStyle)
  if err != nil {
    return nil, err
  }
  log.Trace("Unwrapping UnorderedListElement CheckStyle")
  checkStyle, err = unwrap(checkStyle)
  if err != nil {
    return nil, err
  }
  log.Trace("Unwrapping UnorderedListElement Attributes")
  attributes, err = unwrap(attributes)
  if err != nil {
    return nil, err
  }
  log.Trace("Unwrapping UnorderedListElement Elements")
  elements, err = unwrap(elements)
  if err != nil {
    return nil, err
  }

  // assert the types of the parts
  assertedBulletStyle, ok := bulletStyle.(types.UnorderedListElementBulletStyle)
  if ! ok {
    return nil, errors.New("UnorderedListElement BulletStyle is not type UnorderedListElementBulletStyle")
  }
  assertedCheckStyle, ok := checkStyle.(types.UnorderedListElementCheckStyle)
  if ! ok {
    return nil, errors.New("UnorderedListElement CheckStyle is not type UnorderedListElementCheckStyle")
  }
  assertedAttributes, ok := attributes.(types.Attributes)
  if ! ok {
    return nil, errors.New("UnorderedListElement Attributes is not type Attributes")
  }
  assertedElements, ok := elements.([]interface{})
  if ! ok {
    return nil, errors.New("UnorderedListElement Elements is not type []interface{}")
  }

  // build object
  return &types.UnorderedListElement{
    BulletStyle: assertedBulletStyle,
    CheckStyle: assertedCheckStyle,
    Attributes: assertedAttributes,
    Elements: assertedElements,
  }, nil
}

func unwrapUnorderedListElementBulletStyle(value interface{}) (types.UnorderedListElementBulletStyle, error) {
  log.Trace("Unwrapping a UnorderedListElementBulletStyle")

  // assert the value
  ulebs, ok := value.(string)
  if ! ok {
    return "", errors.New("'UnorderedListElementBulletStyle' type not string")
  }
  return types.UnorderedListElementBulletStyle(ulebs), nil
}

func unwrapUnorderedListElementCheckStyle(value interface{}) (types.UnorderedListElementCheckStyle, error) {
  log.Trace("Unwrapping a UnorderedListElementCheckStyle")

  // assert the value
  ulecs, ok := value.(string)
  if ! ok {
    return "", errors.New("'UnorderedListElementCheckStyle' type not string")
  }
  return types.UnorderedListElementCheckStyle(ulecs), nil
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
  case "AttributeDeclaration":
    returnValue, err = unwrapAttributeDeclaration(value)
  case "DelimitedBlock":
    returnValue, err = unwrapDelimitedBlock(value)
  case "Document":
    returnValue, err = unwrapDocument(value)
  case "DocumentHeader":
    returnValue, err = unwrapDocumentHeader(value)
  case "ElementReferences":
    returnValue, err = unwrapElementReferences(value)
  case "[]Footnote":
    returnValue, err = unwrapFootnoteSlice(value)
  case "InlineImage":
    returnValue, err = unwrapInlineImage(value)
  case "InlineLink":
    returnValue, err = unwrapInlineLink(value)
  case "InlinePassthrough":
    returnValue, err = unwrapInlinePassthrough(value)
  case "[]interface{}":
    returnValue, err = unwrapInterfaceSlice(value)
  case "InternalCrossReference":
    returnValue, err = unwrapInternalCrossReference(value)
  case "List":
    returnValue, err = unwrapList(value)
  case "[]ListElement":
    returnValue, err = unwrapListElementSlice(value)
  case "ListKind":
    returnValue, err = unwrapListKind(value)
  case "Location":
    returnValue, err = unwrapLocation(value)
  case "OrderedListElement":
    returnValue, err = unwrapOrderedListElement(value)
  case "Paragraph":
    returnValue, err = unwrapParagraph(value)
  case "PassthroughKind":
    returnValue, err = unwrapPassthroughKind(value)
  case "Preamble":
    returnValue, err = unwrapPreamble(value)
  case "QuotedString":
    returnValue, err = unwrapQuotedString(value)
  case "QuotedStringKind":
    returnValue, err = unwrapQuotedStringKind(value)
  case "QuotedText":
    returnValue, err = unwrapQuotedText(value)
  case "QuotedTextKind":
    returnValue, err = unwrapQuotedTextKind(value)
  case "Section":
    returnValue, err = unwrapSection(value)
  case "SpecialCharacter":
    returnValue, err = unwrapSpecialCharacter(value)
  case "string":
    returnValue, err = unwrapString(value)
  case "StringElement":
    returnValue, err = unwrapStringElement(value)
  case "Symbol":
    returnValue, err = unwrapSymbol(value)
  case "TableOfContents":
    returnValue, err = unwrapTableOfContents(value)
  case "[]ToCSection":
    returnValue, err = unwrapToCSectionSlice(value)
  case "ToCSection":
    returnValue, err = unwrapToCSection(value)
  case "UnorderedListElement":
    returnValue, err = unwrapUnorderedListElement(value)
  case "UnorderedListElementBulletStyle":
    returnValue, err = unwrapUnorderedListElementBulletStyle(value)
  case "UnorderedListElementCheckStyle":
    returnValue, err = unwrapUnorderedListElementCheckStyle(value)
  default:
    err = fmt.Errorf("unwrap: unknown type %s", srcType)
  }

  log.Tracef("unwrap returning:\nreturnValue = %serr = %s", spew.Sdump(returnValue), spew.Sdump(err))
  return returnValue, err
}
