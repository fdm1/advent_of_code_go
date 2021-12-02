package aoc_utils

import (
  "strings"
  "strconv"
)

func InputToIntSlice(content string) []int {
  lines := InputToSlice(content)
  ints := make([]int, len(lines))
  for i, line := range lines {
    if line != "" {
      num, _ := strconv.Atoi(line)
      ints[i] = num
    }
  }
  return ints
}

func InputToSlice(content string) []string {
  lines := strings.Split(string(content), "\n")
  return lines[0:len(lines)-1]
}
