package advent_of_code

import (
  "strconv"

  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)


func Year2020Day1Part1(cache bool) string {
  target := 2020
  content := aoc_utils.DownloadInput(2020, 1, cache)
  inputInts := aoc_utils.InputToIntSlice(content)
  filteredInts := Year2020Day1FilterInts(target, inputInts)

  for i, valI := range filteredInts {
    for _, valJ := range filteredInts[i:len(filteredInts)] {
      if valI + valJ == target {
        return strconv.Itoa(valI * valJ)
      }
    }
  }

  return "no answer"
}

func Year2020Day1Part2(cache bool) string {
  target := 2020
  content := aoc_utils.DownloadInput(2020, 1, cache)
  inputInts := aoc_utils.InputToIntSlice(content)
  filteredInts := Year2020Day1FilterInts(target, inputInts)

  for _, valI := range filteredInts {
    for j, valJ := range filteredInts {
      for _, valK := range filteredInts[j:len(filteredInts)] {
        if valI != valJ {
          if valI + valJ + valK == target {
            return strconv.Itoa(valI * valJ * valK)
          }
        }
      }
    }
  }

  return "no answer"
}

func Year2020Day1FilterInts(maxInt int, allInts []int) []int {
  result := []int{}
  for i := range allInts {
    if allInts[i] < maxInt {
      result = append(result, allInts[i])
    }
  }
  return result
}
