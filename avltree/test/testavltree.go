package test

import (
	"github.com/teamelehyean/data_structures/avltree"
	"fmt"
	"strconv"
)

func Test() {
	tree := avltree.NewTree(func(a, b *avltree.Node) int {
		first := a.Value.(int)
		second := b.Value.(int)

		if first > second {
			return -1
		} else if first == second {
			return 0
		} else {
			return 1
		}
	})

	tree.Insert(18)
	tree.Insert(12)
	tree.Insert(20)
	tree.Insert(15)
	tree.Insert(14)

	traversalFunction := func(a *avltree.Node) {
		var parentValue string
		if a.GetParent() != nil {
			parentValue = strconv.Itoa(a.GetParent().Value.(int))
		} else {
			parentValue = ""
		}
		fmt.Printf("[value: %d, parent: %s, height: %d]\n", a.Value.(int), parentValue, a.GetHeight())
	}
	tree.InOrderTraversal(traversalFunction)

	err, node := tree.Find(18)
	if err != nil {
		fmt.Println(err)
	} else {
		tree.Delete(node)
	}

	fmt.Println()

	tree.InOrderTraversal(traversalFunction)

	fmt.Println()
	tree.PostOrderTraversal(traversalFunction)
}
