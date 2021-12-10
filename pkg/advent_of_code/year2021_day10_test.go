package advent_of_code

import (
  "testing"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

func TestYear2021Day10Part1(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 10)
  result := Year2021Day10Part1(false)
  expected_result := "26397"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestYear2021Day10Part2(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 10)
  result := Year2021Day10Part2(false)
  expected_result := "288957"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestY21D10FindIllegalChar(t *testing.T) {
  result, _ := Y21D10FindIllegalChar("{([(<{}[<>[]}>{[]{[(<()>")
  expected_result := "}"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestY21D10ScoreForAutocomplete(t *testing.T) {
  result := Y21D10ScoreForAutocomplete([]string{"[", "(", "{", "(", "[", "[", "{", "{"})
  expected_result := 288957

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}
