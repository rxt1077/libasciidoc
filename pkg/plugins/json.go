package plugins

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

  "encoding/json"
)

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
