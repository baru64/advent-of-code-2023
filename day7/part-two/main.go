package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	cards [5]uint8
	bid   uint64
	typ   int
}

func getKeys(m map[uint8]int) []uint8 {
	i := 0
	keys := make([]uint8, len(m))
	for k := range m {
		keys[i] = k
		i++
	}
	slices.Sort(keys)
	return keys
}

func NewHand(s string) Hand {
	cardMap := map[byte]uint8{
		'J': 0,
		'2': 1,
		'3': 2,
		'4': 3,
		'5': 4,
		'6': 5,
		'7': 6,
		'8': 7,
		'9': 8,
		'T': 9,
		'Q': 10,
		'K': 11,
		'A': 12,
	}

	fields := strings.Fields(s)
	bid, _ := strconv.Atoi(fields[1])
	hand := Hand{
		bid: uint64(bid),
	}
	for i, c := range fields[0] {
		hand.cards[i] = cardMap[byte(c)]
	}

	// get card kind amounts
	kinds := make(map[uint8]uint8)
	jokers := uint8(0)
	for _, card := range hand.cards {
		if card == 0 {
			jokers++
		} else {
			kinds[card] = kinds[card] + 1
		}
	}
	amounts := make([]uint8, 0, 5)
	for _, num := range kinds {
		amounts = append(amounts, num)
	}
	slices.Sort(amounts)
	if len(amounts) == 0 {
		amounts = append(amounts, 5)
	} else if jokers > 0 {
		amounts[len(amounts)-1] += jokers
	}

	switch len(amounts) {
	case 1:
		hand.typ = 7
	case 2:
		if amounts[1] == 4 {
			hand.typ = 6
		} else {
			hand.typ = 5
		}
	case 3:
		if amounts[2] == 3 {
			hand.typ = 4
		} else {
			hand.typ = 3
		}
	case 4:
		hand.typ = 2
	case 5:
		hand.typ = 1
	}

	return hand
}

func compareHands(a Hand, b Hand) int {
	c := cmp.Compare(a.typ, b.typ)
	if c != 0 {
		return c
	}
	for i := 0; i < 5; i++ {
		c = cmp.Compare(a.cards[i], b.cards[i])
		if c != 0 {
			return c
		}
	}
	return 0
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

	result := uint64(0)
	hands := make([]Hand, 0, 1000)
	// parse
	for fs.Scan() {
		line := fs.Text()
		hands = append(hands, NewHand(line))
	}
	fmt.Println(hands)

	slices.SortFunc[[]Hand, Hand](hands, compareHands)
	for i := range hands {
		result += hands[i].bid * uint64(i+1)
	}
	fmt.Printf("Result: %d\n", result)
}
