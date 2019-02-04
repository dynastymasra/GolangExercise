package main

import (
	"fmt"
	"math/rand"
)

var r = rand.New(rand.NewSource(99))

type Node struct {
	Left  *Node
	Value int
	Right *Node
}

func New(d int) *Node {
	var node *Node

	node = insert(node, 0, d)

	return node
}

func (node *Node) Print(d int) {
	if node == nil {
		fmt.Println("Binary Tree is empty")
		return
	}

	for i := 0; i <= d; i++ {
		printNode(node, i, true)
		fmt.Println()
	}
}

func printNode(node *Node, i int, end bool) {
	if node == nil {
		return
	}

	if i == 0 {
		if end {
			fmt.Printf("(%d)", node.Value)
		} else {
			fmt.Printf("(%d),", node.Value)
		}
	}

	if i > 0 {
		printNode(node.Left, i-1, false)
		printNode(node.Right, i-1, end)
	}
}

func insert(node *Node, nd, d int) *Node {
	if nd > d {
		return nil
	}

	if node == nil {
		node = &Node{nil, r.Intn(10), nil}
	}

	node.Left = insert(node.Left, nd+1, d)
	node.Right = insert(node.Right, nd+1, d)

	return node
}

func main() {
	var n int

	fmt.Print("Input depth value : ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}

	node := New(n)
	node.Print(n)
}
