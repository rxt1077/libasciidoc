package plugins

// this file houses interfaces and variables used for setting up the
// and using the go-lang plugin package: https://github.com/hashicorp/go-plugin 

import (
	"github.com/hashicorp/go-plugin"
  "github.com/bytesparadise/libasciidoc/pkg/types"
//  "github.com/davecgh/go-spew/spew"

	"net/rpc"
)

// plugins must have the same handshake to be able to connect
var HandshakeConfig = plugin.HandshakeConfig{
  ProtocolVersion: 1,
  MagicCookieKey: "LIBASCIIDOC",
  MagicCookieValue: "libasciidoc",
}

var PluginMap = map[string]plugin.Plugin {
  "docPlugin": &DocGoPlugin{},
}

// net/rpc passed arguments must be bundled in a single struct
type PreRenderRPCArgs struct {
  Doc *types.Document
}

// DocPlugin is the interface that we're exposing as a plugin.
// the term "plugin" is just too overloaded to use alone
// currently we only hook in at one point: PreRender
type DocPlugin interface {
	PreRender(*types.Document) (*types.Document, error)
}

// This is a client implementation that talks over RPC
// NOTE: the main process is technically the client
type DocPluginRPC struct {
  client *rpc.Client
}

func (g *DocPluginRPC) PreRender(doc *types.Document) (*types.Document, error) {
//  spew.Dump(doc)
	var resp types.Document
	err := g.client.Call("Plugin.PreRender", PreRenderRPCArgs{doc}, &resp)
//  spew.Dump(resp)
	if err != nil {
    return nil, err
	}
	return &resp, nil
}

// This is the RPC server that DocPluginRPC talks to, conforming to
// the requirements of net/rpc
// NOTE: the plugin is technically the server
type DocPluginRPCServer struct {
	// This is the real implementation
	Impl DocPlugin
}

func (s *DocPluginRPCServer) PreRender(args PreRenderRPCArgs, resp *types.Document) error {
  resp, _ = s.Impl.PreRender(args.Doc)
	return nil
}

// This is the implementation of plugin.Plugin so we can serve/consume this
//
// This has two methods: Server must return an RPC server for this plugin
// type. We construct a GreeterRPCServer for this.
//
// Client must return an implementation of our interface that communicates
// over an RPC client. We return GreeterRPC for this.
//
// Ignore MuxBroker. That is used to create more multiplexed streams on our
// plugin connection and is a more advanced use case.
type DocGoPlugin struct {
	// Impl Injection
	Impl DocPlugin
}

func (p *DocGoPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &DocPluginRPCServer{Impl: p.Impl}, nil
}

func (DocGoPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &DocPluginRPC{client: c}, nil
}
