package advent_of_code

import (
  "strconv"
  "strings"
  "fmt"

  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)


func Year2021Day3Part1() string {
  content := aoc_utils.DownloadInput(2021, 3)
  input := aoc_utils.InputToSlice(content)
  counts := make([]int, len(strings.Split(input[0], "")))

  for _, val := range input {
    binaryRow := strings.Split(val, "")
    for i, binaryVal := range binaryRow {
      if binaryVal == "1" {
        counts[i] += 1
      }
    }
  }

  epsilonBinary := ""
  gammaBinary := ""

  for _, val := range counts {
    mostAreOne := val >= len(input) / 2
    if mostAreOne {
      epsilonBinary += "0"
      gammaBinary += "1"
    } else {
      epsilonBinary += "1"
      gammaBinary += "0"
    }
  }

  epsilon, _ := strconv.ParseInt(epsilonBinary, 2, 64)
  gamma, _ := strconv.ParseInt(gammaBinary, 2, 64)

  return fmt.Sprintf("%v", epsilon * gamma)
}

func Year2021Day3Part2() string {
  content := aoc_utils.DownloadInput(2021, 3)
  input := aoc_utils.InputToSlice(content)
  counts := make([]int, len(strings.Split(input[0], "")))

  for _, val := range input {
    binaryRow := strings.Split(val, "")
    for i, binaryVal := range binaryRow {
      if binaryVal == "1" {
        counts[i] += 1
      }
    }
  }

  oxygenBinaryCrit := ""
  co2BinaryCrit := ""

  for _, val := range counts {
    mostAreOne := val >= len(input) / 2
    if mostAreOne {
      oxygenBinaryCrit += "0"
      co2BinaryCrit += "1"
    } else {
      oxygenBinaryCrit += "1"
      co2BinaryCrit += "0"
    }
  }

  oxygenBinaryInputs := input
  co2BinaryInputs := input
  for i, _ := range strings.Split(oxygenBinaryCrit, "") {
    oxygenBinaryInputs = Year2021Day3Part2Filter(oxygenBinaryInputs, true, i, "1")
  }
  for i, _ := range strings.Split(co2BinaryCrit, "") {
    co2BinaryInputs = Year2021Day3Part2Filter(co2BinaryInputs, false, i, "0")
  }

  oxygen, _ := strconv.ParseInt(oxygenBinaryInputs[0], 2, 64)
  co2, _ := strconv.ParseInt(co2BinaryInputs[0], 2, 64)

  return fmt.Sprintf("%v", oxygen * co2)
}

func Year2021Day3Part2Filter(input []string, takeMore bool, charIndex int, tieChar string) []string {
  if len(input) == 1 {
    return input
  }
  resultIfZero := []string{}
  resultIfOne := []string{}
  for _, val := range input {
    if strings.Split(val, "")[charIndex] == "0" {
      resultIfZero = append(resultIfZero, val)
    } else {
      resultIfOne = append(resultIfOne, val)
    }
  }

  lenZero := len(resultIfZero)
  lenOne := len(resultIfOne)
  if lenZero == lenOne {
    if tieChar == "0" {
      return resultIfZero
    } else {
      return resultIfOne
    }
  } else if takeMore {
    if lenZero > lenOne {
      return resultIfZero
    } else {
      return resultIfOne
    }
  } else {
    if lenZero > lenOne {
      return resultIfOne
    } else {
      return resultIfZero
    }
  }
}


