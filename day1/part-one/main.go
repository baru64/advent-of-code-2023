package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func findFirstDigit(s string) byte {
	for k := 0; k < len(s); k++ {
		c := s[k]
		if i := c - '0'; i >= 0 && i <= 9 {
			return c
		}
	}
	panic("digit not found")
}

func findLastDigit(s string) byte {
	for k := len(s) - 1; k >= 0; k-- {
		c := s[k]
		if i := int(c) - '0'; i >= 0 && i <= 9 {
			return c
		}
	}
	panic("digit not found")
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
	for fs.Scan() {
		line := fs.Text()
		var number [2]byte
		number[0] = findFirstDigit(line)
		number[1] = findLastDigit(line)
		n, err := strconv.Atoi(string(number[:]))
		if err != nil {
			panic(err)
		}
		sum += n
	}
	fmt.Printf("Result: %d\n", sum)
}
