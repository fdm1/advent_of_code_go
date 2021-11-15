package advent_of_code

import "fmt"

func PuzzleMap(year string, day string, part string) func(string) string {
  puzzleFunctions := map[string]func(string) string {
    "2020-1-1": Year2020Day1Part1,
    "2020-1-2": Year2020Day1Part2,
  }

  return puzzleFunctions[fmt.Sprintf("%v-%v-%v", year, day, part)]
}
