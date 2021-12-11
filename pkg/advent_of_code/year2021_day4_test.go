package advent_of_code

import (
  "strings"
  "testing"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

func TestYear2021Day4Part1(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 4)
  result := Year2021Day4Part1(false)
  expectedResult := "4512"

  if result != expectedResult {
    t.Fatalf("actual result %v != expected result %v", result, expectedResult)
  }
}

func TestYear2021Day4Part2(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 4)
  result := Year2021Day4Part2(false)
  expectedResult := "1924"

  if result != expectedResult {
    t.Fatalf("actual result %v != expected result %v", result, expectedResult)
  }
}

func TestY21D4ExtractBingoBoard(t *testing.T) {
  input := []string{}
  input = append(input, "1 2  3", "4 5 6", "7 8 9")
  expectedResult := [][]string{}
  expectedResult = append(expectedResult,
    strings.Split("1 2 3", " "),
    strings.Split("4 5 6", " "),
    strings.Split("7 8 9", " "),
    strings.Split("1 4 7", " "),
    strings.Split("2 5 8", " "),
    strings.Split("3 6 9", " "),
  )
  result := Y21D4ExtractBingoBoard(input)
  for i, valI := range result {
    for j, _ := range valI {
      if result[i][j] != expectedResult[i][j] {
        t.Fatalf("contents don't match (%v vs %v)", result, expectedResult)
      }
    }
  }
}
