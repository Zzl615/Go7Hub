package basic

type TreeNode struct {
	value int32
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(value int32) *TreeNode {
	return &TreeNode{value: value}
}

func (node *TreeNode) AddLeft(left *TreeNode) {
	node.left = left
}

func (node *TreeNode) AddRight(right *TreeNode) {
	node.right = right
}

func (node *TreeNode) GetValue() int32 {
	return node.value
}

func (node *TreeNode) GetLeft() *TreeNode {
	return node.left
}

func (node *TreeNode) GetRight() *TreeNode {
	return node.right
}

func (node *TreeNode) SetValue(value int32) {
	node.value = value
}

type BinarySearchTree struct {
	root *TreeNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (tree *BinarySearchTree) Insert(value int32) {
	if tree.root == nil {
		tree.root = NewTreeNode(value)
		return
	}
	if value < tree.root.GetValue() {
		tree.root.AddLeft(NewTreeNode(value))
	} else {
		tree.root.AddRight(NewTreeNode(value))
	}
}
