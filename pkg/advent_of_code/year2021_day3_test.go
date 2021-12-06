package advent_of_code

import (
  "fmt"
  "testing"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

func TestYear2021Day3Part1(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 3)
  result := Year2021Day3Part1(false)
  expected_result := "198"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestYear2021Day3Part2(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 3)
  result := Year2021Day3Part2(false)
  expected_result := "230"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestYear2021Day3Part2FilterTakeMore(t *testing.T) {
  filename := fmt.Sprintf("puzzle_tests/%v_%v", 2021, 3)
  content := aoc_utils.ReadFile(aoc_utils.TestInput(filename))
  input := aoc_utils.InputToSlice(content)

  expected_result := []string{}
  expected_result = append(expected_result, "11110", "10110", "10111", "10101", "11100", "10000", "11001")

  result := Year2021Day3Part2Filter(input, true, 0, "1")

  for i := range result {
    if result[i] != expected_result[i] {
      t.Fatalf("contents don't match (%v vs %v)", result, expected_result)
    }
  }
}

func TestYear2021Day3Part2FilterTakeLess(t *testing.T) {
  filename := fmt.Sprintf("puzzle_tests/%v_%v", 2021, 3)
  content := aoc_utils.ReadFile(aoc_utils.TestInput(filename))
  input := aoc_utils.InputToSlice(content)

  expected_result := []string{}
  expected_result = append(expected_result, "00100", "01111", "00111", "00010", "01010")

  result := Year2021Day3Part2Filter(input, false, 0, "1")

  for i := range result {
    if result[i] != expected_result[i] {
      t.Fatalf("contents don't match (%v vs %v)", result, expected_result)
    }
  }
}


func TestYear2021Day3Part2FilterOneValueLeft(t *testing.T) {
  input := []string{}
  input = append(input, "foobar")

  expected_result := []string{}
  expected_result = append(expected_result, "foobar")

  result := Year2021Day3Part2Filter(input, true, 0, "1")

  for i := range result {
    if result[i] != expected_result[i] {
      t.Fatalf("contents don't match (%v vs %v)", result, expected_result)
    }
  }
}

func TestYear2021Day3Part2FilterTie(t *testing.T) {
  input := []string{}
  input = append(input, "10110", "10111")

  expected_result := []string{}
  expected_result = append(expected_result, "10111")

  result := Year2021Day3Part2Filter(input, true, 4, "1")

  for i := range result {
    if result[i] != expected_result[i] {
      t.Fatalf("contents don't match (%v vs %v)", result, expected_result)
    }
  }
}
