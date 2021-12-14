package advent_of_code

import (
  "fmt"
  "strings"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)


func Year2021Day14Part1(cache bool) string {
  content := aoc_utils.DownloadInput(2021, 14, cache)
  input := aoc_utils.InputToSlice(content)
  inputChars := strings.Split(input[0], "")

  steps := [][]string{}
  for _, row := range input[2:len(input)] {
    steps = append(steps, strings.Split(row, " -> "))
  }

  for i := 0; i < 10; i++ {
    inputChars = Y21D14Iterate(inputChars, steps)
  }

  return fmt.Sprintf("%v", Y21D14ScorePt1(inputChars))
}

func Year2021Day14Part2(cache bool) string {
  content := aoc_utils.DownloadInput(2021, 14, cache)
  input := aoc_utils.InputToSlice(content)
  inputChars := strings.Split(input[0], "")

  steps := [][]string{}
  for _, row := range input[2:len(input)] {
    steps = append(steps, strings.Split(row, " -> "))
  }

  for i := 0; i < 40; i++ {
    inputChars = Y21D14Iterate(inputChars, steps)
  }

  return fmt.Sprintf("%v", Y21D14ScorePt1(inputChars))
}

func Y21D14Iterate(input []string, steps [][]string) []string {
  result := []string{}
  for inputI := 0; inputI < len(input) - 1; inputI++ {
    result = append(result, input[inputI])

    for _, step := range steps {
      if fmt.Sprintf("%v%v", input[inputI], input[inputI + 1]) == step[0] {
        result = append(result, step[1])
        break  // assume no repeats?
      }
    }
  }
  result = append(result, input[len(input) - 1])

  return result
}

func Y21D14ScorePt1(input []string) int {
  charCounts := map[string]int{}
  for _, val := range input {
    charCounts[val]++
  }

  minCount := -1
  maxCount := -1

  for _, count := range charCounts {
    if minCount == -1 {
      minCount = count
      maxCount = count
    } else {
      if count < minCount { minCount = count }
      if count > maxCount { maxCount = count }
    }
  }
  return maxCount - minCount
}
