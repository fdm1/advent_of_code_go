package aoc_utils

import "testing"

func TestTestInput(t *testing.T) {
  result := TestInput("some-input")
  if result != "../../test_inputs/some-input.txt" {
    t.Fatal("result unexpected")
  }
}

func TestInputFile(t *testing.T) {
  result := InputFile("1234", "123")
  if result != "puzzle_inputs/year_1234_day_123.txt" {
    t.Fatal("result unexpected")
  }
}
