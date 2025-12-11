package day11

import (
	"maps"
	"strings"
	"unicode"
)

type Node struct {
	label string
	conns []*Node
}

func SolvePart1(input <-chan string) int {
	start, end := parse(input)
	return solve(start, end, make(map[*Node]struct{}))
}

func solve(start, end *Node, visited map[*Node]struct{}) int {
	if start == end {
		return 1
	}

	visited[start] = struct{}{}

	paths := 0
	for _, conn := range start.conns {
		if _, ok := visited[conn]; !ok {
			paths += solve(conn, end, maps.Clone(visited))
		}
	}

	return paths
}

func SolvePart2(input <-chan string) int {
	parse(input)
	return 0
}

func parse(input <-chan string) (*Node, *Node) {
	nodeByLabel := make(map[string]*Node)

	for line := range input {
		if line == "" {
			continue
		}

		label, conns, ok := strings.Cut(line, ": ")
		if !ok {
			panic("invalid input")
		}

		node, ok := nodeByLabel[label]
		if !ok {
			node = &Node{label: label}
			nodeByLabel[label] = node
		}

		for _, connLabel := range strings.FieldsFunc(conns, unicode.IsSpace) {
			connNode, ok := nodeByLabel[connLabel]
			if !ok {
				connNode = &Node{label: connLabel}
				nodeByLabel[connLabel] = connNode
			}

			node.conns = append(node.conns, connNode)
		}
	}

	return nodeByLabel["you"], nodeByLabel["out"]
}
