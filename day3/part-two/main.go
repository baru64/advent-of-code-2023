package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	lid   int
	rid   int
	value int
}

type Gear struct {
	id      int
	numbers []Number
	ratio   int
}

func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func parseNumber(s string, i int) Number {
	number := Number{}
	min := 0
	max := len(s) - 1
	// find left
	number.lid = i
	for (number.lid-1 >= min) && isDigit(s[number.lid-1]) {
		number.lid--
	}
	number.rid = number.lid
	// find right
	for (number.rid+1 <= max) && isDigit(s[number.rid+1]) {
		number.rid++
	}
	v, err := strconv.Atoi(s[number.lid : number.rid+1])
	if err != nil {
		panic(err)
	}
	number.value = v
	return number
}

func (n Number) isOverlapping(o Number) bool {
	if (n.lid >= o.lid && n.lid <= o.rid) || (n.rid >= o.lid && n.rid <= o.rid) {
		return true
	}
	return false
}

func (n Number) isOverlappingList(list []Number) bool {
	for _, o := range list {
		if n.isOverlapping(o) {
			return true
		}
	}
	return false
}

func NewGear(s string, i int, lineLen int) Gear {
	gear := Gear{
		id:      i,
		ratio:   0,
		numbers: make([]Number, 0, 6),
	}
	// find numbers
	indices := [8]int{
		i - 1 - lineLen, i - lineLen, i + 1 - lineLen,
		i - 1, i + 1,
		i - 1 + lineLen, i + lineLen, i + 1 + lineLen,
	}
	for _, j := range indices {
		if j > 0 && j < len(s) && isDigit(s[j]) {
			n := parseNumber(s, j)
			if !n.isOverlappingList(gear.numbers) {
				fmt.Println(n)
				fmt.Println("appending to")
				fmt.Println(gear.numbers)
				gear.numbers = append(gear.numbers, n)
			}
		}
	}
	if len(gear.numbers) == 2 {
		gear.ratio = gear.numbers[0].value * gear.numbers[1].value
	}
	return gear
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
		if c == '*' {
			g := NewGear(schematic, i, lineLen)
			sum += g.ratio
		}
	}
	fmt.Printf("Result: %d\n", sum)
}
