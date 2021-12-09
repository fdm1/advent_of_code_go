package advent_of_code

import (
  "testing"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

func TestYear2021Day9Part1(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 9)
  result := Year2021Day9Part1(false)
  expected_result := "15"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestYear2021Day9Part2(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 9)
  result := Year2021Day9Part2(false)
  expected_result := "1134"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestY21D9CalculateRisk(t *testing.T) {
  lowPoints := map[string]int{}
  lowPoints["foo"] = 1
  lowPoints["bar"] = 0
  lowPoints["baz"] = 5
  lowPoints["foobar"] = 5
  result := Y21D9CalculateRisk(lowPoints)
  expected_result := 15

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestY21D9FindLowPointsForRow(t *testing.T) {
  row0 := []int{3, 2, 3, 6, 9, 8, 9, 8, 7, 6, 6, 7, 6, 7, 8, 9, 5, 4, 3, 9, 9, 7, 6, 3, 4, 6, 7, 8, 9, 4, 3, 4, 5, 9, 8, 6, 9, 8, 7, 7, 9, 9, 7, 6, 8, 9, 3, 6, 7, 8, 9, 9, 8, 5, 4, 3, 2, 3, 4, 5, 9, 8, 5, 4, 3, 4, 6, 9, 6, 5, 9, 3, 9, 8, 7, 6, 7, 8, 4, 5, 6, 9, 7, 8, 9, 8, 6, 7, 9, 9, 7, 9, 9, 8, 7, 6, 7, 5, 6, 7}
  row1 := []int{2, 1, 3, 5, 6, 7, 8, 9, 8, 7, 7, 9, 8, 9, 9, 4, 3, 1, 2, 4, 9, 8, 5, 4, 5, 6, 8, 9, 6, 4, 2, 3, 9, 9, 7, 5, 9, 8, 6, 6, 8, 5, 9, 7, 8, 9, 2, 4, 5, 6, 9, 9, 9, 7, 8, 7, 5, 4, 6, 6, 7, 9, 6, 5, 4, 5, 7, 8, 9, 9, 8, 4, 5, 9, 8, 9, 7, 6, 5, 6, 7, 9, 6, 7, 9, 8, 7, 8, 9, 5, 6, 9, 9, 9, 5, 4, 3, 4, 5, 8}
  row2 := []int{3, 4, 4, 6, 8, 8, 9, 3, 9, 8, 9, 1, 9, 8, 6, 5, 3, 2, 4, 9, 8, 7, 6, 5, 6, 7, 9, 6, 5, 3, 0, 9, 8, 7, 6, 4, 9, 7, 5, 4, 3, 4, 9, 8, 9, 0, 1, 3, 4, 5, 8, 9, 9, 8, 9, 7, 6, 5, 7, 7, 8, 9, 7, 6, 5, 7, 9, 9, 9, 8, 7, 5, 6, 7, 9, 9, 8, 9, 7, 7, 8, 9, 5, 4, 5, 9, 8, 9, 3, 4, 9, 8, 9, 8, 7, 4, 2, 4, 5, 5}
  intRows := [][]int{row0, row1, row2}
  result := Y21D9FindLowPointsForRow(1, intRows)
  expectedResult := map[string]int{}
  expectedResult["1,1"] = 1
  expectedResult["1,17"] = 1

  if len(result) != len(expectedResult) {
    t.Fatalf("contents don't match (%v vs %v)", result, expectedResult)
  }

  for i := range result {
    if result[i] != expectedResult[i] {
      t.Fatalf("contents don't match (%v vs %v)", result, expectedResult)
    }
  }
}

func TestY21D9PointAboveOrBelowIsPartOfBasin(t *testing.T) {
  row1 := []int{3, 9, 8, 7, 8, 9, 4, 9, 2, 1}
  point := "0,6"
  val := 3
  result := Y21D9PointAboveOrBelowIsPartOfBasin(point, val, row1)
  expectedResult := true

  if result != expectedResult {
    t.Fatalf("contents don't match (%v vs %v)", result, expectedResult)
  }
}
