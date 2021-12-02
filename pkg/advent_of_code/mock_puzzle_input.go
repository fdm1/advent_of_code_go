package advent_of_code

import (
  "fmt"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
  "github.com/jarcoal/httpmock"
)

func MockAdventOfCodeInput(year int, day int) {
  httpmock.Activate()
  // defer httpmock.DeactivateAndReset()

  url := fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", year, day)
  filename := fmt.Sprintf("puzzle_tests/%v_%v", year, day)
  content := aoc_utils.ReadFile(aoc_utils.TestInput(filename))

  httpmock.RegisterResponder("GET", url,
    httpmock.NewStringResponder(200, content))
}
