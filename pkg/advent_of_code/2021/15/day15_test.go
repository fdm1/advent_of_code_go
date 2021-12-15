package y21d15

import (
  "testing"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

func TestPart1(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 15)
  result := Part1(false)
  expectedResult := "40"

  if result != expectedResult {
    t.Fatalf("actual result %v != expected result %v", result, expectedResult)
  }
}

// func TestYear2021Day15Part2(t *testing.T) {
//   aoc_utils.MockAdventOfCodeInput(2021, 15)
//   result := Year2021Day15Part2(false)
//   expectedResult := "part2_answer"
//
//   if result != expectedResult {
//     t.Fatalf("actual result %v != expected result %v", result, expectedResult)
//   }
// }

func TestInitialRisk(t *testing.T) {
  grid := [][]int{
    []int{1, 2, 3},
    []int{4, 5, 6},
    []int{7, 8, 9},
    []int{8, 7, 6},
  }
  result := InitialRisk(grid)
  expectedResult := 2 + 3 + 6 + 9 + 6

  if result != expectedResult {
    t.Fatalf("actual result %v != expected result %v", result, expectedResult)
  }
}
