package bst

import (
	"bytes"
	"fmt"
)

type Node struct {
	Key   []byte
	Value []byte
	Left  *Node
	Right *Node
}

func NewNode(key []byte, value []byte) *Node {
	return &Node{Key: key, Value: value, Left: nil, Right: nil}
}

func (n *Node) Insert(key []byte, value []byte) {
	newNode := NewNode(key, value)
	if n == nil {
		n = newNode
	} else {
		insertNode(n, newNode)
	}

}

func insertNode(node, newNode *Node) {
	if bytes.Compare(newNode.Key, node.Key) == -1 {
		if node.Left == nil {
			node.Left = newNode
		} else {
			insertNode(node.Left, newNode)
		}
	} else if bytes.Compare(newNode.Key, node.Key) == 1 {
		if node.Right == nil {
			node.Right = newNode
		} else {
			insertNode(node.Right, newNode)
		}
	} else if bytes.Compare(newNode.Key, node.Key) == 0 {
		//over write value
		node.Value = newNode.Value
	}
}

func (n *Node) Get(key []byte) *[]byte {
	if bytes.Compare(n.Key, key) == 0 {
		return &n.Value
	} else if n.Left != nil && bytes.Compare(key, n.Key) == -1 {
		return n.Left.Get(key)
	} else if n.Right != nil && bytes.Compare(key, n.Key) == 1 {
		return n.Right.Get(key)
	}
	return nil
}

func (n *Node) get(key []byte) *Node {
	if bytes.Compare(n.Key, key) == 0 {
		return n
	} else if n.Left != nil && bytes.Compare(key, n.Key) == -1 {
		return n.Left.get(key)
	} else if n.Right != nil && bytes.Compare(key, n.Key) == 1 {
		return n.Right.get(key)
	}
	return nil
}

func deleteNode(root *Node, key []byte) *Node {
	if root == nil {
		return nil
	}
	if bytes.Compare(key, root.Key) == -1 {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if bytes.Compare(key, root.Key) == 1 {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	if root.Left == nil && root.Right == nil {
		root = nil
		return nil
	}
	if root.Left == nil {
		root = root.Right
		return root
	}
	if root.Right == nil {
		root = root.Left
		return root
	}
	lmrs := root.Right
	for {
		if lmrs != nil && lmrs.Left != nil {
			lmrs = lmrs.Left
		} else {
			break
		}
	}
	root.Key, root.Value = lmrs.Key, lmrs.Value
	root.Right = deleteNode(root.Right, root.Key)
	return root
}

func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += ""
		}
		//format += "---[ "
		level++
		stringify(n.Left, level)
		fmt.Printf(format+"%s\n", string(n.Key))
		stringify(n.Right, level)
	}
}
