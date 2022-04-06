package plugins

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
  "github.com/davecgh/go-spew/spew"

  log "github.com/sirupsen/logrus"

  "reflect"
)

// basic struct for storing a value and a type
// uses json field tags to make it easier to look at
type Wrap struct {
  Type string `json:"type"`
  Value interface{} `json:"value"`
}

func wrapAttributes(attr types.Attributes) Wrap {
  log.Trace("Wrapping a Attributes")

  // add type
  wrapped := Wrap {
    Type: "Attributes",
    Value: nil,
  }

  // return nil values
  if attr == nil {
    return wrapped
  }

  // wrap every attribute in map
  newMap := make(map[string]interface{})
  for key, val := range attr {
    newMap[key] = wrap(val)
  }
  wrapped.Value = newMap

  return wrapped
}

func wrapDocument(doc *types.Document) Wrap {
  log.Trace("Wrapping a Document")

  // add type
  wrapped := Wrap {
    Type: "Document",
    Value: nil,
  }

  // return nil values
  if doc == nil {
    return wrapped
  }

  // wrap non-atomic parts
  wrapped.Value = map[string]interface{}{
    "elements": wrap(doc.Elements),
    "elementReferences": wrap(doc.ElementReferences),
    "footnotes": wrap(doc.Footnotes),
    "tableOfContents": wrap(doc.TableOfContents),
  }

  return wrapped
}

func wrapDocumentHeader(docHead *types.DocumentHeader) Wrap {
  log.Trace("Wrapping a DocumentHeader")

  // add type
  wrapped := Wrap {
    Type: "DocumentHeader",
    Value: nil,
  }

  // return nil values
  if docHead == nil {
    return wrapped
  }

  // wrap non-atomic parts
  log.Tracef("Wrapping DocumentHeader Title")
  title := wrap(docHead.Title)
  log.Tracef("Wrapping DocumentHeader Attributes")
  attributes := wrap(docHead.Attributes)
  log.Tracef("Wrapping DocumentHeader Elements")
  elements := wrap(docHead.Elements)

  // create object
  wrapped.Value = map[string]interface{}{
    "title": title,
    "attributes": attributes,
    "elements": elements,
  }

  return wrapped
}

func wrapElementReferences(refs types.ElementReferences) Wrap {
  // add type
  wrapped := Wrap {
    Type: "ElementReferences",
    Value: nil,
  }

  // return nil values
  if refs == nil {
    return wrapped
  }

  // wrap every ref in map
  newMap := make(map[string]interface{})
  for key, val := range refs {
    newMap[key] = wrap(val)
  }
  wrapped.Value = newMap

  return wrapped
}

func wrapFootnoteSlice(footnotes []*types.Footnote) Wrap {
  log.Tracef("Wrapping []Footnote")

  // add type
  wrapped := Wrap {
    Type: "[]Footnote",
    Value: nil,
  }

  // return nil values
  if footnotes == nil {
    return wrapped
  }

  // wrap every footnote in slice
  var newSlice []interface{}
  for _, val := range footnotes {
    newSlice = append(newSlice, wrap(val))
  }
  wrapped.Value = newSlice

  return wrapped
}

func wrapInlineLink(inlineLink *types.InlineLink) Wrap {
  log.Trace("Wrapping InlineLink")

  // add type
  wrapped := Wrap {
    Type: "InlineLink",
    Value: nil,
  }

  // return nil values
  if inlineLink == nil {
    return wrapped
  }

  // wrap non-atomic parts
  wrapped.Value = map[string]interface{}{
    "attributes": wrap(inlineLink.Attributes),
    "location": wrap(inlineLink.Location),
  }

  return wrapped
}

func wrapInlinePassthrough(inlinePassthrough *types.InlinePassthrough) Wrap {
  log.Trace("Wrapping InlinePassthrough")

  // add type
  wrapped := Wrap {
    Type: "InlinePassthrough",
    Value: nil,
  }

  // return nil values
  if inlinePassthrough == nil {
    return wrapped
  }

  // wrap non-atomic parts
  wrapped.Value = map[string]interface{}{
    "kind": wrap(inlinePassthrough.Kind),
    "elements": wrap(inlinePassthrough.Elements),
  }

  return wrapped
}

