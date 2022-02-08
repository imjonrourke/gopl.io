// salaryCalculator will either multiply a current salary by a multiplier or return a multiplier given a target salary.
package main

import (
  "errors"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type Salary = float64
type Multiplier = float64
type ArgName = string

type ArgResult struct {
  Multiplier
  Salary
  name ArgName
}

const currentArg = ArgName("--current")
const cArg = ArgName("-c")
const targetArg = ArgName("--target")
const tArg = ArgName("-t")
const multiplierArg = ArgName("--multiplier")
const mArg = ArgName("-m")

func main() {
  var current ArgResult
  var multi ArgResult
  var target ArgResult
  for _, arg := range os.Args[1:] {
    val, err := parseArg(arg)
    if err != nil {
      fmt.Fprintf(os.Stderr, "%v is invalid\n", err)
    }
    if val.name == currentArg || val.name == cArg {
      current = val
    }
    if val.name == multiplierArg || val.name == mArg {
      multi = val
    }
    if val.name == targetArg || val.name == tArg {
      target = val
    }
  }
  if current.name != "" {
    if multi.name != "" {
      fmt.Printf("Current salary %v will become %v\n", current.Salary, current.Salary * multi.Multiplier)
      os.Exit(1)
    }
    if target.name != "" {
      fmt.Printf("%v would be %v times the size of current salary %v\n", target.Salary, target.Salary / current.Salary, current.Salary)
      os.Exit(1)
    }
    fmt.Fprintf(os.Stderr, "Not enough data supplied. %v\n", "Missing either a target salary (--target or -t) or multiplier (--multiplier or -m).")
    os.Exit(1)
  }
  fmt.Fprintf(os.Stderr, "Not enough data supplied. Missing a current salary (--target or -t) %v\n", "")
  os.Exit(1)
}

func parseArg(arg string) (ArgResult, error) {
  result := strings.Split(arg, "=")
  if len(result) < 2 {
    fmt.Fprintf(os.Stderr, "Not enough data for %v\n", result[0])
    os.Exit(1)
  }
  name :=result[0]
  value := parseArgValue(result[1])
  if name == multiplierArg || name == mArg {
    multiplier := Multiplier(value)
    return ArgResult{
      name: name,
      Salary: 0,
      Multiplier: multiplier,
    }, nil
  }
  if name == currentArg || name == cArg {
    current := Salary(value)
    return ArgResult{
      name: name,
      Multiplier: 0,
      Salary: current,
    }, nil
  }
  if name == targetArg || name == tArg {
    target := Salary(value)
    return ArgResult{
      name: name,
      Multiplier: 0,
      Salary: target,
    }, nil
  }
  return ArgResult{
    name: "",
    Multiplier: 0,
    Salary: 0,
  }, errors.New(name)
}

func parseArgValue(result string) float64 {
  value, err := strconv.ParseFloat(result, 64)
  if err != nil {
    fmt.Fprintf(os.Stderr, "value %v is not a number\n", err)
    os.Exit(1)
  }
  return value
}
