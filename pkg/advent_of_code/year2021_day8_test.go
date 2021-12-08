package advent_of_code

import (
  "testing"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

func TestYear2021Day8Part1(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 8)
  result := Year2021Day8Part1(false)
  expected_result := "26"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestYear2021Day8Part2(t *testing.T) {
  aoc_utils.MockAdventOfCodeInput(2021, 8)
  result := Year2021Day8Part2(false)
  expected_result := "61229"

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestY21D8DecodeRow(t *testing.T) {
  row := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"
  result := Y21D8DecodeRow(row)
  expected_result := 5353

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestY21D8DecodeRowMore(t *testing.T) {
  row := "bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef"
  result := Y21D8DecodeRow(row)
  expected_result := 1625

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}

func TestY21D8DecodeRowEarlyReturn(t *testing.T) {
  row := "stuf | aa abc cadb cdfebaf"
  result := Y21D8DecodeRow(row)
  expected_result := 1748

  if result != expected_result {
    t.Fatalf("actual result %v != expected result %v", result, expected_result)
  }
}
