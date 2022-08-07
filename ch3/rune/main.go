package main

import (
  "fmt"
  "unicode/utf8"
)

func HasPrefix(s, prefix string) bool {
  return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
  return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func main()  {
  s := "Hello, 世界"
  fmt.Println(len(s))
  fmt.Println(utf8.RuneCountInString(s))

  for i := 0; i < len(s); {
    r, size := utf8.DecodeLastRuneInString(s[i:])
    fmt.Printf("%d\t%c\n", i, r)
    i += size
  }

  for i, r := range s {
    fmt.Printf("%d\t%q\t%d\n", i, r, r)
  }
}
