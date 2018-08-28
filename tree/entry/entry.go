package main

import (
	"learngo/tree"
	"fmt"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder () {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()

	myNode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value:3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	//nodes := []Node{
	//	{value:3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)

	root.Print()

	root.Right.Left.SetValue(4)
	root.Right.Left.Print()

	//var pRoot *Node
	//pRoot.SetValue(200)
	//pRoot = &root
	//pRoot.SetValue(300)
	//pRoot.Print()

	fmt.Println()
	root.Traverse()
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()

	fmt.Println()
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount ++
	})
	fmt.Println("Node count:", nodeCount)

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node value:", maxNode)

}
