package advent_of_code

import (
  "testing"

  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

func TestYear2020Day2Part1(t *testing.T) {
  filename := aoc_utils.TestInput("puzzle_tests/year_2020_day_2")
  result := Year2020Day2Part1(filename)
  expected_result := "2"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestYear2020Day2Part2(t *testing.T) {
  filename := aoc_utils.TestInput("puzzle_tests/year_2020_day_2")
  result := Year2020Day2Part2(filename)
  expected_result := "1"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestParseY20D2InputLine(t *testing.T) {
  inputLine := "1-3 b: cdefg"

  minCount, maxCount, char, password := ParseY20D2InputLine(inputLine)

  if minCount != 1 {
    t.Fatalf("actual minCount %v != 1", minCount)
  }

  if maxCount != 3 {
    t.Fatalf("actual maxCount %v != 1", maxCount)
  }

  if char != "b" {
    t.Fatalf("actual char %v != 1", char)
  }

  if password != "cdefg" {
    t.Fatalf("actual password %v != 1", password)
  }


}
