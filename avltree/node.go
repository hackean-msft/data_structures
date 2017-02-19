package avltree

import (
	"math"
	"errors"
)

type Node struct {
	parent *Node
	Value  interface{}
	left   *Node
	right  *Node
	height int
}

var comp func(a, b *Node) int

func create(value interface{}) (node *Node) {
	return &Node{Value:value}
}

func getHeight(node *Node) int {
	if node != nil {
		return node.height
	}
	return -1
}

func balancedFactor(node *Node) int {
	left := getHeight(node.left)
	right := getHeight(node.right)
	return left - right
}

func (n *Node) add(child *Node, tree *Tree) {
	greaterFactor := comp(n, child)
	if greaterFactor > 0 {
		if n.right != nil {
			n.right.add(child, tree)
		} else {
			n.setRight(child)
		}
	} else {
		if n.left != nil {
			n.left.add(child, tree)
		} else {
			n.setLeft(child)
		}

	}

	n.updateHeight()
	if !n.IsBalanced() {
		rotate(n, tree)
	}

}

func (n *Node) updateHeight() {
	hleft := getHeight(n.left)
	hright := getHeight(n.right)
	n.height = int(math.Max(float64(hleft), float64(hright))) + 1
}

func (n *Node) Left() (*Node) {
	return n.left
}

func (n *Node) Right() (*Node) {
	return n.right
}

func (n *Node) IsLeave() bool {
	return n.right == nil && n.left == nil
}

func (n *Node) IsBalanced() bool {
	return math.Abs(float64(balancedFactor(n))) < 2
}

func (n *Node) GetHeight() int {
	return n.height
}

func (n *Node) GetParent() (*Node) {
	return n.parent
}

func (n *Node) setLeft(node *Node) {
	n.left = node
	if node != nil {
		node.parent = n
	}
}

func (n *Node) setRight(node *Node) {
	n.right = node
	if node != nil {
		node.parent = n
	}
}

func (n *Node) detachChild(node *Node) error {
	if n.left == node {
		n.left = nil
		node.parent = nil
		return nil
	}

	if n.right == node {
		n.right = nil
		node.parent = nil
		return nil
	}

	return errors.New("Node wasn't found")
}

func rotate(n *Node, tree *Tree) {
	if !n.IsBalanced() {
		bf := balancedFactor(n)
		if bf > 0 {
			bfNext := balancedFactor(n.left)
			if bfNext > 0 {
				leftLeftRotation(n, tree)
			} else {
				leftRightRotation(n, tree)
			}
		} else {
			bfNext := balancedFactor(n.right)
			if bfNext > 0 {
				rightLeftRotation(n, tree)
			} else {
				rightRightRotation(n, tree)
			}
		}
	}
}

func leftLeftRotation(pivot *Node, tree *Tree) {
	child := pivot.left

	// House keeping
	fixTree(pivot, child, tree)

	pivot.setLeft(child.right)
	child.setRight(pivot)

	updateChild(child)

}

func leftRightRotation(pivot *Node, tree *Tree) {
	rightRightRotation(pivot.left, tree)
	leftLeftRotation(pivot, tree)
}

func rightRightRotation(pivot *Node, tree *Tree) {
	child := pivot.right

	fixTree(pivot, child, tree)

	pivot.setRight(child.left)
	child.setLeft(pivot)

	updateChild(child)
}

func rightLeftRotation(pivot *Node, tree *Tree) {
	leftLeftRotation(pivot.right, tree)
	rightRightRotation(pivot, tree)
}

func fixTree(pivot, pivotChild *Node, tree *Tree) {
	if tree.root == pivot {
		tree.root = pivotChild
		pivotChild.parent = nil
	} else {
		if pivot.parent.left == pivot {
			pivot.parent.setLeft(pivotChild)
		} else {
			pivot.parent.setRight(pivotChild)
		}
	}
}

func updateChild(child *Node)  {
	if child.left != nil {
		child.left.updateHeight()
	}

	if child.right != nil {
		child.right.updateHeight()
	}
	child.updateHeight()
}
