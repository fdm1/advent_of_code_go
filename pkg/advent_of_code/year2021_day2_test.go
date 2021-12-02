package advent_of_code

import (
  "testing"

)

func TestYear2021Day2Part1(t *testing.T) {
  MockAdventOfCodeInput(2021, 2)
  result := Year2021Day2Part1()
  expected_result := "150"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestYear2021Day2Part2(t *testing.T) {
  MockAdventOfCodeInput(2021, 2)
  result := Year2021Day2Part2()
  expected_result := "900"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}
