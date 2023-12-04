package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func findFirstDigit(s string, dict map[string]byte) byte {
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			return s[i]
		}
		for k, v := range dict {
			if len(s) >= i+len(k) && k == s[i:i+len(k)] {
				return v
			}
		}
	}
	panic("digit not found")
}

func findLastDigit(s string, dict map[string]byte) byte {
	for i := len(s)-1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			return s[i]
		}
		for k, v := range dict {
			if len(s) >= i+len(k) && k == s[i:i+len(k)] {
				return v
			}
		}
	}
	panic("digit not found")
}


func main() {
	digitDict := map[string]byte{
		"one": byte('1'),
		"two": byte('2'),
		"three": byte('3'),
		"four": byte('4'),
		"five": byte('5'),
		"six": byte('6'),
		"seven": byte('7'),
		"eight": byte('8'),
		"nine": byte('9'),
	}

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
		number[0] = findFirstDigit(line, digitDict)
		number[1] = findLastDigit(line, digitDict)
		n, err := strconv.Atoi(string(number[:]))
		if err != nil {
			panic(err)
		}
		sum += n
	}
	fmt.Printf("Result: %d\n", sum)
}
