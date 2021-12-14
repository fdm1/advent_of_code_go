package advent_of_code

import (
  "testing"
  "strings"
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

// func TestYear2021Day14Part2(t *testing.T) {
//   aoc_utils.MockAdventOfCodeInput(2021, 14)
//   result := Year2021Day14Part2(false)
//   expectedResult := "2188189693529"
//
//   if result != expectedResult {
//     t.Fatalf("actual result %v != expected result %v", result, expectedResult)
//   }
// }

func TestY21D14Iterate(t *testing.T) {
  input := strings.Split("NNCB", "")

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
  result := strings.Join(Y21D14Iterate(input, steps), "")
  expectedResult := "NCNBCHB"

  if result != expectedResult {
    t.Fatalf("actual result %v != expected result %v", result, expectedResult)
  }
}
