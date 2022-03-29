package plugins

import (
  log "github.com/sirupsen/logrus"
//	"github.com/pkg/errors"

	"github.com/bytesparadise/libasciidoc/pkg/types"

  "os"
  "os/exec"
  "bytes"
)

func LoadPlugins(pluginPaths []string) ([]*exec.Cmd, error) {
  // attempt to load every plugin in the paths specified
  loadedPlugins := []*exec.Cmd{}
  for _, path := range(pluginPaths) {
    log.Debugf("Loading plugin %s", path)
    loadedPlugins = append(loadedPlugins, exec.Command(path))
  }
  return loadedPlugins, nil
}

func RunPreRender(doc *types.Document, plugins []*exec.Cmd) (*types.Document, error) {
  for _, cmd := range plugins {
    log.Debugf("Running plugin %s", cmd.Path)

    // convert doc to JSON
    json, err := MarshalJSON(doc)
    if err != nil {
      return nil, err
    }

    // run command passing the JSON on stdin and catching stdout and stderr
    var stdout bytes.Buffer
    cmd.Stdin = bytes.NewBuffer(json)
    cmd.Stdout = &stdout
    cmd.Stderr = os.Stderr
    err = cmd.Run()
    if err != nil {
      return nil, err
    }

    // convert response JSON back to doc
    doc, err = UnmarshalJSON(stdout.Bytes())
    if err != nil {
      return nil, err
    }
  }
  return doc, nil
}
