package advent_of_code

import (
  "testing"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

func TestYearYEARDayDAYPart1(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(YEAR, DAY)
  result := YearYEARDayDAYPart1(false)
  expectedResult := "part1_answer"

  if result != expectedResult {
    t.Fatalf("actual result %v != expected result %v", result, expectedResult)
  }
}

// func TestYearYEARDayDAYPart2(t *testing.T) {
//   aoc_utils.MockAdventOfCodeInput(YEAR, DAY)
//   result := YearYEARDayDAYPart2(false)
//   expectedResult := "part2_answer"
//
//   if result != expectedResult {
//     t.Fatalf("actual result %v != expected result %v", result, expectedResult)
//   }
// }
