package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id   int
	sets []Set
}

func parseSet(s string) Set {
	cubes := strings.Split(s, ",")
	set := Set{
		red:   0,
		green: 0,
		blue:  0,
	}
	for _, cube := range cubes {
		splitted := strings.Split(strings.TrimSpace(cube), " ")
		count, err := strconv.Atoi(splitted[0])
		if err != nil {
			panic(err)
		}
		switch splitted[1] {
		case "red":
			set.red = count
		case "blue":
			set.blue = count
		case "green":
			set.green = count
		}
	}
	return set
}

func (s Set) getPower() int {
	return s.red * s.green * s.blue
}

func NewGame(s string) Game {
	game := Game{
		id:   0,
		sets: make([]Set, 0, 5),
	}
	gameSplit := strings.Split(s, ":")
	idSplit := strings.Split(strings.TrimSpace(gameSplit[0]), " ")
	id, err := strconv.Atoi(idSplit[1])
	if err != nil {
		panic(err)
	}
	game.id = id
	setsSplit := strings.Split(gameSplit[1], ";")
	for _, set := range setsSplit {
		game.sets = append(game.sets, parseSet(set))
	}
	return game
}

func (g Game) isPossible(bag Set) bool {
	for _, set := range g.sets {
		if set.red > bag.red || set.green > bag.green || set.blue > bag.blue {
			return false
		}
	}
	return true
}

func (g Game) getMinimalSet() Set {
	set := Set{
		red:   0,
		green: 0,
		blue:  0,
	}
	for _, s := range g.sets {
		if s.red > set.red {
			set.red = s.red
		}
		if s.green > set.green {
			set.green = s.green
		}
		if s.blue > set.blue {
			set.blue = s.blue
		}
	}
	return set
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
		game := NewGame(line)
		sum += game.getMinimalSet().getPower()
	}
	fmt.Printf("Result: %d\n", sum)
}
