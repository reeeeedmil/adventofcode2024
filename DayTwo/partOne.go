package daytwo

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func poParseInput(data []byte) [][]int {
  lines := strings.Split(string(data), "\n");

  parsedLines := make([][]int, len(lines));
  for i := range lines {
    parsedLines[i] = make([]int, 5)
  }

  for index := range lines {
    splitLine := strings.Split(strings.Trim(lines[index], " "), " ")
    fmt.Println(splitLine)

    for i := range 5 {
      parsedInt, err := strconv.ParseInt(splitLine[i], 10, 32);
      if err != nil {
        fmt.Printf("ERROR WHILE PARSING LINE INDEX %d: \n%s\n", index, err);
        os.Exit(1);
      }
      parsedLines[index][i] = int(parsedInt);
    }
  }

  return parsedLines;
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

func poDayTwo() {
  data, err := os.ReadFile("./DayTwo/testInput.txt");
  if err != nil {
    fmt.Printf("ERROR WHILE READING FILE: \n%s", err);
  }

  parsedInput := poParseInput(data);
  fmt.Println(parsedInput);
}
