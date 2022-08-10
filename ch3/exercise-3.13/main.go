package main

import (
  "fmt"
  "os"
  "strconv"
)

const (
  B = 1
  KB = B * 1000
  MB = KB * 1000
  GB = MB * 1000
  TB = GB * 1000
  YB = TB * 1000
)

func main() {
  fmt.Printf("%v\n", constDeclarations(os.Args[1]))
}

func constDeclarations(input string) string {
  numValue, err := strconv.Atoi(input)
  switch {
  case err != nil:
    return "you didn't enter a number"
  case numValue >= YB:
    return "you've entered a massive number"
  case numValue >= TB &&  numValue < YB:
    return "you've entered a terabyte"
  case numValue >= GB && numValue < TB:
    return "you've enetered a gigabyte"
  case numValue >= MB && numValue < GB:
    return "you've entered a megabyte"
  case numValue >= KB && numValue < MB:
    return "you've entered a kilobyte"
  default:
    return "you've entered a byte"
  }
}
