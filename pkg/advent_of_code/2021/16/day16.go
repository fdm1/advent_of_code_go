package y2021d16

import (
  "fmt"
  "strconv"
  "strings"
  "github.com/fdm1/advent_of_code_go/pkg/aoc_utils"
)


func Part1(cache bool) string {
  content := aoc_utils.DownloadInput(2021, 16, cache)
  input := aoc_utils.InputToSlice(content)
  return fmt.Sprintf("%v", CalculateSum(input[0], 0, []int64{}, []int64{}))
}

func Part2(cache bool) string {
  content := aoc_utils.DownloadInput(2021, 16, cache)
  input := aoc_utils.InputToSlice(content)
  return fmt.Sprintf("%v", input)
}

func BinaryToInt(binString string) int64 {
  output, _ := strconv.ParseInt(binString, 2, 64)
  return output
}

// Probably a "right" way to do this, but I hate this stuff
func ConvertToBinary(hexString string) []string {
  chars := strings.Split(hexString, "")
  result := ""
  for _, char := range chars {
    if char == "0" { result += "0000" }
    if char == "1" { result += "0001" }
    if char == "2" { result += "0010" }
    if char == "3" { result += "0011" }
    if char == "4" { result += "0100" }
    if char == "5" { result += "0101" }
    if char == "6" { result += "0110" }
    if char == "7" { result += "0111" }
    if char == "8" { result += "1000" }
    if char == "A" { result += "1001" }
    if char == "B" { result += "1010" }
    if char == "C" { result += "1011" }
    if char == "D" { result += "1101" }
    if char == "E" { result += "1110" }
    if char == "F" { result += "1111" }
    fmt.Println(result)
  }
  return strings.Split(result, "")
}

// return versions and values
func CalculateSum(hexString string, pointer int, versions []int64, values []int64) int64 {
  bits := ConvertToBinary(hexString)
  fmt.Printf("starting over @ pointer = %v; len(bits)=%v\n", pointer, len(bits))
  if pointer < len(bits) - 1 {
    fmt.Println(strings.Join(bits, ""))
    newVersions := versions
    newValues := values

    versionString := strings.Join(bits[pointer:pointer+3], "")
    currentVersion := BinaryToInt(versionString)
    newVersions = append(newVersions, currentVersion)
    pointer += 3
    fmt.Printf("pointer = %v; version = %v (%v)\n", pointer, currentVersion, versionString)

    typeString := strings.Join(bits[pointer:pointer+3], "")
    typeId := BinaryToInt(typeString)
    pointer += 3
    fmt.Printf("pointer = %v; typeId = %v (%v)\n", pointer, typeId, typeString)
    if typeId == 4 {
      for i := 0; i < 3; i++ {
        pointer += 1
        newValues = append(newValues, BinaryToInt(strings.Join(bits[pointer:pointer+4], "")))
        pointer += 4
        fmt.Printf("pointer = %v\n", pointer)
      }
      for pointer % 4 != 0 {
        pointer += 1
      }
      return CalculateSum(hexString, pointer, newVersions, newValues)
    }
  } else {
    fmt.Println(versions)
    fmt.Println(values)
    var result int64
    for _, version := range versions {
      result += version
    }
    return result
  }
  return 0
}
