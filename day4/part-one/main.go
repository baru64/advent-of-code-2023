package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"slices"
)

func parseNumbers(s string) []int {
		numbers := make([]int, 0, 30)
		scanner := bufio.NewScanner(strings.NewReader(s))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			number, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, number)
		}
		return numbers
}

func main() {
	var filename = "testinput"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	fmt.Printf("Processing file %s ...\n", filename)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)

	sum := 0
	// load whole file
	for fs.Scan() {
		line := fs.Text()
		line = line[strings.IndexByte(line, ':')+1:]
		numberStrings := strings.Split(line, " | ")
		winingNumbers := parseNumbers(numberStrings[0])
		yourNumbers := parseNumbers(numberStrings[1])
		won := 0
		for _, number := range yourNumbers {
			if slices.Contains[[]int, int](winingNumbers, number) {
				if won == 0 {
					won = 1
				} else {
					won *= 2
				}
			}
		}
		fmt.Printf("won %d\n", won)
		sum += won
	}
	fmt.Printf("Result: %d\n", sum)
}
