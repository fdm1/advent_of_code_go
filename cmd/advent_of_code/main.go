package main

import (
  "fmt"

  "github.com/fdm1/advent_of_code_go/pkg/advent_of_code"
)

func main() {
  fmt.Printf("Year: ")
  var year string
  fmt.Scanln(&year)

  fmt.Printf("Day: ")
  var day string
  fmt.Scanln(&day)

  fmt.Printf("Part: ")
  var part string
  fmt.Scanln(&part)


  function := advent_of_code.PuzzleMap(year, day, part)
  result := function()
  fmt.Printf("Result: %v\n", result)
}