func wrapInterface(obj interface{}) Wrap {
  log.Trace("Wrapping interface{}")

  // add type
  wrapped := Wrap {
    Type: "interface{}",
    Value: nil,
  }

  // return nil values
  if obj == nil {
    return wrapped
  }

  wrapped.Value = wrap(obj)

  return wrapped
}

func wrapInterfaceSlice(obj []interface{}) Wrap {
  log.Trace("Wrapping []interface{}")

  // add type
  wrapped := Wrap {
    Type: "[]interface{}",
    Value: nil,
  }

  // return nil values
  if obj == nil {
    return wrapped
  }

  // wrap every value in slice
  var newSlice []interface{}
  for _, val := range obj {
    newSlice = append(newSlice, wrap(val))
  }
  wrapped.Value = newSlice

  return wrapped
}

func wrapListKind(kind types.ListKind) Wrap {
  log.Trace("Wrapping a ListKind")

  return Wrap {
    Type: "ListKind",
    Value: kind,
  }
}

func wrapParagraph(paragraph *types.Paragraph) Wrap {
  log.Trace("Wrapping a Paragraph")

  // add type
  wrapped := Wrap {
    Type: "Paragraph",
    Value: nil,
  }

  if paragraph == nil {
    return wrapped
  }

  // wrap non-atomic parts
  wrapped.Value = map[string]interface{}{
    "attributes": wrap(paragraph.Attributes),
    "elements": wrap(paragraph.Elements),
  }

  return wrapped
}

func wrapList(list *types.List) Wrap {
  log.Trace("Wrapping a Location")

  // add type
  wrapped := Wrap {
    Type: "List",
    Value: nil,
  }

  if list == nil {
    return wrapped
  }

  // wrap atomic and non-atomic parts
  wrapped.Value = map[string]interface{}{
    "kind": wrap(list.Kind),
    "attributes": wrap(list.Attributes),
    "elements": wrap(list.Elements),
  }

  return wrapped
}

func wrapListElementSlice(listElements []types.ListElement) Wrap {
  log.Trace("Wrapping a []ListElement")

  // add type
  wrapped := Wrap {
    Type: "[]ListElement",
    Value: nil,
  }

  // return nil values
  if listElements == nil {
    return wrapped
  }

  // wrap each part of the slice
  newSlice := []interface{}{}
  for i, listElement := range listElements {
    log.Tracef("Wrapping []ListElement element %d", i)
    newSlice = append(newSlice, wrap(listElement))
  }
  wrapped.Value = newSlice

  return wrapped
}

func wrapLocation(location *types.Location) Wrap {
  log.Trace("Wrapping a Location")

  // add type
  wrapped := Wrap {
    Type: "Location",
    Value: nil,
  }

  if location == nil {
    return wrapped
  }

  // wrap atomic and non-atomic parts
  wrapped.Value = map[string]interface{}{
    "scheme": location.Scheme,
    "path": wrap(location.Path),
  }

  return wrapped
}

func wrapPassthroughKind(kind types.PassthroughKind) Wrap {
  log.Trace("Wrapping a PassthroughKind")

  return Wrap {
    Type: "PassthroughKind",
    Value: string(kind),
  }
}

func wrapQuotedString(quotedString *types.QuotedString) Wrap {
  log.Trace("Wrapping a QuotedString")

  // add type
  wrapped := Wrap {
    Type: "QuotedString",
    Value: nil,
  }

  if quotedString == nil {
    return wrapped
  }

  // wrap non-atomic parts
  wrapped.Value = map[string]interface{}{
    "kind": wrap(quotedString.Kind),
    "elements": wrap(quotedString.Elements),
  }

  return wrapped
}

func wrapQuotedStringKind(kind types.QuotedStringKind) Wrap {
  log.Trace("Wrapping a QuotedStringKind")

  return Wrap {
    Type: "QuotedStringKind",
    Value: kind,
  }
}

