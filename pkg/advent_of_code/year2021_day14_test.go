package advent_of_code

import (
  "testing"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

func TestYear2021Day14Part1(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 14)
  result := Year2021Day14Part1(false)
  expectedResult := "1588"

  if result != expectedResult {
    t.Fatalf("actual result %v != expected result %v", result, expectedResult)
  }
}

func TestYear2021Day14Part2(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 14)
  result := Year2021Day14Part2(false)
  expectedResult := "2188189693529"

  if result != expectedResult {
    t.Fatalf("actual result %v != expected result %v", result, expectedResult)
  }
}

func TestY21D14Iterate(t *testing.T) {
  input := map[string]int{}
  input["CB"] = 1
  input["NC"] = 1
  input["NN"] = 1

  steps := [][]string{
    []string{"CH", "B"},
    []string{"HH", "N"},
    []string{"CB", "H"},
    []string{"NH", "C"},
    []string{"HB", "C"},
    []string{"HC", "B"},
    []string{"HN", "C"},
    []string{"NN", "C"},
    []string{"BH", "H"},
    []string{"NC", "B"},
    []string{"NB", "B"},
    []string{"BN", "B"},
    []string{"BB", "N"},
    []string{"BC", "B"},
    []string{"CC", "N"},
    []string{"CN", "C"},
  }
  result := Y21D14Iterate(input, steps)
  expectedResult := map[string]int{}
  expectedResult["NC"] = 1
  expectedResult["CN"] = 1
  expectedResult["NB"] = 1
  expectedResult["BC"] = 1
  expectedResult["CH"] = 1
  expectedResult["HB"] = 1

  for x, _ := range result {
    if result[x] != expectedResult[x] {
      t.Fatalf("actual result %v != expected result %v", result, expectedResult)
    }
  }
  if len(result) != len(expectedResult) {
    t.Fatalf("actual result %v != expected result %v", result, expectedResult)
  }
}
