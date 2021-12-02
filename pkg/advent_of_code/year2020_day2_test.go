package advent_of_code

import (
  "testing"
)

func TestYear2020Day2Part1(t *testing.T) {
  MockAdventOfCodeInput(2020, 2)
  result := Year2020Day2Part1()
  expected_result := "2"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestYear2020Day2Part2(t *testing.T) {
  MockAdventOfCodeInput(2020, 2)
  result := Year2020Day2Part2()
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
