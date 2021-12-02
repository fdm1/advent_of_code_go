package advent_of_code

import (
  "testing"

  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

func TestYear2021Day2Part1(t *testing.T) {
  filename := aoc_utils.TestInput("puzzle_tests/year_2021_day_2")
  result := Year2021Day2Part1(filename)
  expected_result := "150"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestYear2021Day2Part2(t *testing.T) {
  filename := aoc_utils.TestInput("puzzle_tests/year_2021_day_2")
  result := Year2021Day2Part2(filename)
  expected_result := "900"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}
