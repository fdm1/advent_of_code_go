package aoc_utils

import "testing"

func TestReadFileListOfInt(t *testing.T) {
  input_file := TestInput("file_reader_tests/one_int_per_row")
  result := ReadFileListOfInt(input_file)
  if len(result) != 3 {
    t.Fatalf("result has wrong length (%v)", len(result))
  }

  expected_result := []int{}
  expected_result = append(expected_result, 123)
  expected_result = append(expected_result, 456)
  expected_result = append(expected_result, 789)

  for i := range result {
    if result[i] != expected_result[i] {
      t.Fatalf("contents don't match (%v vs %v)", result, expected_result)
    }
  }

}
