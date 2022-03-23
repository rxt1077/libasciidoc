package main

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
  "github.com/hashicorp/go-plugin"
  "github.com/hashicorp/go-hclog"
	"github.com/bytesparadise/libasciidoc/pkg/plugins"
//   "github.com/davecgh/go-spew/spew"

  "encoding/gob"
  "os"
)

// the handshake must match what's expected
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "LIBASCIIDOC",
	MagicCookieValue: "libasciidoc",
}

// this basic plugin just returns the doc given to it
type PassThroughPlugin struct {
  logger hclog.Logger
}

func (g *PassThroughPlugin) PreRender(doc *types.Document) (*types.Document, error) {
//  g.logger.Debug(spew.Sdump(doc))
	return doc, nil
}

func main() {
  // gob is used to serialize our types, they must be registered beforehand
//  gob.RegisterName("github.com/bytesparadise/libasciidoc/pkg/types.DocumentHeader", types.DocumentHeader{})
  gob.Register(types.DocumentHeader{})
  gob.Register(types.StringElement{})
  gob.Register(types.Preamble{})
  gob.Register(types.Paragraph{})
  gob.Register(types.InlineImage{})
  gob.Register(types.Section{})
  gob.Register(types.List{})
//  gob.Register(types.UnorderedListElement{})
  gob.RegisterName("github.com/bytesparadise/libasciidoc/pkg/types.UnorderedListElement", &types.UnorderedListElement{})
  gob.Register(types.QuotedText{})
  gob.Register(types.InlinePassthrough{})
  gob.Register(types.QuotedString{})
  gob.Register(types.Symbol{})
  gob.Register(types.InlineLink{})
  gob.Register(types.SpecialCharacter{})
  gob.Register(types.DelimitedBlock{})
//  gob.Register(types.OrderedListElement{})
  gob.RegisterName("github.com/bytesparadise/libasciidoc/pkg/types.OrderedListElement", &types.OrderedListElement{})
//  gob.Register(types.ListElement{})
  gob.Register([]interface{}{})

  logger := hclog.New(&hclog.LoggerOptions{
    Level: hclog.Trace,
    Output: os.Stderr,
    JSONFormat: true,
  })

	passPlugin := &PassThroughPlugin{
    logger: logger,
  }

	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{
		"docPlugin": &plugins.DocGoPlugin{Impl: passPlugin},
	}

  // the plugin acts as an RPC server
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugins.HandshakeConfig,
		Plugins:         pluginMap,
	})
}