func wrapQuotedText(quotedText *types.QuotedText) Wrap {
  log.Trace("Wrapping a QuotedText")

  // add type
  wrapped := Wrap {
    Type: "QuotedText",
    Value: nil,
  }

  if quotedText == nil {
    return wrapped
  }

  // wrap non-atomic parts
  wrapped.Value = map[string]interface{}{
    "kind": wrap(quotedText.Kind),
    "elements": wrap(quotedText.Elements),
    "attributes": wrap(quotedText.Attributes),
  }

  return wrapped
}

func wrapQuotedTextKind(kind types.QuotedTextKind) Wrap {
  log.Trace("Wrapping a QuotedTextKind")

  return Wrap {
    Type: "QuotedTextKind",
    Value: kind,
  }
}

func wrapSection(section *types.Section) Wrap {
  log.Trace("Wrapping a Section")

  // add type
  wrapped := Wrap {
    Type: "Section",
    Value: nil,
  }

  if section == nil {
    return wrapped
  }

  // wrap non-atomic parts and set atomic ones
  wrapped.Value = map[string]interface{}{
    "level": section.Level,
    "attributes": wrap(section.Attributes),
    "title": wrap(section.Title),
    "elements": wrap(section.Elements),
  }

  return wrapped
}

func wrapSpecialCharacter(char *types.SpecialCharacter) Wrap {
  log.Trace("Wrapping a SpecialCharacter")

  // add type
  wrapped := Wrap {
    Type: "SpecialCharacter",
    Value: nil,
  }

  if char == nil {
    return wrapped
  }

  // wrap non-atomic parts
  wrapped.Value = map[string]interface{}{
    "name": char.Name,
  }

  return wrapped
}

// sometimes a string will need to be wrapped, like when it is stored as an
// interface{} and we need type data for it
func wrapString(str string) Wrap {
  log.Tracef("Wrapping string")

  return Wrap {
    Type: "string",
    Value: str,
  }
}

// StringElements are atomic single-valued structs so we just store Content as
// a string instead of inside a separate map[string]inteface{}
func wrapStringElement(str *types.StringElement) Wrap {
  log.Tracef("Wrapping StringElement")

  // string elements are atomic, return the object
  return Wrap {
    Type: "StringElement",
    Value: str.Content,
  }
}

func wrapSymbol(symbol *types.Symbol) Wrap {
  log.Tracef("Wrapping Symbol")

  return Wrap{
    Type: "Symbol",
    Value: map[string]interface{}{
      "prefix": symbol.Prefix,
      "name": symbol.Name,
    },
  }
}

func wrapTableOfContents(toc *types.TableOfContents) Wrap {
  // add type
  wrapped := Wrap {
    Type: "TableOfContents",
    Value: nil,
  }

  // return nil values
  if toc == nil {
    return wrapped
  }

  // wrap non-atomic parts and set atomic ones
  wrapped.Value = map[string]interface{}{
    "maxDepth": toc.MaxDepth,
    "sections": wrap(toc.Sections),
  }

  return wrapped
}

// TODO: Can we generalize ALL slice wraps?
func wrapToCSectionSlice(slice []*types.ToCSection) Wrap {
  log.Tracef("Wrapping []ToCSection")

  // add type
  wrapped := Wrap {
    Type: "[]ToCSection",
    Value: nil,
  }

  // return nil values
  if slice == nil {
    return wrapped
  }

  // wrap every element in the slice
  newSlice := []Wrap{}
  for i, elem := range slice {
    log.Tracef("Wrapping []ToCSection index %d", i)
    newSlice = append(newSlice, wrap(elem))
  }
  wrapped.Value = newSlice

  return wrapped
}

func wrapUnorderedListElement(ule *types.UnorderedListElement) Wrap {
  // add type
  wrapped := Wrap {
    Type: "UnorderedListElement",
    Value: nil,
  }

  // return nil values
  if ule == nil {
    return wrapped
  }

  // wrap non-atomic parts
  wrapped.Value = map[string]interface{}{
    "bulletStyle": wrap(ule.BulletStyle),
    "checkStyle": wrap(ule.CheckStyle),
    "attributes": wrap(ule.Attributes),
    "elements": wrap(ule.Elements),
  }

  return wrapped
}

