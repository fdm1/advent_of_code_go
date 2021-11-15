package advent_of_code

import (
  "testing"

  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

func TestYear2020Day1Part1(t *testing.T) {
  filename := aoc_utils.TestInput("puzzle_tests/year_2020_day_1")
  result := Year2020Day1Part1(filename)
  expected_result := "514579"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestYear2020Day1Part2(t *testing.T) {
  filename := aoc_utils.TestInput("puzzle_tests/year_2020_day_1")
  result := Year2020Day1Part2(filename)
  expected_result := "241861950"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestFilterInts(t *testing.T) {
  inputs := []int{}
  inputs = append(inputs, 1)
  inputs = append(inputs, 2)
  inputs = append(inputs, 3)

  expected_result := []int{}
  expected_result = append(expected_result, 1)
  expected_result = append(expected_result, 2)


  result := FilterInts(3, inputs)

  for i := range result {
    if result[i] != expected_result[i] {
      t.Fatalf("contents don't match (%v vs %v)", result, expected_result)
    }
  }
}
