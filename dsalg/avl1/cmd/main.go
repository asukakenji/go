package main

type NodeInterface interface {
	NodeInterface()
}

type Node struct {
}

func (n *Node) NodeInterface() {
}

func f1(ptrN **Node) {
}

func f2(ptrN *NodeInterface) {
}

func main() {
	n := &Node{}
	f1(&n)
	// f2(&n) // Does not work!
}
