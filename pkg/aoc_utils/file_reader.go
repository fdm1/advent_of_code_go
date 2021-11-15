package aoc_utils

import (
  "os"
  "strconv"
  "strings"
)

func ReadFileListOfInt(filename string) []int {
  lines := ReadFileLines(filename)
  ints := make([]int, len(lines))
  for i, line := range lines {
    if line != "" {
      num, _ := strconv.Atoi(line)
      ints[i] = num
    }
  }
  return ints
}

func ReadFileLines(filename string) []string {
  content, err := os.ReadFile(filename)
  if err != nil {
      //Do something
  }
  lines := strings.Split(string(content), "\n")
  return lines[0:len(lines)-1]
}
