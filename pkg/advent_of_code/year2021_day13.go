package advent_of_code

import (
  "fmt"
  "strings"
  "strconv"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)


func Year2021Day13Part1(cache bool) string {
  content := aoc_utils.DownloadInput(2021, 13, cache)
  input := aoc_utils.InputToSlice(content)
  coordinates := map[string]bool{}
  foldInstructions := []string{}
  reachedFolds := false
  maxHeight := 0
  maxWidth := 0
  for _, row := range input {
    if row == "" { reachedFolds = true }
    if !reachedFolds {
      x, y := Y21D13CoordinatesToInts(row)
      if x >= maxWidth { maxWidth = x }
      if y >= maxHeight { maxHeight = y }
      coordinates[row] = true
    } else {
      if row != "" { foldInstructions = append(foldInstructions, strings.Split(row, " ")[2]) }
    }
  }

  coordinates, maxWidth, maxHeight = Y21D13Fold(coordinates, maxWidth, maxHeight, foldInstructions[0])

  return fmt.Sprintf("%v", len(coordinates))
}

func Year2021Day13Part2(cache bool) string {
  content := aoc_utils.DownloadInput(2021, 13, cache)
  input := aoc_utils.InputToSlice(content)
  coordinates := map[string]bool{}
  foldInstructions := []string{}
  reachedFolds := false
  maxHeight := 0
  maxWidth := 0
  for _, row := range input {
    if row == "" { reachedFolds = true }
    if !reachedFolds {
      x, y := Y21D13CoordinatesToInts(row)
      if x >= maxWidth { maxWidth = x }
      if y >= maxHeight { maxHeight = y }
      coordinates[row] = true
    } else {
      if row != "" { foldInstructions = append(foldInstructions, strings.Split(row, " ")[2]) }
    }
  }

  for _, instruction := range foldInstructions {
    coordinates, maxWidth, maxHeight = Y21D13Fold(coordinates, maxWidth, maxHeight, instruction)
  }

  result := "\n"
  for y := 0; y <= maxHeight; y += 1 {
    for x := 0; x <= maxWidth; x += 1 {
      pointFound := false
      for point, _ := range coordinates {
        if point == Y21D13IntsToCoordinates(x, y) {
          pointFound = true
        }
      }
      if pointFound {
        result += "#"
      } else {
        result += " "
      }
    }
    result += "\n"
  }
  return result
}

func Y21D13Fold(coordinateMap map[string]bool, maxWidth int, maxHeight int, foldInstruction string) (map[string]bool, int, int) {
  splitInstruction := strings.Split(foldInstruction, "=")
  axis := splitInstruction[0]
  lineN, _ := strconv.Atoi(splitInstruction[1])
  newCoordinateMap := map[string]bool{}
  newMaxHeight := 0
  newMaxWidth := 0
  if axis == "y" {
    newMaxWidth = maxWidth
    newMaxHeight = maxHeight - lineN - 1
  } else {
    newMaxHeight = maxHeight
    newMaxWidth = maxWidth - lineN - 1
  }

  for coordinates, _ := range coordinateMap {
    x, y := Y21D13CoordinatesToInts(coordinates)
    if axis == "y" {
      if y < lineN {
        newCoordinateMap[coordinates] = true
      } else {
        newY := maxHeight - y
        newCoordinateMap[Y21D13IntsToCoordinates(x, newY)] = true
      }
    } else {
      if x < lineN {
        newCoordinateMap[coordinates] = true
      } else {
        newX := maxWidth - x
        newCoordinateMap[Y21D13IntsToCoordinates(newX, y)] = true
      }
    }
  }
  return newCoordinateMap, newMaxWidth, newMaxHeight

}

func Y21D13IntsToCoordinates(x int, y int) string {
  return fmt.Sprintf("%v,%v", x, y)
}

func Y21D13CoordinatesToInts(coordinates string) (int, int) {
  ints := strings.Split(coordinates, ",")
  x, _ := strconv.Atoi(ints[0])
  y, _ := strconv.Atoi(ints[1])
  return x, y
}
