package advent_of_code

import (
  "testing"
)

func TestYear2021Day1Part1(t *testing.T) {
  MockAdventOfCodeInput(2021, 1)
  result := Year2021Day1Part1()
  expected_result := "7"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestYear2021Day1Part2(t *testing.T) {
  MockAdventOfCodeInput(2021, 1)
  result := Year2021Day1Part2()
  expected_result := "5"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}
