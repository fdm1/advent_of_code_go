package y2021d16

import (
  "fmt"
  "testing"
  "strings"
  // "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

// func TestPart1(t *testing.T) {
//   aoc_utils.MockAdventOfCodeInput(2021, 16)
//   result := Part1(false)
//   expectedResult := "31"
//
//   if result != expectedResult {
//     t.Fatalf("actual result %v != expected result %v", result, expectedResult)
//   }
// }

// func TestPart2(t *testing.T) {
//   aoc_utils.MockAdventOfCodeInput(2021, 16)
//   result := Part2(false)
//   expectedResult := "part2_answer"
//
//   if result != expectedResult {
//     t.Fatalf("actual result %v != expected result %v", result, expectedResult)
//   }
// }

func TestConvertToBinary(t *testing.T) {
  result := ConvertToBinary("D2FE28")
  expectedResult := strings.Split("110100101111111000101000", "")
  if len(result) != len(expectedResult) {
    t.Fatalf("actual result %v != expected result %v", result, expectedResult)
  }
  for i, _ := range result {
    if result[i] != expectedResult[i] {
      t.Fatalf("actual result %v != expected result %v", result, expectedResult)
    }
  }
}

func TestCalculateSum(t *testing.T) {
  hexStrings := []string{
    "D2FE28",
    "D2FE28D2FE28",
    // "8A004A801A8002F478",
    // "620080001611562C8802118E34",
    // "C0015000016115A2E0802F182340",
    // "A0016C880162017C3686B18A3D4780",
  }
  sums := []int64{
    6,
    12,
    16,
    12,
    23,
    31,
  }
  failed := false
  for i, hexString := range hexStrings {
    result := CalculateSum(hexString, 0, []int64{}, []int64{})
    expectedResult := sums[i]

    if result != expectedResult {
      failed = true
      fmt.Printf("actual result %v != expected result %v (hexString = %v)\n", result, expectedResult, hexString)
    }
  }

  if failed { t.Fatalf("Test failed") }
}
