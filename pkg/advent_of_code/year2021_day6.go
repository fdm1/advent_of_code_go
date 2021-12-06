package advent_of_code

import (
  "fmt"
  "strings"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)


func Year2021Day6Part1(cache bool) string {
  content := aoc_utils.DownloadInput(2021, 6, cache)
  fish := aoc_utils.StringSliceToIntSlice(strings.Split(aoc_utils.InputToSlice(content)[0], ","))

  fishCounters := make(map[int]int)
  for i := 0; i < 8; i++ {
    fishCounters[i] = 0
  }
  for _, val := range fish {
    fishCounters[val] += 1
  }
  for i := 0; i < 80; i++ {
    fishCounters = Y21D6IterateFishOnce(fishCounters)
  }
  result := 0
  for _, val := range fishCounters {
    result += val
  }
  return fmt.Sprintf("%v", result)
}

func Year2021Day6Part2(cache bool) string {
  content := aoc_utils.DownloadInput(2021, 6, cache)
  fish := aoc_utils.StringSliceToIntSlice(strings.Split(aoc_utils.InputToSlice(content)[0], ","))

  fishCounters := make(map[int]int)
  for i := 0; i <= 8; i++ {
    fishCounters[i] = 0
  }
  for _, val := range fish {
    fishCounters[val] += 1
  }
  for i := 0; i < 256; i++ {
    fishCounters = Y21D6IterateFishOnce(fishCounters)
  }
  result := 0
  for _, val := range fishCounters {
    result += val
  }
  return fmt.Sprintf("%v", result)
}


func Y21D6IterateFishOnce(fishCounters map[int]int) map[int]int {
  newFishCounters := make(map[int]int)

  for num, count := range fishCounters {
    if num == 0 {
      newFishCounters[8] += count
      newFishCounters[6] += count
    } else {
      newFishCounters[num - 1] += count
    }
  }

  return newFishCounters
}
