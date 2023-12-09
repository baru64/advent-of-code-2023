package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func allZeros(v []int64) bool {
	for i := range v {
		if v[i] != 0 {
			return false
		}
	}
	return true
}

func makeDiffs(hist []int64) [][]int64 {
	diffs := make([][]int64, 0, len(hist)-1)
	for i := 0; i < len(hist)-1; i++ {
		diffs = append(diffs, make([]int64, 0, len(hist)-i))
		for j := 0; j < len(hist)-1-i; j++ {
			//fmt.Printf("%d %d\n", i, j)
			if i == 0 {
				diffs[i] = append(diffs[i], hist[j+1]-hist[j])
			} else {
				diffs[i] = append(diffs[i], diffs[i-1][j+1]-diffs[i-1][j])
			}
		}
		if allZeros(diffs[i]) {
			return diffs
		}
	}
	return diffs
}

func last(v []int64) int64 {
	return v[len(v)-1]
}

func extrapolate(hist []int64) int64 {
	diffs := makeDiffs(hist)
	for i := len(diffs) - 1; i >= 0; i-- {
		if i == len(diffs)-1 {
			diffs[i] = append(diffs[i], 0)
		} else {
			diffs[i] = append(diffs[i], last(diffs[i])+last(diffs[i+1]))
		}
	}
	fmt.Println(hist)
	for _, diff := range diffs {
		fmt.Println(diff)
	}
	fmt.Printf("returning %d+%d=%d\n", last(hist), last(diffs[0]), last(hist)+last(diffs[0]))
	return last(hist) + last(diffs[0])
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

	result := int64(0)

	// parse
	for fs.Scan() {
		line := fs.Text()
		history := make([]int64, 0, 30)
		for _, field := range strings.Fields(line) {
			n, err := strconv.Atoi(field)
			if err != nil {
				panic(err)
			}
			history = append(history, int64(n))
		}
		result = result + extrapolate(history)
	}

	fmt.Printf("Result: %d\n", result)
}
