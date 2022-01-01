package tree

import "github.com/pkg/errors"

func sumInts(ns []int) int {
	var sum int
	for _, n := range ns {
		sum += n
	}
	return sum
}

type Node struct {
	Metadata []int
	Childs   []*Node
}

func (n *Node) SumMeta() int {
	sum := sumInts(n.Metadata)
	for _, c := range n.Childs {
		sum += c.SumMeta()
	}
	return sum
}

func (n *Node) Value() int {
	if len(n.Childs) == 0 {
		return sumInts(n.Metadata)
	}
	var vsum int
	for _, ci := range n.Metadata {
		ci -= 1
		if ci >= 0 && ci < len(n.Childs) {
			child := n.Childs[ci]
			vsum += child.Value()
		}
	}
	return vsum
}

type Tree struct {
	Root *Node
}

func Build(ns []int) (*Tree, error) {
	root, _, err := readNode(ns)
	if err != nil {
		return nil, errors.Wrap(err, "read-node")
	}
	return &Tree{
		Root: root,
	}, nil
}

func (t *Tree) SumMeta() int {
	return t.Root.SumMeta()
}

func (t *Tree) Value() int {
	return t.Root.Value()
}

func readNode(ns []int) (*Node, int, error) {
	if len(ns) < 2 {
		return nil, -1, errors.Errorf("too less data to read node")
	}
	node := &Node{}
	numChilds := ns[0]
	numMeta := ns[1]
	off := 2
	for i := 0; i < numChilds; i++ {
		child, cnt, err := readNode(ns[off:])
		if err != nil {
			return nil, -1, err
		}
		node.Childs = append(node.Childs, child)
		off += cnt
	}
	for mi := 0; mi < numMeta; mi++ {
		node.Metadata = append(node.Metadata, ns[off])
		off++
	}
	return node, off, nil
}
