package bst

import (
	"fmt"
	"sync"
)

type Tree struct {
	root *Node
	lock sync.RWMutex
}

func (t *Tree) Insert(key []byte, value []byte) {
	t.lock.Lock()
	defer t.lock.Unlock()
	if t.root == nil {
		t.root = NewNode(key, value)
	} else {
		t.root.Insert(key, value)
	}
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Get(key []byte) []byte {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return *t.root.Get(key)
}

func (t *Tree) Delete(key []byte) []byte {
	t.lock.Lock()
	defer t.lock.Unlock()
	n := deleteNode(t.root, key)
	if n != nil {
		return n.Value
	} else {
		return nil
	}
}

func (t *Tree) String() {
	t.lock.RLock()
	defer t.lock.RUnlock()
	fmt.Println("------------------------------------------------")
	stringify(t.root, 0)
	fmt.Println("------------------------------------------------")

}
