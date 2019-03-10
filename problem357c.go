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
	"os"
)

func main() {
	stringrep := []rune(os.Args[1])

	count := 0
	maxcount := 0

	for _, r := range stringrep {
		switch r {
		case '(':
			count++
			if count > maxcount {
				maxcount = count
			}
		case ')':
			count--
		case '0':
		default:
		}
	}

	fmt.Printf("Depth %d\n", maxcount)
}
