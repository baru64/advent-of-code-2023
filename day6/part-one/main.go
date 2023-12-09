package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	result := 1
	times := make([]int, 0, 5)
	distances := make([]int, 0, 5)
	// parse
	fs.Scan()
	line := fs.Text()
	timesFields := strings.Fields(line[strings.IndexByte(line, ':')+1:])
	for _, field := range timesFields {
		n, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		times = append(times, n)
	}
	fs.Scan()
	line = fs.Text()
	distFields := strings.Fields(line[strings.IndexByte(line, ':')+1:])
	for _, field := range distFields {
		n, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		distances = append(distances, n)
	}
	for i := range times {
		ways := 0
		for j := 1; j < times[i]; j += 1 {
			if j*(times[i]-j) > distances[i] {
				ways++
			}
		}
		fmt.Println(ways)
		result *= ways
	}
	fmt.Printf("Result: %d\n", result)
}
