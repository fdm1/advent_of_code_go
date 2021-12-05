package advent_of_code

import "fmt"

func PuzzleMap(year string, day string, part string) func(bool) string {
  puzzleFunctions := map[string]func(bool) string {
    "2020-1-1": Year2020Day1Part1,
    "2020-1-2": Year2020Day1Part2,
    "2020-2-1": Year2020Day2Part1,
    "2020-2-2": Year2020Day2Part2,
    "2021-1-1": Year2021Day1Part1,
    "2021-1-2": Year2021Day1Part2,
    "2021-2-1": Year2021Day2Part1,
    "2021-2-2": Year2021Day2Part2,
    "2021-3-1": Year2021Day3Part1,
    "2021-3-2": Year2021Day3Part2,
    "2021-4-1": Year2021Day4Part1,
    "2021-4-2": Year2021Day4Part2,
    "2021-5-1": Year2021Day5Part1,
    "2021-5-2": Year2021Day5Part2,
  }

  return puzzleFunctions[fmt.Sprintf("%v-%v-%v", year, day, part)]
}
