package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// 工厂方法
func CreateNode(value int) *Node {
	return &Node{Value:value} // go里可以返回局部变量地址给外层用
}

func (node Node) Print()  {
	fmt.Print(" ", node.Value)
}

func (node *Node) SetValue(value int)  {
	if node == nil {
		fmt.Println("node is nil ignored!")
		return
	}
	node.Value = value
}




