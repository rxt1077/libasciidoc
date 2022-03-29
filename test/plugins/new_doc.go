package main

import (
  "os"
  "fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/pkg/plugins"
)

// this plugin replaces whatever it is given with a new blank Document
func main() {
  doc := &types.Document{}
  json, err := plugins.MarshalJSON(doc)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error converting doc to JSON: %s", err)
    return
  }
  _, err = os.Stdout.Write(json)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error writing to Stdout: %s", err)
  }
  return
}
