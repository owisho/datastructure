package tree

import "fmt"

type RedBlackTree struct {
	root *RedBlackTreeNode
}

func (self *RedBlackTree) ToString() {
	toString(self.root)
}

func toString(node *RedBlackTreeNode) {
	if node.leftChild != nil {
		toString(node.leftChild)
	}
	isRoot := node.parent == nil
	color := "black"
	if node.isRed {
		color = "red"
	}
	if isRoot {
		fmt.Print(node.data, "+", color, "+root", ",")
	} else {
		fmt.Print(node.data, "+", color, ",")
	}

	if node.rightChild != nil {
		toString(node.rightChild)
	}
}

func (self *RedBlackTree) rotateLeft() {

}

func (self *RedBlackTree) rotateRight() {

}

func (self *RedBlackTree) Insert(data int) (bool, *RedBlackTreeNode) {
	if self.root == nil {
		self.root = &RedBlackTreeNode{
			data:       data,
			isRed:      false,
			parent:     nil,
			leftChild:  nil,
			rightChild: nil,
		}
		return true, self.root
	} else {
		b, node := self.root.Insert(data)
		self.root = searchRoot(self.root)
		return b, node
	}
}

func searchRoot(node *RedBlackTreeNode) *RedBlackTreeNode {
	if node.parent == nil {
		return node
	}
	return searchRoot(node.parent)
}

//搜索数据
func (self *RedBlackTree) search(data int) *RedBlackTreeNode {
	return search(self.root, data)
}

//在树中递归搜索数据，如果搜索到返回节点信息，搜索不到返回nil
func search(node *RedBlackTreeNode, data int) *RedBlackTreeNode {
	if node.data == data {
		return node
	} else if node.data > data {
		if node.leftChild == nil {
			fmt.Println("data not find in tree")
			return nil
		} else {
			return search(node.leftChild, data)
		}
	} else { // node.data < data
		if node.rightChild == nil {
			fmt.Println("data not find in tree")
			return nil
		} else {
			return search(node.rightChild, data)
		}
	}
}

func (self *RedBlackTree) delete(data int) bool {
	//TODO
	return false
}
