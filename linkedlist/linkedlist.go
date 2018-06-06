package linkedlist

import (
	"fmt"
)

// ordered collection of data
type Node struct {
	Data interface{}
	Next *Node
}

func NewNode(data interface{}) *Node {
	return &Node{Data: data}
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) InsertFirst(n *Node) {
	n.Next = l.Head
	l.Head = n
}

func (l *LinkedList) InsertLast(nn *Node) {
	if l.Head == nil {
		l.Head = nn
		return
	}
	n := l.Head
	for n != nil {
		if n.Next == nil {
			n.Next = nn
			return
		}
		n = n.Next
	}
}

func (l *LinkedList) Print() {
	fmt.Println("LinkedList Print()")
	n := l.Head
	for n != nil {
		fmt.Println(n.Data)
		n = n.Next
	}
}

func (l *LinkedList) Size() int {
	cnt := 0
	n := l.Head
	for n != nil {
		cnt++
		n = n.Next
	}
	return cnt
}

func (l *LinkedList) First() *Node {
	return l.Head
}

func (l *LinkedList) At(pos int) *Node {
	cnt := 0
	n := l.Head
	for n != nil {
		if cnt == pos {
			return n
		}
		n = n.Next
		cnt++
	}
	return nil
}

func (l *LinkedList) Last() *Node {
	if l.Head == nil {
		return nil
	}
	n := l.Head
	for n != nil {
		if n.Next == nil {
			return n
		}
		n = n.Next
	}
	return n
}

func (l *LinkedList) Clear() {
	l.Head = nil
}

func (l *LinkedList) RemoveFirst() {
	l.Head = l.Head.Next
}

// Not possible in Go? (https://stackoverflow.com/questions/42066797/how-to-delete-struct-object-in-go)
func (l *LinkedList) RemoveLast() {
	n := l.Head
	moreNodes := true
	for moreNodes {
		/*if n.Next == nil {
			//this is the last node!
			n = nil
			moreNodes = false
		} else {
			n = n.Next
		}*/

		//alternate solution for Go
		if n.Next == nil {
			moreNodes = false
		} else {
			if n.Next.Next == nil {
				//peek ahead to remove the next reference
				n.Next = nil
				moreNodes = false
			}
			n = n.Next
		}
	}
}

// Mid returns the node at the mid point of the LinkedList
// It would be easy to just use LinkedList.At(Size()/2) to get the midpoint,
// but this would required multiple loops internally. This approach uses a single loop,
// the slow var is incremented by one and fast var is incremented by 2 with each
// iteration. This means that when fast has reached the end, slow is now at the midpoint.
func (l *LinkedList) Mid() *Node {
	slow := l.First()
	fast := l.First()
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// Circular detect if the LinkedList has a circular reference
func (l *LinkedList) Circular() bool {
	slow := l.First()
	fast := l.First()
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
