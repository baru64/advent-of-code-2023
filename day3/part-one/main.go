package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func containsSymbol(s string) bool {
	for i := range s {
		if (s[i] > '9' || s[i] < '0') && s[i] != '.' {
			return true
		}
	}
	return false
}

func isPart(schematic string, part []byte, i int, lineLen int) bool {
	// line before
	if i > lineLen+len(part)+1 {
		//fmt.Printf("[%d:%d]\n", i-lineLen-len(part)-1,i-lineLen+1)
		s := schematic[i-lineLen-len(part)-1 : i-lineLen+1]
		if containsSymbol(s) {
			return true
		}
	}
	// next line
	if i+lineLen+1 <= len(schematic) {
		//fmt.Printf("[%d:%d]\n", i+lineLen-len(part)-1, i+lineLen+1)
		s := schematic[i+lineLen-len(part)-1 : i+lineLen+1]
		if containsSymbol(s) {
			return true
		}
	}
	// character before
	if i-len(part) > 0 {
		//fmt.Printf("[%d:%d]\n", i-lineLen-len(part)-1,i-lineLen+1)
		s := schematic[i-len(part)-1 : i-len(part)]
		if containsSymbol(s) {
			return true
		}
	}
	// character after
	if i < len(schematic)-1 {
		//fmt.Printf("[%d:%d]\n", i-lineLen-len(part)-1,i-lineLen+1)
		s := schematic[i : i+1]
		if containsSymbol(s) {
			return true
		}
	}
	return false
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
	lineLen := 0
	var buffer strings.Builder
	// load whole file
	for fs.Scan() {
		line := fs.Text()
		lineLen = len(line)
		buffer.WriteString(line)
	}
	schematic := buffer.String()
	part := make([]byte, 0, 3)
	fmt.Println(part)
	for i := range schematic {
		c := schematic[i]
		if c >= '0' && c <= '9' {
			part = append(part, c)
		} else if len(part) > 0 {
			fmt.Println(part)
			if isPart(schematic, part, i, lineLen) {
				num, err := strconv.Atoi(string(part))
				if err != nil {
					panic(err)
				}
				sum += num
				fmt.Printf("num: %d\n", num)
			}
			part = nil
		}
	}
	fmt.Printf("Result: %d\n", sum)
}
