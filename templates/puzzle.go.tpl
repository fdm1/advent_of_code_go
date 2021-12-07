package advent_of_code

import (
  "fmt"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)


func YearYEARDayDAYPart1(cache bool) string {
  content := aoc_utils.DownloadInput(YEAR, DAY, cache)
  input := aoc_utils.InputToSlice(content)
  return fmt.Sprintf("%v", input)
}

func YearYEARDayDAYPart2(cache bool) string {
  content := aoc_utils.DownloadInput(YEAR, DAY, cache)
  input := aoc_utils.InputToSlice(content)
  return fmt.Sprintf("%v", input)
}
