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

func navigate(ghost string, nodes map[string]Node, instruction string) int {
	steps := 0
	currNode := ghost
	for currNode[2] != 'Z' {
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

	nodes := make(map[string]Node, 800)
	ghosts := make([]string, 0, 6)

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
		if nodeName[2] == 'A' {
			ghosts = append(ghosts, nodeName)
		}
	}
	fmt.Println(instruction)
	fmt.Println(nodes)
	fmt.Println(ghosts)

	for _, g := range ghosts {
		result := navigate(g, nodes, instruction)
		fmt.Println(result)
	}
	// Solution is LCM of all results
}
