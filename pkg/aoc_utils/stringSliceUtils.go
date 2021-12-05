package aoc_utils

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
