package dayone

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ptParseInput(data []byte) ([]int, []int) {
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

func ptCalculateOccurence(inp []int) map[int]int {
  right := make(map[int]int);
  for index := 0; index < len(inp)-1; index++ {
    if right[inp[index]] == 0 {
      right[inp[index]] = 1;
      continue;
    }
    right[inp[index]] += 1;
  }

  return right;

}

func ptGetResult(left, right []int) int {
  result := 0;

  rightMap := ptCalculateOccurence(right);

  for index := 0; index < len(left)-1; index++ {
    result += left[index] * rightMap[left[index]];
  }
  
  return result;
}

func ptDayOne() {
  data, err := os.ReadFile("./DayOne/input.txt");
  if err != nil {
    fmt.Printf("ERROR WHILE READING FILE: \n%s", err);
  }

  left, right := ptParseInput(data);

  result := ptGetResult(left, right);

  fmt.Printf("DAY ONE, PART TWO RESULT: %d\n", result);
}
