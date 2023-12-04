package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
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

func countNumbers(w []int, y []int) int {
	count := 0
	for _, n := range y {
		if slices.Contains[[]int, int](w, n) {
			count++
		}
	}
	return count
}

func count(i int, games []int) int {
	v := games[i]
	c := 0
	for k := 0; k < v; k++ {
		c += count(i+1+k, games)
	}
	return c + 1
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
	games := make([]int, 0, 220)
	for fs.Scan() {
		line := fs.Text()
		line = line[strings.IndexByte(line, ':')+1:]
		numberStrings := strings.Split(line, " | ")
		winingNumbers := parseNumbers(numberStrings[0])
		yourNumbers := parseNumbers(numberStrings[1])
		games = append(games, countNumbers(winingNumbers, yourNumbers))
	}
	for i := range games {
		sum += count(i, games)
	}
	fmt.Printf("Result: %d\n", sum)
}
