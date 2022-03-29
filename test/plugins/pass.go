package main

import (
  "os"
  "io"
  "fmt"

	"github.com/bytesparadise/libasciidoc/pkg/plugins"
  log "github.com/sirupsen/logrus"
)

func ifErrorExit(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "%s", err)
    os.Exit(1)
  }
  return
}

func main() {
  log.SetFormatter(&log.TextFormatter{
    EnvironmentOverrideColors: true,
    DisableLevelTruncation:    true,
    DisableTimestamp:          true,
    DisableQuote:              true,
  })
  log.SetLevel(log.TraceLevel)

  bytes, err := io.ReadAll(os.Stdin)
  ifErrorExit(err)
  doc, err := plugins.UnmarshalJSON(bytes)
  ifErrorExit(err)
  json, err := plugins.MarshalJSON(doc)
  ifErrorExit(err)
  _, err = os.Stdout.Write(json)
  ifErrorExit(err)

  return
}
