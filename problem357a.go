/*
You are given a binary tree in a peculiar string representation. Each node is
written in the form (lr), where l corresponds to the left child and r
corresponds to the right child.

If either l or r is null, it will be represented as a zero. Otherwise, it will
be represented by a new (lr) pair.

Here are a few examples:

    A root node with no children: (00)
    A root node with two children: ((00)(00))
    An unbalanced tree with three consecutive left children: ((((00)0)0)0)

Given this representation, determine the depth of the tree.
*/

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type TreeNode interface {
	Left() TreeNode
	Right() TreeNode
	Depth() int
}

type InteriorNode struct {
	left  TreeNode
	right TreeNode
}

func (node *InteriorNode) Left() TreeNode {
	return node.left
}

func (node *InteriorNode) Right() TreeNode {
	return node.right
}

func (node *InteriorNode) Depth() int {
	leftDepth := node.Left().Depth()
	rightDepth := node.Right().Depth()

	if leftDepth > rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
}

func (node *LeafNode) Depth() int {
	return 0
}

type LeafNode struct {
	zork int // compiler knows about empty structs, always uses the same one
}

func (leaf *LeafNode) Left() TreeNode {
	return nil
}

func (leaf *LeafNode) Right() TreeNode {
	return nil
}

func main() {
	stringrep := []rune(os.Args[1])

	root, remainder := constructSubtree(stringrep)
	// remainder should be zero-length string
	if len(remainder) != 0 {
		fmt.Printf("Remainder not zero length: %q\n", remainder)
	}

	printtree(root)
	fmt.Println()
	fmt.Printf("Depth %d\n", root.Depth())

	if len(os.Args) > 2 {
		filename := os.Args[2]
		graphtree(root, filename)
	}
}

func constructSubtree(subtree []rune) (TreeNode, []rune) {
	if subtree[0] == '(' {
		left, remainder := constructSubtree(subtree[1:])
		right, remainder := constructSubtree(remainder)
		if remainder[0] != ')' {
			fmt.Printf("Subtree %d, Remainder not correct: %q\n", string(subtree), string(remainder))
		}
		return &InteriorNode{left: left, right: right}, remainder[1:]
	}
	if subtree[0] == '0' {
		return new(LeafNode), subtree[1:]
	}
	panic(fmt.Sprintf("subtree: %q, subtree[0] = %c\n", string(subtree), subtree[0]))
}

func printtree(node TreeNode) {
	switch node.(type) {
	case *InteriorNode:
		fmt.Print("(")
		printtree(node.Left())
		printtree(node.Right())
		fmt.Print(")")
	case *LeafNode:
		fmt.Print("0")
	}
}

func graphtree(node TreeNode, filename string) {
	f, e := os.Create(filename)
	if e != nil {
		log.Printf("creating %s: %v\n", filename, e)
		return
	}
	fmt.Fprintf(f, "digraph g {\n")
	realgraph(node, f)
	fmt.Fprintf(f, "}\n")
	f.Close()
}

func realgraph(node TreeNode, out io.Writer) {
	switch node.(type) {
	case *InteriorNode:
		fmt.Fprintf(out, "n%p;\n", node)

		left := node.Left()
		switch left.(type) {
		case *LeafNode:
			fmt.Fprintf(out, "n%pL [shape=point];\n", left)
			fmt.Fprintf(out, "n%p -> n%pL;\n", node, left)
		case *InteriorNode:
			realgraph(node.Left(), out)
			fmt.Fprintf(out, "n%p -> n%p;\n", node, left)
		}

		right := node.Right()
		switch right.(type) {
		case *LeafNode:
			fmt.Fprintf(out, "n%pR [shape=point];\n", right)
			fmt.Fprintf(out, "n%p -> n%pR;\n", node, right)
		case *InteriorNode:
			realgraph(right, out)
			fmt.Fprintf(out, "n%p -> n%p;\n", node, right)
		}
	}
}
