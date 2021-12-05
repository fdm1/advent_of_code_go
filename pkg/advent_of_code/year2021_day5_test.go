package advent_of_code

import (
  "testing"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

func TestYear2021Day5Part1(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 5)
  result := Year2021Day5Part1(false)
  expected_result := "5"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestYear2021Day5Part2(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 5)
  result := Year2021Day5Part2(false)
  expected_result := "12"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}
