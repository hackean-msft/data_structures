package avltree

import (
	"errors"
)

type Tree struct {
	root *Node
}

func NewTree(comparator func(a, b *Node) int) (*Tree) {
	comp = comparator
	return &Tree{}
}

func (tree *Tree) Insert(value interface{}) {
	n := create(value)
	rootNode := tree.root
	if rootNode == nil {
		n.height = 0
		tree.root = n
	} else {
		rootNode.add(n, tree)
	}
}

func (tree *Tree) InOrderTraversal(do func(a *Node)) {
	inOrderTraversal(tree.root, do)
}

func (tree *Tree) PreOrderTraversal(do func(a *Node))  {
	preOrderTraversal(tree.root, do)
}
func (tree *Tree) PostOrderTraversal(do func(a *Node))  {
	postOrderTraversal(tree.root, do)
}

func (tree *Tree) Delete(node *Node) error {
	if node == nil {
		return errors.New("Couldn't found node with the value specified")
	}

	if node.IsLeave() {
		parent := node.parent
		parent.detachChild(node)
		tree.rebalance(parent)
	}

	if node.left != nil && node.right == nil {
		parent := node.parent
		parent.setLeft(node.left)
		tree.rebalance(parent)
	} else if node.left == nil && node.right != nil {
		parent := node.parent
		parent.setRight(node.right)
		tree.rebalance(parent)
	}

	if node.left != nil && node.right != nil {
		_, successor := tree.Successor(node)
		node.Value = successor.Value
		tree.Delete(successor)
	}
	
	return nil
}

func (tree *Tree) Find(value interface{}) (error, *Node){
	n := create(value)
	foundNode := tree.compare(n, tree.root)
	if foundNode == nil {
		return errors.New("Node was not found"), nil
	}
	return nil, foundNode
}

func (tree *Tree) Successor(node *Node) (error, *Node) {
	if node.right != nil {
		return nil, tree.min(node.right)
	} else {
		return errors.New("Node has no right child"), nil
	}
}

func (tree *Tree) min(node *Node) *Node{
	if node.left == nil {
		return node
	}
	return tree.min(node.left)
}

func (tree *Tree) compare(node, current *Node) (*Node) {
	if current == nil {
		return nil
	}
	comparison := comp(current, node)
	if comparison == 0 {
		return current
	} else if comparison > 0 {
		return tree.compare(node, current.right)
	} else {
		return tree.compare(node, current.left)
	}
}

func (tree *Tree) rebalance(node *Node)  {
	if node != nil {
		node.updateHeight()
		if !node.IsBalanced() {
			bf := balancedFactor(node)
			if bf > 0 {
				bfNext := balancedFactor(node.left)
				if bfNext == 0 || bfNext == 1{
					leftLeftRotation(node, tree)
				} else {
					leftRightRotation(node, tree)
				}
			} else {
				bfNext := balancedFactor(node.right)
				if bfNext == 0 || bfNext == 1{
					rightRightRotation(node, tree)
				} else {
					rightLeftRotation(node, tree)
				}
			}
		} else {
			tree.rebalance(node.parent)
		}
	}
}

func inOrderTraversal(node *Node, do func(a *Node)) {
	if node != nil {
		inOrderTraversal(node.left, do)
		do(node)
		inOrderTraversal(node.right, do)
	}
}

func preOrderTraversal(node *Node, do func(a *Node))  {
	if node != nil {
		do(node)
		preOrderTraversal(node.left, do)
		preOrderTraversal(node.right, do)
	}
}

func postOrderTraversal(node *Node, do func(a *Node))  {
	if node != nil {
		postOrderTraversal(node.left, do)
		postOrderTraversal(node.right, do)
		do(node)
	}
}