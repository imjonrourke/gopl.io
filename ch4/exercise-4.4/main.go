package main

import "fmt"

func main() {
  slicedUp := []int{1, 2, 3, 4, 5}
  //reverse(&slicedUp)
  fmt.Printf("brooo\n%v", rotate(slicedUp, 1))
}

//func reverse(s *[]int) {
//  for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
//    (*s)[i], (*s)[j] = (*s)[j], (*s)[i]
//  }
//}

func rotate(s []int, shift int) []int {
  arrLimit := len(s) - 1
  newArr := make([]int, len(s), cap(s))
  for i := 0; i <= arrLimit; i++ {
    if i <= arrLimit-shift {
      newArr[i] = s[i+shift]
    } else {
      newArr[i] = s[(i - arrLimit + shift - 1)]
    }
  }
  return newArr
}
