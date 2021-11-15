package main

import (
  "fmt"

  "github.com/fdm1/advent_of_code_go/pkg/advent_of_code"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

func main() {
  fmt.Println("Year: ")
  var year string
  fmt.Scanln(&year)

  fmt.Println("Day: ")
  var day string
  fmt.Scanln(&day)

  fmt.Println("Part: ")
  var part string
  fmt.Scanln(&part)


  function := advent_of_code.PuzzleMap(year, day, part)
  input := aoc_utils.InputFile(year, day)
  result := function(input)
  fmt.Printf("Result: %v", result)
}

