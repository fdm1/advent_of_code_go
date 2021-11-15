package advent_of_code

import (
  "strconv"

  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)

const Part1Target = 2020

func Year2020Day1Part1(filename string) string {
  inputInts := aoc_utils.ReadFileListOfInt(filename)
  filteredInts := FilterInts(Part1Target, inputInts)

  for i, valI := range filteredInts {
    for _, valJ := range filteredInts[i:len(filteredInts)] {
      if valI + valJ == Part1Target {
        return strconv.Itoa(valI * valJ)
      }
    }
  }

  return "no answer"
}

func Year2020Day1Part2(filename string) string {
  inputInts := aoc_utils.ReadFileListOfInt(filename)
  filteredInts := FilterInts(Part1Target, inputInts)

  for _, valI := range filteredInts {
    for j, valJ := range filteredInts {
      for _, valK := range filteredInts[j:len(filteredInts)] {
        if valI != valJ {
          if valI + valJ + valK == Part1Target {
            return strconv.Itoa(valI * valJ * valK)
          }
        }
      }
    }
  }

  return "no answer"
}

func FilterInts(maxInt int, allInts []int) []int {
  result := []int{}
  for i := range allInts {
    if allInts[i] < maxInt {
      result = append(result, allInts[i])
    }
  }
  return result
}
