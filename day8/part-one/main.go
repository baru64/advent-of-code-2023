package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	l string
	r string
}

func navigate(nodes map[string]Node, instruction string) int {
	steps := 0
	currNode := "AAA"
	for currNode != "ZZZ" {
		fmt.Println(currNode)
		fmt.Printf("%c\n", instruction[steps%(len(instruction))])
		switch instruction[steps%(len(instruction))] {
		case 'R':
			currNode = nodes[currNode].r
		case 'L':
			currNode = nodes[currNode].l
		}
		steps++
	}
	return steps
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

	result := 0
	nodes := make(map[string]Node, 800)

	fs.Scan()
	instruction := fs.Text()

	fs.Scan() // skip empty

	// load nodes
	for fs.Scan() {
		line := fs.Text()
		node := Node{}
		var nodeName string
		_, err := fmt.Sscanf(line, "%3s = (%3s, %3s)", &nodeName, &node.l, &node.r)
		if err != nil {
			panic(err)
		}
		nodes[nodeName] = node
	}
	fmt.Println(instruction)
	fmt.Println(nodes)

	result = navigate(nodes, instruction)

	fmt.Printf("Result: %d\n", result)
}
