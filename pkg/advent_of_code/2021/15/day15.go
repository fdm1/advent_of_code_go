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
  grid := aoc_utils.InputToSliceOfIntSlices(content)
  initialRows := len(grid)
  initialColumns := len(grid[0])
  grid = ExpandGrid(grid, initialRows, initialColumns)
  grid = ExpandGrid(grid, initialRows, initialColumns)
  grid = ExpandGrid(grid, initialRows, initialColumns)
  grid = ExpandGrid(grid, initialRows, initialColumns)

  maxRisk := InitialRisk(grid)
  endpointRisks := map[string]int{}
  endpointRisks["0,1"] = grid[0][1]
  endpointRisks["1,0"] = grid[1][0]

  running := true
  risk := 0

  for running {
    endpointRisks = Move(endpointRisks, grid, maxRisk)
    if len(endpointRisks) == 1 {
      for point, riskVal := range endpointRisks {
        x, y := aoc_utils.GridPointToInts(point)
        if x == len(grid[0]) - 1 && y == len(grid) - 1 {
          risk = riskVal
          running = false
        }
      }
    }
  }

  return fmt.Sprintf("%v", risk)
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
      if newEndpoints[newPoint] == 0 && newRisk <= maxRisk {
        newEndpoints[newPoint] = newRisk
      } else if newEndpoints[newPoint] > 0 && newRisk < newEndpoints[newPoint] {
        newEndpoints[newPoint] = newRisk
      }
    }
    if y < len(grid) - 1 {
      newPoint := aoc_utils.GridIntsToPoint(x, y+1)
      newRisk := currentRisk + grid[x][y+1]
      if newEndpoints[newPoint] == 0 && newRisk <= maxRisk {
        newEndpoints[newPoint] = newRisk
      } else if newEndpoints[newPoint] > 0 && newRisk < newEndpoints[newPoint] {
        newEndpoints[newPoint] = newRisk
      }
    }
  }
  return newEndpoints
}

func ExpandGrid(grid [][]int, rows int, columns int) [][]int {
  newGrid := [][]int{}
  for row, rowVals := range grid {
    newRow := grid[row]
    for c := 0; c < columns; c++ {
      oldVal := rowVals[len(rowVals) - (columns - c)]
      if oldVal == 9 {
        newRow = append(newRow, 1)
      } else {
        newRow = append(newRow, oldVal + 1)
      }
    }
    newGrid = append(newGrid, newRow)
  }
  for r := 0; r < rows; r++ {
    newRow := []int{}
    for _, val := range newGrid[len(newGrid) - rows] { newRow = append(newRow, val) }
    for i, val := range newRow {
      if val == 9 {
        newRow[i] = 1
      } else {
        newRow[i] = val + 1
      }
    }
    newGrid = append(newGrid, newRow)
  }

  return newGrid
}
