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

func wrap(src interface{}) Wrap {
  log.Tracef("wrap:\nsrc = %s", spew.Sdump(src))

  var returnValue Wrap

  switch elem := src.(type) {
  case types.Attributes:
    returnValue = wrapAttributes(elem)
  case *types.Document:
    returnValue = wrapDocument(elem)
  case *types.DocumentHeader:
    returnValue = wrapDocumentHeader(elem)
  case types.ElementReferences:
    returnValue = wrapElementReferences(elem)
  case []*types.Footnote:
    returnValue = wrapFootnoteSlice(elem)
  case []interface{}:
    returnValue = wrapInterfaceSlice(elem)
  case *types.Paragraph:
    returnValue = wrapParagraph(elem)
  case *types.Section:
    returnValue = wrapSection(elem)
  case *types.StringElement:
    returnValue = wrapStringElement(elem)
  case *types.TableOfContents:
    returnValue = wrapTableOfContents(elem)
  default:
    returnValue.Type = "unknown"
    returnValue.Value = reflect.TypeOf(src).String()
    log.Warnf("wrap: Unknown type %s", returnValue.Value)
  }

  log.Tracef("wrap returning:\nreturnValue = %s", spew.Sdump(returnValue))
  return returnValue
}
