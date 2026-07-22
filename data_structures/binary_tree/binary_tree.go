package binarytree

type Node struct {
	value int
	right *Node
	left  *Node
}

func newNode(value int) *Node {
	return &Node{value: value}
}

func (self *Node) Value() int {
	return self.value
}

func (self *Node) Right() *Node {
	return self.right
}

func (self *Node) Left() *Node {
	return self.left
}

type BinaryTree struct {
	root *Node
}

func New() *BinaryTree {
	return &BinaryTree{}
}

func (self *BinaryTree) Root() *Node {
	return self.root
}

func (self *BinaryTree) Insert(value int) {
	self.root = insert(self.root, value)
}

func insert(node *Node, value int) *Node {
	if node == nil {
		return newNode(value)
	}
	if value < node.value {
		node.left = insert(node.left, value)
	} else if value > node.value {
		node.right = insert(node.right, value)
	}
	return node
}

func (self *BinaryTree) InsertIterative(value int) {
	if self.root == nil {
		self.root = newNode(value)
		return
	}
	current := self.root
	for {
		if value < current.value {
			if current.left == nil {
				current.left = newNode(value)
				return
			}
			current = current.left
		} else if value > current.value {
			if current.right == nil {
				current.right = newNode(value)
				return
			}
			current = current.right
		} else {
			return
		}
	}
}

func (self *BinaryTree) Search(value int) bool {
	return self.search(self.root, value)
}

func (self *BinaryTree) search(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if value == node.value {
		return true
	}
	if value < node.value {
		return self.search(node.left, value)
	} else {
		return self.search(node.right, value)
	}
}

func (self *BinaryTree) SearchIterative(value int) bool {
	if self.root == nil {
		return false
	}
	current := self.root
	for current != nil {
		if value == current.value {
			return true
		}
		if value < current.value {
			current = current.left
		} else {
			current = current.right
		}
	}
	return false
}

func (self *BinaryTree) InOrder() []int {
	result := make([]int, 0)
	self.inOrder(self.root, &result)
	return result
}

func (self *BinaryTree) inOrder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	self.inOrder(node.left, result)
	*result = append(*result, node.value)
	self.inOrder(node.right, result)
}
