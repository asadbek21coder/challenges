package veryhard

// Create a BinaryTree struct with methods Insert
// (to add a value to the tree) and Search (to find if a
// value exists in the tree). Implement these methods
// using a binary search tree algorithm.
//

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func (bt *BinaryTree) Insert(value int) {
	if bt.Root == nil {
		bt.Root = &Node{Value: value}
		return
	}

	current := bt.Root
	for {
		if value < current.Value {
			if current.Left == nil {
				current.Left = &Node{Value: value}
				return
			}
			current = current.Left
		} else if value > current.Value {
			if current.Right == nil {
				current.Right = &Node{Value: value}
				return
			}
			current = current.Right
		} else {
			// Value already exists in the tree, no duplicate allowed.
			return
		}
	}
}

func (bt *BinaryTree) Search(value int) bool {
	current := bt.Root
	for current != nil {
		if value < current.Value {
			current = current.Left
		} else if value > current.Value {
			current = current.Right
		} else {
			// Value found in the tree.
			return true
		}
	}
	// Value not found in the tree.
	return false
}
