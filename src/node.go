package src

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("{Value:%d Next:%+v}", n.Value, *n.Next)
}