func wrapUnorderedListElementBulletStyle(ulebs types.UnorderedListElementBulletStyle) Wrap {
  return Wrap {
    Type: "UnorderedListElementBulletStyle",
    Value: string(ulebs),
  }
}

func wrapUnorderedListElementCheckStyle(ulecs types.UnorderedListElementCheckStyle) Wrap {
  return Wrap {
    Type: "UnorderedListElementCheckStyle",
    Value: string(ulecs),
  }
}

func wrap(src interface{}) Wrap {
  log.Tracef("wrap:\nsrc = %s", spew.Sdump(src))

  var returnValue Wrap

  switch elem := src.(type) {
  case types.Attributes:
    returnValue = wrapAttributes(elem)
  case *types.AttributeDeclaration:
    returnValue = wrapAttributeDeclaration(elem)
  case *types.DelimitedBlock:
    returnValue = wrapDelimitedBlock(elem)
  case *types.Document:
    returnValue = wrapDocument(elem)
  case *types.DocumentHeader:
    returnValue = wrapDocumentHeader(elem)
  case types.ElementReferences:
    returnValue = wrapElementReferences(elem)
  case []*types.Footnote:
    returnValue = wrapFootnoteSlice(elem)
  case *types.InlineImage:
    returnValue = wrapInlineImage(elem)
  case *types.InlineLink:
    returnValue = wrapInlineLink(elem)
  case *types.InlinePassthrough:
    returnValue = wrapInlinePassthrough(elem)
  case []interface{}:
    returnValue = wrapInterfaceSlice(elem)
  case *types.InternalCrossReference:
    returnValue = wrapInternalCrossReference(elem)
  case *types.List:
    returnValue = wrapList(elem)
  case types.ListKind:
    returnValue = wrapListKind(elem)
  case []types.ListElement:
    returnValue = wrapListElementSlice(elem)
  case *types.Location:
    returnValue = wrapLocation(elem)
  case *types.OrderedListElement:
    returnValue = wrapOrderedListElement(elem)
  case *types.Paragraph:
    returnValue = wrapParagraph(elem)
  case types.PassthroughKind:
    returnValue = wrapPassthroughKind(elem)
  case *types.Preamble:
    returnValue = wrapPreamble(elem)
  case *types.QuotedString:
    returnValue = wrapQuotedString(elem)
  case types.QuotedStringKind:
    returnValue = wrapQuotedStringKind(elem)
  case *types.QuotedText:
    returnValue = wrapQuotedText(elem)
  case types.QuotedTextKind:
    returnValue = wrapQuotedTextKind(elem)
  case *types.Section:
    returnValue = wrapSection(elem)
  case *types.SpecialCharacter:
    returnValue = wrapSpecialCharacter(elem)
  case string:
    returnValue = wrapString(elem)
  case *types.StringElement:
    returnValue = wrapStringElement(elem)
  case *types.Symbol:
    returnValue = wrapSymbol(elem)
  case *types.TableOfContents:
    returnValue = wrapTableOfContents(elem)
  case []*types.ToCSection:
    returnValue = wrapToCSectionSlice(elem)
  case *types.ToCSection:
    returnValue = wrapToCSection(elem)
  case *types.UnorderedListElement:
    returnValue = wrapUnorderedListElement(elem)
  case types.UnorderedListElementBulletStyle:
    returnValue = wrapUnorderedListElementBulletStyle(elem)
  case types.UnorderedListElementCheckStyle:
    returnValue = wrapUnorderedListElementCheckStyle(elem)
  default:
    returnValue.Type = "unknown"
    returnValue.Value = reflect.TypeOf(src).String()
    log.Warnf("wrap: Unknown type %s", returnValue.Value)
  }

  log.Tracef("wrap returning:\nreturnValue = %s", spew.Sdump(returnValue))
  return returnValue
}
