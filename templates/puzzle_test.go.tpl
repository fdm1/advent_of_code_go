package advent_of_code

import (
  "testing"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

func TestYearYEARDayDAYPart1(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(YEAR, DAY)
  result := YearYEARDayDAYPart1(false)
  expected_result := "part1_answer"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestYearYEARDayDAYPart2(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(YEAR, DAY)
  result := YearYEARDayDAYPart2(false)
  expected_result := "part2_answer"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}
