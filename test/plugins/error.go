package main

import (
  "os"
  "fmt"
)

// this plugin prints an error on Stderror
func main() {
  fmt.Fprintf(os.Stderr, "Testing error passing")
  return
}
