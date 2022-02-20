package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  length := 0

  if len(os.Args) > 1 {
    length = len(strings.Fields(os.Args[1]))
  }

  fmt.Println(length);
}

