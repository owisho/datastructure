package tree

type RedBlackTreeNode struct {
	parent     *RedBlackTreeNode
	leftChild  *RedBlackTreeNode
	rightChild *RedBlackTreeNode
	data       int
	isRed      bool
}

func (self *RedBlackTreeNode) Insert(data int) (bool, *RedBlackTreeNode) {
	if data == self.data {
		return false, self
	} else if data < self.data {
		if self.leftChild == nil {
			nnode := &RedBlackTreeNode{
				parent: self,
				data:   data,
				isRed:  true,
			}
			self.leftChild = nnode
			nnode.adjust()
			return true, nnode
		} else {
			return self.leftChild.Insert(data)
		}
	} else { // data > self.data
		if self.rightChild == nil {
			nnode := &RedBlackTreeNode{
				parent: self,
				data:   data,
				isRed:  true,
			}
			self.rightChild = nnode
			nnode.adjust()
			return true, nnode
		} else {
			return self.rightChild.Insert(data)
		}
	}
}

//调整数，参数为新增加的红色节点
func (self *RedBlackTreeNode) adjust() {
	if self.parent == nil {
		self.isRed = false
		return
	}
	if !self.parent.isRed {
		//不需要做任何调整
		return
	}
	// 存在父节点并且父节点的颜色为红色
	parent := self.parent
	uncle := self.getUncle()
	if uncle != nil && uncle.isRed { //叔父节点为红色
		uncle.isRed = false
		parent.isRed = false
		parent.parent.isRed = true
		parent.parent.adjust()
	} else {
		gp := parent.parent
		if (self.data < parent.data && parent.data < gp.data) ||
			(self.data > parent.data && parent.data > gp.data) {

			if self.data < parent.data {
				parent.rightRotate()
			} else {
				parent.leftRotate()
			}
		} else {
			if self.data > parent.data { //因为self.data > parent.data和上面的if条件不成立，所以  parent.data < gp.data
				self.parent = gp
				parent.rightChild = self.leftChild
				self.leftChild = parent
				parent.parent = self
				self.rightRotate()
			} else {
				self.parent = gp
				parent.leftChild = self.rightChild
				self.rightChild = parent
				parent.parent = self
				self.leftRotate()
			}
		}
	}
}

//当前节点是左旋的中间节点
func (self *RedBlackTreeNode) leftRotate() {
	parent := self.parent
	gp := parent.parent

	//修改节点的祖父节点的子节点（将数据挂在祖父节点下）
	if gp != nil {
		if parent == gp.rightChild {
			gp.rightChild = self
		} else {
			gp.leftChild = self
		}
	}

	//
	self.parent = gp
	parent.parent = self
	parent.rightChild = self.leftChild
	self.leftChild = parent
	parent.isRed = true
	self.isRed = false
}

//当前节点是左旋的中间节点
func (self *RedBlackTreeNode) rightRotate() {
	parent := self.parent
	gp := parent.parent

	//修改节点的祖父节点的子节点（将数据挂在祖父节点下）
	if gp != nil {
		if parent == gp.rightChild {
			gp.rightChild = self
		} else {
			gp.leftChild = self
		}
	}

	self.parent = gp
	parent.parent = self
	parent.leftChild = self.rightChild
	self.rightChild = parent
	parent.isRed = true
	self.isRed = false
}

func (self *RedBlackTreeNode) getUncle() *RedBlackTreeNode {
	parent := self.parent
	gp := parent.parent
	if parent == gp.leftChild {
		return gp.rightChild
	} else {
		return gp.leftChild
	}
}

func (self *RedBlackTreeNode) delete() {
	next := self.findNext()
	if next == nil {
		parent := self.parent
		parent.rightChild = nil
		if parent.isRed {
			parent.isRed = false
			parent.leftChild.isRed = true
		}
	} else {
		self.data = next.data
		next.delete()
	}
}

//查找后继节点（给删除使用)
func (self *RedBlackTreeNode) findNext() *RedBlackTreeNode {
	if self.rightChild != nil {
		return self.rightChild.findMin()
	} else {
		return self.findSmallestParent()
	}
}

func (self *RedBlackTreeNode) findSmallestParent() *RedBlackTreeNode {
	p := self.parent
	for p != nil {
		if p.data > self.data {
			return p
		}
		p = p.parent
	}
	return nil
}

func (self *RedBlackTreeNode) findMin() *RedBlackTreeNode {
	if self.leftChild != nil {
		return self.leftChild.findMin()
	} else {
		return self
	}
}
