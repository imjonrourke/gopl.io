package main

import "fmt"

func main() {
  dupArr := []string{"yo", "yo", "what", "is", "is", "up", "up"}

  fmt.Printf("ayyy\n%v", adjacentDedup(dupArr))
}

func adjacentDedup(s []string) []string {
  deduppedArr := s[:0]
  for i := 0; i < len(s); i++ {
    if i == 0 || s[i] != s[i-1] {
      deduppedArr = append(deduppedArr, s[i])
    }
  }
  return deduppedArr
}
