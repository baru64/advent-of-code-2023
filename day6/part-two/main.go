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

	result := 0
	var timeBuilder strings.Builder
	var distanceBuilder strings.Builder
	// parse 
	fs.Scan()
	line := fs.Text()
	timesFields := strings.Fields(line[strings.IndexByte(line, ':')+1:])
	for _, field := range timesFields {
		timeBuilder.Write([]byte(field))
	}
	time, err := strconv.Atoi(timeBuilder.String())
	if err != nil {
		panic(err)
	}
	fmt.Println(time)
	fs.Scan()
	line = fs.Text()
	distanceFields := strings.Fields(line[strings.IndexByte(line, ':')+1:])
	for _, field := range distanceFields {
		distanceBuilder.Write([]byte(field))
	}
	distance, err := strconv.Atoi(distanceBuilder.String())
	if err != nil {
		panic(err)
	}
	fmt.Println(distance)

	for j := 1; j < time; j += 1 {
		if j*(time-j) > distance {
			result++
		}
	}
	fmt.Printf("Result: %d\n", result)
}
