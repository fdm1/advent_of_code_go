package advent_of_code

import (
  "testing"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

func TestYear2021Day13Part1(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 13)
  result := Year2021Day13Part1(false)
  expectedResult := "17"

  if result != expectedResult {
    t.Fatalf("actual result %v != expected result %v", result, expectedResult)
  }
}

func TestYear2021Day13Part2(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 13)
  result := Year2021Day13Part2(false)
  expectedResult := "\n#####\n#   #\n#   #\n#   #\n#####\n     \n     \n"

  if result != expectedResult {
    t.Fatalf("actual result %v != expected result %v", result, expectedResult)
  }
}
