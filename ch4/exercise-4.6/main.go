package main

import (
  "bytes"
  "fmt"
  "unicode"
)

func main() {
  fmt.Printf("%v", squashSpaces("I said    you   have no ideaaa      "))
}

func squashSpaces(s string) string {
  byteSlice := []byte(s)
  var newString = bytes.Buffer{}
  for i := 0; i < len(s); i++ {
    if i > 0 && unicode.IsSpace(rune(byteSlice[i])) && unicode.IsSpace(rune(byteSlice[i-1])) {
      newString.WriteString(" ")
    } else {
      newString.WriteByte(byteSlice[i])
    }
  }
  return newString.String()
}
