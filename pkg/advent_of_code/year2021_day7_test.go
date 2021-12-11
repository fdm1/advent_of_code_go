package advent_of_code

import (
  "testing"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

func TestYear2021Day7Part1(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 7)
  result := Year2021Day7Part1(false)
  expectedResult := "37"

  if result != expectedResult {
    t.Fatalf("actual result %v != expected result %v", result, expectedResult)
  }
}

func TestYear2021Day7Part2(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 7)
  result := Year2021Day7Part2(false)
  expectedResult := "168"

  if result != expectedResult {
    t.Fatalf("actual result %v != expected result %v", result, expectedResult)
  }
}

func TestY21D7P2FuelCalculator(t *testing.T) {
  result := Y21D7P2FuelCalculator(11)
  expectedResult := 66
  if result != expectedResult {
    t.Fatalf("actual result %v != expected result %v", result, expectedResult)
  }
}
