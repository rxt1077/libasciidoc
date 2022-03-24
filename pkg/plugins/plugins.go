package plugins

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

   "github.com/davecgh/go-spew/spew"

  log "github.com/sirupsen/logrus"
	"github.com/pkg/errors"

  "os/exec"
  "bytes"
)

func LoadPlugins(pluginPaths []string) ([]*exec.Cmd, error) {
  // attempt to load every plugin in the paths specified
  loadedPlugins := []*exec.Cmd{}
  for _, path := range(pluginPaths) {
    loadedPlugins = append(loadedPlugins, exec.Command(path))
  }
  return loadedPlugins, nil
}


func RunPreRender(doc *types.Document, plugins []*exec.Cmd) (*types.Document, error) {
  spew.Dump(doc)
  for _, cmd := range plugins {
    log.Debugf("Running plugin %s", cmd.Path)
    // convert doc to JSON
    json, err := MarshalJSON(doc)
    if err != nil {
      return nil, err
    }
    // run command passing the JSON on stdin and catching stdout and stderr
    var stdout, stderr bytes.Buffer
    cmd.Stdin = bytes.NewBuffer(json)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err = cmd.Run()
    if err != nil {
      return nil, err
    }
    errStr := string(stderr.Bytes())
//    log.Debugf("out:\n%s\nerr:\n%s\n", outStr, errStr)
    // any output on stderr is considered an error
    if errStr != "" {
      return nil, errors.New(errStr)
    }
    // convert response JSON back to doc
    doc, err = UnmarshalJSON(stdout.Bytes())
    if err != nil {
      return nil, err
    }
  }
  spew.Dump(doc)
  /*
  spew.Dump(doc)
  result, err := MarshalJSON(doc)
  if err != nil {
    return nil, err
  }
  newDoc, err := UnmarshalJSON(result)
  if err != nil {
    return nil, err
  }*/
//  spew.Dump(newDoc)
  return doc, nil
}
