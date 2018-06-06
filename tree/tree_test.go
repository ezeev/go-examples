package tree_test

import (
	"testing"

	"github.com/ezeev/go-examples/tree"
)

func TestTree(t *testing.T) {

	n1 := tree.NewNode("n1")
	n2 := tree.NewNode("n2")
	n3 := tree.NewNode("n3")

	n1.Add(n2)
	n1.Add(n3)

	if len(n1.Children) != 2 {
		t.Error("children count != 2")
	}

	n1.Remove("n2")

	if len(n1.Children) != 1 {
		t.Error("children count != 1")
	}

	t.Log(n1)
	n1.Add(n2)
	t.Log(n1)

	//build a tree
	root := tree.NewNode(20)
	root.Add(tree.NewNode(0))
	root.Add(tree.NewNode(40))
	root.Add(tree.NewNode(-15))

	root.Children[0].Add(tree.NewNode(12))
	root.Children[0].Add(tree.NewNode(-2))
	root.Children[0].Add(tree.NewNode(1))

	root.Children[2].Add(tree.NewNode(-2))

	tr := tree.NewTree()
	tr.Root = root

	t.Log("Breadth first Traversal")
	tr.TraverseBF(
		func(n *tree.Node) {
			t.Log(n.Data)
		})

	t.Log("Depth first traversal")
	tr.TraverseDF(
		func(n *tree.Node) {
			t.Log(n.Data)
		})

	t.Log(tr.LevelWidth())

}
