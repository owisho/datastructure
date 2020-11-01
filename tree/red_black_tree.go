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
		b, node := compareAndInsert(self.root, data)
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

func compareAndInsert(node *RedBlackTreeNode, data int) (bool, *RedBlackTreeNode) {
	if data == node.data {
		fmt.Println("data already existed in tree")
		return false, node
	}
	if data < node.data {
		lcnode := node.leftChild
		if lcnode == nil {
			nnode := &RedBlackTreeNode{
				data:       data,
				isRed:      true,
				parent:     node,
				leftChild:  nil,
				rightChild: nil,
			}
			node.leftChild = nnode
			adjustTree(nnode)
			return true, nnode
		} else if data < lcnode.data {
			return compareAndInsert(lcnode, data)
		} else if data == lcnode.data {
			return false, lcnode
		} else {
			nnode := &RedBlackTreeNode{
				data:       data,
				isRed:      true,
				parent:     node,
				leftChild:  node.leftChild,
				rightChild: nil,
			}
			node.leftChild = nnode
			adjustTree(nnode)
			return true, nnode
		}
	}
	if data > node.data {
		rcnode := node.rightChild
		if rcnode == nil {
			nnode := &RedBlackTreeNode{
				data:       data,
				isRed:      true,
				parent:     node,
				leftChild:  nil,
				rightChild: nil,
			}
			node.rightChild = nnode
			adjustTree(nnode)
			return true, nnode
		} else if data > node.rightChild.data {
			return compareAndInsert(node.rightChild, data)
		} else if data == node.rightChild.data {
			return false, node.rightChild
		} else { // node.data < data < node.rightChild.data
			nnode := &RedBlackTreeNode{
				data:       data,
				isRed:      true,
				parent:     node,
				leftChild:  nil,
				rightChild: node.rightChild,
			}
			node.rightChild = nnode
			adjustTree(nnode)
			return true, nnode
		}
	}
	fmt.Println("error status, all status should dealed before")
	panic("error status")
}

//调整数，参数为新增加的红色节点
func adjustTree(node *RedBlackTreeNode) {
	if node.parent == nil {
		node.isRed = false
		return
	}
	if !node.parent.isRed {
		//不需要做任何调整
		return
	}
	// 存在父节点并且父节点的颜色为红色
	parent := node.parent
	uncle := getUncle(parent)
	if uncle != nil && uncle.isRed {
		uncle.isRed = false
		parent.isRed = false
		parent.parent.isRed = true
		adjustTree(parent.parent)
	} else {
		gp := parent.parent
		if (node.data < parent.data && parent.data < gp.data) ||
			(node.data > parent.data && parent.data > gp.data) {

			if node.data < parent.data {
				rightRotate(parent, gp)
			} else {
				leftRotate(parent, gp)
			}
		} else {
			if node.data > parent.data { //因为node.data > parent.data和上面的if条件不成立，所以  parent.data < gp.data
				node.parent = gp
				parent.rightChild = node.leftChild
				node.leftChild = parent
				parent.parent = node
				rightRotate(node, gp)
			} else {
				node.parent = gp
				parent.leftChild = node.rightChild
				node.rightChild = parent
				parent.parent = node
				leftRotate(node, gp)
			}
		}
	}

}

func rightRotate(parent, gp *RedBlackTreeNode) {
	parent.parent = gp.parent
	parent.isRed = false
	gp.isRed = true
	gp.parent = parent
	gp.leftChild = parent.rightChild
	parent.rightChild = gp
}

func leftRotate(parent, gp *RedBlackTreeNode) {
	parent.parent = gp.parent
	parent.isRed = false
	gp.isRed = true
	gp.parent = parent
	gp.rightChild = parent.leftChild
	parent.leftChild = gp
}

func getUncle(parent *RedBlackTreeNode) *RedBlackTreeNode {
	if parent.data < parent.parent.data {
		return parent.parent.rightChild
	} else {
		return parent.parent.leftChild
	}
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
