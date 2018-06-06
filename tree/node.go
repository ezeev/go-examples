package tree

type Node struct {
	Data     interface{}
	Children []*Node
}

func NewNode(data interface{}) *Node {

	ch := make([]*Node, 0)

	n := &Node{Data: data, Children: ch}
	return n
}

func (n *Node) Add(node *Node) {
	n.Children = append(n.Children, node)
}

func (n *Node) Remove(data interface{}) {
	for i, v := range n.Children {
		if v.Data == data {
			n.Children = append(n.Children[:i], n.Children[i+1:]...)
		}
	}
}
