package linkedlist_test

import (
	"fmt"
	"testing"

	"github.com/ezeev/go-examples/linkedlist"
)

func TestLinkedList(t *testing.T) {

	n1 := linkedlist.NewNode("a")
	n2 := linkedlist.NewNode("b")
	n3 := linkedlist.NewNode("c")
	n4 := linkedlist.NewNode("d")

	l := linkedlist.LinkedList{}
	l.InsertFirst(n1)
	l.InsertFirst(n2)
	l.InsertFirst(n3)
	l.InsertFirst(n4)

	l.Print()

	// last node
	t.Logf("last node: %s", l.Last().Data)

	//first node
	t.Logf("first node: %s", l.First().Data)

	//size
	t.Logf("size: %d", l.Size())

	//at
	t.Logf("element pos 0: %s", l.At(0).Data)
	t.Logf("element pos 2: %s", l.At(2).Data)
	t.Logf("element pos 3: %s", l.At(3).Data)

	//remove first
	l.RemoveFirst()
	t.Log("removed first.")
	l.Print()

	//remove last
	l.RemoveLast()
	t.Log("removed last.")
	l.Print()

	//clear it
	l.Clear()
	t.Log("cleared list.")
	l.Print()

	//insert last
	l.InsertLast(linkedlist.NewNode("a"))
	l.InsertLast(linkedlist.NewNode("b"))
	l.InsertLast(linkedlist.NewNode("c"))
	l.InsertLast(linkedlist.NewNode("d"))
	l.InsertLast(linkedlist.NewNode("e"))
	l.Print()
	t.Log("Added a,b,c,d using InsertLast()")
	t.Logf("element pos 3: %s", l.At(3).Data)

	//loop example
	n := l.First()
	out := "iteration example:"
	for n != nil {
		out = out + fmt.Sprintf(" %s", n.Data)
		n = n.Next
	}
	t.Log(out)

	t.Logf("mid node: %s", l.Mid().Data)

}
