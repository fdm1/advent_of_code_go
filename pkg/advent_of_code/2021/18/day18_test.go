package y2021d18

import (
  "testing"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

func TestPart1(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 18)
  result := Part1(false)
  expectedResult := "4140"

  if result != expectedResult {
    t.Fatalf("actual result %v != expected result %v", result, expectedResult)
  }
}

// func TestPart2(t *testing.T) {
//   aoc_utils.MockAdventOfCodeInput(2021, 18)
//   result := Part2(false)
//   expectedResult := "part2_answer"
//
//   if result != expectedResult {
//     t.Fatalf("actual result %v != expected result %v", result, expectedResult)
//   }
// }
//
// func TestReduceRow(t *testing.T) {
//   input := "[1,2]"
//   result, _ := ReduceRow(input)
//   expectedResult := "7"
//
//   if result != expectedResult {
//     t.Fatalf("actual result %v != expected result %v", result, expectedResult)
//   }
// }
