package tree

import (
	"fmt"
)

type Tree struct {
	Root *Node
}

func NewTree() *Tree {
	t := &Tree{}
	return t
}

func (t *Tree) TraverseBF(action func(n *Node)) {
	arr := make([]*Node, 0)
	arr = append(arr, t.Root)
	for len(arr) > 0 {
		node := arr[0]
		//remove it
		arr = append(arr[:0], arr[1:]...)
		// add the node's children
		arr = append(arr, node.Children...)
		action(node)
	}

}

func (t *Tree) LevelWidth() []int {
	counters := make([]int, 0)
	counters = append(counters, 0)
	arr := make([]*Node, 0)
	arr = append(arr, t.Root)
	arr = append(arr, NewNode("s")) // s will be used as the level terminator
	for len(arr) > 1 {
		node := arr[0]
		fmt.Println(node.Data)
		arr = append(arr[:0], arr[1:]...)
		if node.Data == "s" {
			counters = append(counters, 0)
			arr = append(arr, node)
		} else {
			arr = append(arr, node.Children...)
			//increment last value in counters
			if len(counters) > 0 {
				counters[len(counters)-1] = counters[len(counters)-1] + 1
				fmt.Printf("level counter %d\n", counters[len(counters)-1])
			}
		}
	}
	return counters
}

func (t *Tree) TraverseDF(action func(n *Node)) {

	arr := make([]*Node, 0)
	arr = append(arr, t.Root)
	for len(arr) > 0 {
		node := arr[0]
		arr = append(arr[:0], arr[1:]...)
		// same as TraverseBF, except add nodes to the beginning instead of the end
		arr = append(node.Children, arr...)
		action(node)
	}

}
