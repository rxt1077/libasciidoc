package plugins

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
//	"github.com/pkg/errors"

   "github.com/davecgh/go-spew/spew"

//  "os"
  "os/exec"
)

func LoadPlugins(pluginPaths []string) ([]exec.Cmd, error) {
  // attempt to load every plugin in the paths specified
  loadedPlugins := []exec.Cmd{}
  for _, _ = range(pluginPaths) {
  }
  return loadedPlugins, nil
}


func RunPreRender(doc *types.Document, plugins []exec.Cmd) (*types.Document, error) {
  spew.Dump(doc)
  result, err := MarshalJSON(doc)
  if err != nil {
    return nil, err
  }
  newDoc, err := UnmarshalJSON(result)
  if err != nil {
    return nil, err
  }
  spew.Dump(newDoc)
/*  spew.Dump(doc)
  result, err := MarshalJSON(doc)
  fmt.Println(string(result))
  if err != nil {
    spew.Dump(err)
  }
  var myDoc types.Document
  err = json.Unmarshal([]byte(result), &myDoc)
  if err != nil {
    spew.Dump(err)
  }
  spew.Dump(myDoc)
  return &myDoc, nil

  for _, curPlugin := range plugins {
    var err error
    doc, err = curPlugin.PreRender(doc)
    if err != nil {
      return nil, err
    }
  }*/
  return doc, nil
}

/*
// runs the PreRender plugins
func RunPreRender(doc *types.Document, plugins []Plugin) (*types.Document, error) {
	for _, curPlugin := range plugins {
		log.Debugf("plugins: running %s PreRender\n", curPlugin.Path)
		var err error
		doc, err = curPlugin.PreRender(doc)
		if err != nil {
			return nil, err
		}
	}
	return doc, nil
}*/
