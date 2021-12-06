package advent_of_code

import (
  "testing"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

func TestYear2021Day6Part1(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 6)
  result := Year2021Day6Part1(false)
  expected_result := "5934"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestYear2021Day6Part2(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 6)
  result := Year2021Day6Part2(false)
  expected_result := "26984457539"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}