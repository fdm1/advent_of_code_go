package y21d15

import (
  "fmt"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)


func Part1(cache bool) string {
  content := aoc_utils.DownloadInput(2021, 15, cache)
  grid := aoc_utils.InputToSliceOfIntSlices(content)
  maxRisk := InitialRisk(grid)
  endpointRisks := map[string]int{}
  endpointRisks["0,1"] = grid[0][1]
  endpointRisks["1,0"] = grid[1][0]

  running := true
  for running {
    endpointRisks = Move(endpointRisks, grid, maxRisk)
    if len(endpointRisks) == 1 { running = false }
  }

  risk := 0
  for _, val := range endpointRisks {
    risk = val
  }

  return fmt.Sprintf("%v", risk)
}

func Part2(cache bool) string {
  content := aoc_utils.DownloadInput(2021, 15, cache)
  input := aoc_utils.InputToSlice(content)
  return fmt.Sprintf("%v", input)
}

// Get a value as a baseline
func InitialRisk(grid [][]int) int {
  risk := 0
  for i, val := range grid[0] {
    if i != 0 { risk += val }
  }
  for i, row := range grid {
    if i != 0 { risk += row[len(row) - 1] }
  }

  return risk
}

func Move(pathEndpoints map[string]int, grid [][]int, maxRisk int) map[string]int {
  newEndpoints := map[string]int{}
  for endpoint, currentRisk := range pathEndpoints {
    x, y := aoc_utils.GridPointToInts(endpoint)
    if x < len(grid[0]) - 1 {
      newPoint := aoc_utils.GridIntsToPoint(x+1, y)
      newRisk := currentRisk + grid[x+1][y]
      if newEndpoints[newPoint] == 0 && newRisk < maxRisk {
        newEndpoints[newPoint] = newRisk
      } else if newEndpoints[newPoint] > 0 && newRisk < newEndpoints[newPoint] {
        newEndpoints[newPoint] = newRisk
      }
    }
    if y < len(grid) - 1 {
      newPoint := aoc_utils.GridIntsToPoint(x, y+1)
      newRisk := currentRisk + grid[x][y+1]
      if newEndpoints[newPoint] == 0 && newRisk < maxRisk {
        newEndpoints[newPoint] = newRisk
      } else if newEndpoints[newPoint] > 0 && newRisk < newEndpoints[newPoint] {
        newEndpoints[newPoint] = newRisk
      }
    }
  }
  return newEndpoints
}
