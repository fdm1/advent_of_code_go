package advent_of_code

import (
  "strconv"
  "strings"

  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

func Year2020Day2Part1(filename string) string {
  inputLines := aoc_utils.ReadFileLines(filename)
  result := 0
  if len(inputLines) > 0 {
    result = 0
  }

  for _, inputLine := range inputLines[0:len(inputLines)-1] {
    minCount, maxCount, char, password := ParseY20D2InputLine(inputLine)
    pwLength := len(password)
    charCount := pwLength - len(strings.Replace(password, char, "", -1))
    if minCount <= charCount && charCount <= maxCount {
      result = result + 1
    }
  }
  if result == 0 {
    return "no answer"
  }
  return strconv.Itoa(result)
}

func ParseY20D2InputLine(input string) (int, int, string, string) {
  spaceSplit := strings.Split(input, " ")

  counts := strings.Split(spaceSplit[0], "-")
  char := strings.Replace(spaceSplit[1], ":", "", 1)
  password := spaceSplit[2]

  minCount, _ := strconv.Atoi(counts[0])
  maxCount, _ := strconv.Atoi(counts[1])
  return minCount, maxCount, char, password
}
