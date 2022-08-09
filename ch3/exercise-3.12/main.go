package main

import (
  "fmt"
  "os"
)

func main() {
  for i := 1; i < len(os.Args); i++ {
    fmt.Printf("%v\n", checkString(os.Args[i]))
  }
}

func checkString(s string) bool {
  n := len(s) - 1
  midPoint := (n + 1)/2
  for i := 0; i <= midPoint; i++ {
    if (s[i] != s[n-i]) {
      return false
    }
  }
  return true
}
