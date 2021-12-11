package advent_of_code

import (
  "testing"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

func TestYear2021Day1Part1(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 1)
  result := Year2021Day1Part1(false)
  expectedResult := "7"

  if result != expectedResult {
    t.Fatalf("actual result %v != expected result %v", result, expectedResult)
  }
}

func TestYear2021Day1Part2(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 1)
  result := Year2021Day1Part2(false)
  expectedResult := "5"

  if result != expectedResult {
    t.Fatalf("actual result %v != expected result %v", result, expectedResult)
  }
}
