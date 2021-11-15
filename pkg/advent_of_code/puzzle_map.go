package advent_of_code

import "fmt"

func PuzzleMap(year string, day string, part string) func(string) int {
  puzzleFunctions := map[string]func(string) int {
    "2020-1-1": Year2020Day1Part1,
  }

  return puzzleFunctions[fmt.Sprintf("%v-%v-%v", year, day, part)]
}
