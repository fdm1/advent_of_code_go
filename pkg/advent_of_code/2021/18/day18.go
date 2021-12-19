package y2021d18

import (
  "fmt"
  "strings"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)


func Part1(cache bool) string {
  content := aoc_utils.DownloadInput(2021, 18, cache)
  input := aoc_utils.InputToSlice(content)
  fmt.Println(input[0])
  groups := GetGroups(input[0])
  neeedToReduceSubGroups = len(groups) > 1 || !GroupReadyToReduce[groups[0]]
  for neeedToReduceSubGroups {
  }

  for len(groups) > 1 {
    sub
  }
  fmt.Println(GroupReadyToReduce("[[0,1],[2,3]]"))
  // for _, val := range strings.Split(input[0], "],[") {
  //   fmt.Println(val)
  // }
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

func GetGroups(inputString string) []string {
  groups := []string{}
  pointer := 1
  chars := strings.Split(inputString, "")
  openCount := 0
  for i, val := range chars[pointer:len(chars)] {
    if val == "[" {
      if openCount == 0 { pointer = i }
      openCount++
    }
    if val == "]" {
      openCount--
      if openCount == 0 {
        groups = append(groups, strings.Join(chars[pointer+1:i+2], ""))
      }
    }
  }
  for i, group := range groups {
    fmt.Printf("group %v: %v\n", i, group)
  }
  return groups
}

func GroupReadyToReduce(group string) bool {
  chars := strings.Split(group, "")
  openCount := 0
  for _, val := range chars {
    if val == "[" { openCount++ }
    if val == "]" { openCount-- }
    if openCount > 2 { return false }
  }
  return true
}

func GroupShouldExplode(group string) bool {
  chars := strings.Split(group, "")
  openCount := 0
  for _, val := range chars {
    if val == "[" { openCount++ }
    if val == "]" { openCount-- }
    if openCount > 4 { return true }
  }
  return false
}
