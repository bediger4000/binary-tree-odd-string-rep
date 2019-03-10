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

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
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

func (node *TreeNode) Depth() int {
	leftDepth := 0
	if node.left != nil {
		leftDepth = node.left.Depth()
	}
	rightDepth := 0
	if node.right != nil {
		rightDepth = node.right.Depth()
	}

	if leftDepth > rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
}

func constructSubtree(subtree []rune) (*TreeNode, []rune) {
	switch subtree[0] {
	case '(':
		left, remainder := constructSubtree(subtree[1:])
		right, remainder := constructSubtree(remainder)
		if remainder[0] != ')' {
			fmt.Printf("Subtree %q, Remainder not correct: %q\n", string(subtree), string(remainder))
		}
		return &TreeNode{left: left, right: right}, remainder[1:]
	case '0':
		return nil, subtree[1:]
	default:
		panic(fmt.Sprintf("subtree: %q, subtree[0] = %c\n", string(subtree), subtree[0]))
	}
}

func printtree(node *TreeNode) {
	if node == nil {
		fmt.Print("0")
		return
	}
	fmt.Print("(")
	printtree(node.left)
	printtree(node.right)
	fmt.Print(")")
}

func graphtree(node *TreeNode, filename string) {
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

func realgraph(node *TreeNode, out io.Writer) {
	if node.left != nil {
		fmt.Fprintf(out, "n%p -> n%p;\n", node, node.left)
		realgraph(node.left, out)
	} else {
		fmt.Fprintf(out, "n%pL [shape=point];\n", node)
		fmt.Fprintf(out, "n%p -> n%pL;\n", node, node)
	}
	if node.right != nil {
		fmt.Fprintf(out, "n%p -> n%p;\n", node, node.right)
		realgraph(node.right, out)
	} else {
		fmt.Fprintf(out, "n%pR [shape=point];\n", node)
		fmt.Fprintf(out, "n%p -> n%pR;\n", node, node)
	}
}
