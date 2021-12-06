package aoc_utils

import (
  "strconv"
)

func DeleteEmptyFromStringSlice(input []string) []string {
  result := []string{}
  for _, val := range input {
    if val != "" {
      result = append(result, val)
    }
  }
  return result
}

func StringSliceContainsAll(inputStrings []string, valuesToCheck[]string) bool {
  foundVals := 0
    for _, valToCheck := range valuesToCheck {
      for _, inputVal := range inputStrings {
        if inputVal == valToCheck {
          foundVals += 1
      }
    }
  }

  return foundVals == len(valuesToCheck)
}

func StringSliceToIntSlice(input []string) []int {
  result := []int{}
  for _, val := range input {
    newInt, _ := strconv.Atoi(val)
    result = append(result, newInt)
  }
  return result
}
