package main

import (
  "fmt"
  "os"
  "strconv"
)

const FeetInMile = 5280

type Mile = float64
type Feet = float64

func main() {
  for _, arg := range os.Args[1:] {
    val, err := strconv.ParseFloat(arg, 64)
    if err != nil {
      fmt.Fprintf(os.Stderr, "fm: %v\n", err)
      os.Exit(1)
    }
    f := Feet(val)
    m := Mile(val)
    fmt.Printf("%v feet = %v miles, %v miles = %v feet\n", f, FeetToMiles(f), m, MilesToFeet(m))
  }
}

func FeetToMiles(feet Feet) Mile {
  return Mile(feet / FeetInMile)
}

func MilesToFeet(miles Mile) Feet {
  return Feet(FeetInMile * miles)
}