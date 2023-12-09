package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ConvertRange struct {
	src int
	dst int
	len int
}

type ConvertMap struct {
	ranges []ConvertRange
}

func NewConvertMap() ConvertMap {
	ranges := make([]ConvertRange, 30)
	return ConvertMap{
		ranges: ranges,
	}
}

func (m *ConvertMap) addRange(s string) {
	numbers := strings.Fields(s)
	var r [3]int
	for i, number := range numbers {
		n, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		r[i] = n
	}
	m.ranges = append(m.ranges, ConvertRange{
		src: r[1],
		dst: r[0],
		len: r[2],
	})
}

func (m *ConvertMap) convert(n int) int {
	for _, r := range m.ranges {
		if n >= r.src && n < (r.src+r.len) {
			return r.dst + (n - r.src)
		}
	}
	return n
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

	seeds := make([]int, 0, 30)
	maps := make([]ConvertMap, 0, 10)
	mapId := -1
	// parse file
	for fs.Scan() {
		line := fs.Text()
		if len(line) == 0 {
			continue // skip empty lines
		}

		if strings.Contains(line, "seeds") {
			seedFields := strings.Fields(line[strings.IndexByte(line, ':')+1:])
			for _, seed := range seedFields {
				n, err := strconv.Atoi(seed)
				if err != nil {
					panic(err)
				}
				seeds = append(seeds, n)
			}
		} else if strings.Contains(line, "map") {
			mapId++
			maps = append(maps, NewConvertMap())
		} else {
			maps[mapId].addRange(line)
		}
	}
	lowestLocation := 3999999999
	fmt.Println("Computing lowest location...")
	for j := 0; j < len(seeds); j += 2 {
		for seed := seeds[j]; seed < seeds[j]+seeds[j+1]; seed++ {
			x := seed
			for i := 0; i < mapId+1; i++ {
				x = maps[i].convert(x)
			}
			if x < lowestLocation {
				lowestLocation = x
			}
		}
	}
	fmt.Printf("Result: %d\n", lowestLocation)
}
