package main

import (
  "bytes"
  "fmt"
)

func main() {
  slicedUp := []byte("i said i go i know")
  //slicedUpBuf := bytes.Buffer{"i said i go i know"}

  var bufSliced = bytes.Buffer{}

  fmt.Printf("original\n%v", slicedUp)
  reverse(&slicedUp)
  bufSliced.Write(slicedUp)
  fmt.Printf("broooslice\n%v", bufSliced.String())
}

func reverse(s *[]byte) {
  for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
    (*s)[i], (*s)[j] = (*s)[j], (*s)[i]
  }
}
