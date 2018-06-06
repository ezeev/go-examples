package bst_test

import (
	"testing"

	"github.com/ezeev/go-examples/bst"
)

func TestBSTNode(t *testing.T) {

	bt := bst.NewNode([]byte("10"), []byte(""))
	t.Log(bt)

	bt.Insert([]byte("0"), []byte("my key is 0"))
	bt.Insert([]byte("-1"), []byte("lala -1"))
	bt.Insert([]byte("4"), []byte("whoo hoo 4"))
	bt.Insert([]byte("12"), []byte("this is some stuff 12"))
	bt.Insert([]byte("11"), []byte("hello 11"))
	bt.Insert([]byte("20"), []byte("make me soup 20"))
	bt.Insert([]byte("17"), []byte("look at me 17"))
	bt.Insert([]byte("99"), []byte("this is a value"))

	t.Log(bt.Key)
	v := bt.Get([]byte("99"))
	t.Log(string(*v))

}

func TestBSTree(t *testing.T) {

	bt := bst.NewTree()
	bt.Insert([]byte("0"), []byte("my key is 0"))
	bt.Insert([]byte("-1"), []byte("lala -1"))
	bt.Insert([]byte("4"), []byte("whoo hoo 4"))
	bt.Insert([]byte("12"), []byte("this is some stuff 12"))
	bt.Insert([]byte("11"), []byte("hello 11"))
	bt.Insert([]byte("20"), []byte("make me soup 20"))
	bt.Insert([]byte("17"), []byte("look at me 17"))
	bt.Insert([]byte("99"), []byte("this is a value"))

	bt.String()

	// now remove item

	bt.Delete([]byte("4"))

	bt.String()

	// test overwriting a value
	t.Logf("the value for key 20 is \"%s\"", bt.Get([]byte("20")))
	bt.Insert([]byte("20"), []byte("I overwrote the value"))
	t.Logf("the value for key 20 is \"%s\"", bt.Get([]byte("20")))

}
