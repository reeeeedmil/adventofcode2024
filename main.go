package main

import (
	"adventofcode2024/DayOne"
	"adventofcode2024/DayTwo"
	"fmt"
	"os"
	"strconv"
)

func main() {
  fmt.Println("AOC 2024 entrypoint");
  args := os.Args;

  if args[1] == "--day" && args[2] != "" {
    selectedData, err := strconv.ParseInt(args[2], 10, 32);
    if err != nil {
      fmt.Printf("ERROR WHILE SELECTING DAY: \n%s\n", err);
      os.Exit(1);
    }
    switch selectedData {
      case 1: 
        dayone.DayOne();
      case 2:
        daytwo.DayTwo();
    }
  }
}
