package y2021d18

import (
  "fmt"
  "strings"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)


func Part1(cache bool) string {
  content := aoc_utils.DownloadInput(2021, 18, cache)
  input := aoc_utils.InputToSlice(content)
  return fmt.Sprintf("%v", strings.Split(input[0], "],["))
}

func Part2(cache bool) string {
  content := aoc_utils.DownloadInput(2021, 18, cache)
  input := aoc_utils.InputToSlice(content)
  return fmt.Sprintf("%v", input)
}

func PairMagnitute(pair []int) int {
  return 3 * pair[0] + 2 * pair[1]
}

func ReduceRow(inputString string) (string, bool) {
  // pointer := 0
  // openC
  // chars := strings.Split(inputString, "")
  // for i, val := range chars {
  //   if char == "[" {
  //     pointer = i
  //   op
  // }
  //
  return "7", true
}
// func ReduceRow(inputRow string) []int {
//   
//   
// }
//
