package tree

type RedBlackTreeNode struct {
	parent     *RedBlackTreeNode
	leftChild  *RedBlackTreeNode
	rightChild *RedBlackTreeNode
	data       int
	isRed      bool
}
