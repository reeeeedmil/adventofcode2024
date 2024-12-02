package dayone

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func poParseInput(data []byte) ([]int, []int) {
  compiledRegex := regexp.MustCompile(`[\s\p{Zs}]{2,}`);

  deSpaced := compiledRegex.ReplaceAllString(string(data), " ")
  lineData := strings.Split(deSpaced, "\n");

  var left, right []int = make([]int, len(lineData)), make([]int, len(lineData));
  

  for index := 0; index < len(lineData)-1; index++ {
    splitLine := strings.Split(lineData[index], " ")

    leftNumber, err := strconv.ParseInt(splitLine[0], 10, 32);
    if err != nil {
      fmt.Printf("ERROR WHILE CONVERTING LEFT NUMBER DURING INDEX %d: \n%s\n", index, err);
    }

    rightNumber, err := strconv.ParseInt(splitLine[1], 10, 32);
    if err != nil {
      fmt.Printf("ERROR WHILE CONVERTING RIGHT NUMBER DURING INDEX %d: \n%s\n", index, err);
    }

    left[index] = int(leftNumber);
    right[index] = int(rightNumber);
  }

  return left, right;
}

func poSortArray(arr []int) []int {
  slices.Sort(arr);
  slices.Reverse(arr);
  return arr;
}

func poGetDistance(leftNumber, rightNumber int) int {
  if leftNumber > rightNumber {
    return leftNumber - rightNumber;
  }
  return rightNumber - leftNumber;
}

func poGetResult(left, right []int) int {
  result := 0;

  for index := 0; index < len(left)-1; index++ {
    result += poGetDistance(left[index], right[index]);
  }
  
  return result;
}

func poDayOne() {
  data, err := os.ReadFile("./DayOne/input.txt");
  if err != nil {
    fmt.Printf("ERROR WHILE READING FILE: \n%s", err);
  }

  left, right := poParseInput(data);
  left = poSortArray(left);
  right = poSortArray(right);

  result := poGetResult(left, right);

  fmt.Printf("DAY ONE, PART ONE RESULT: %d\n", result);
}
