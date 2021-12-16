package y2021d16

import (
  "testing"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

func TestPart1(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 16)
  result := Part1(false)
  expectedResult := "31"

  if result != expectedResult {
    t.Fatalf("actual result %v != expected result %v", result, expectedResult)
  }
}

// func TestPart2(t *testing.T) {
//   aoc_utils.MockAdventOfCodeInput(2021, 16)
//   result := Part2(false)
//   expectedResult := "part2_answer"
//
//   if result != expectedResult {
//     t.Fatalf("actual result %v != expected result %v", result, expectedResult)
//   }
// }
