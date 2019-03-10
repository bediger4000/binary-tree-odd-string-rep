# Daily Coding Problem: Problem #357 [Hard]

Problem statement:

> You are given a binary tree in a peculiar string representation. Each node is
> written in the form (lr), where l corresponds to the left child and r
> corresponds to the right child.
> 
> If either l or r is null, it will be represented as a zero. Otherwise, it will
> be represented by a new (lr) pair.
> 
> Here are a few examples:
> 
>     A root node with no children: (00)
>     A root node with two children: ((00)(00))
>     An unbalanced tree with three consecutive left children: ((((00)0)0)0)
> 
> Given this representation, determine the depth of the tree.

The problem statement frames this as a binary tree algorithm.
Thinking from there, parsing the "peculiar string representation" constitutes
the crux of the matter.
The problem uses depth of the tree as a way of determining if you got that parse correct.

## Iteration 1

I wrote a standard, [two-class, object oriented binary tree](problem357a.go):
one class for interior nodes,
one class for exterior nodes.
Because I wrote this in Go, 
I wrote an interface namedd `TreeNode`:

    type TreeNode interface {
        Left() TreeNode
        Right() TreeNode
        Depth() int
    }

I wrote two structs, `InteriorNode` and `LeafNode`,
pointers to which  fit that interface.

### Odd thing I learned

You have to have a struct element to get the Go compiler
to create *different* `LeafNode` structs,
which the [graphviz](http://www.graphviz.org) output depends on to make visual sense.

## Iteration 2

## Iteration 3

You can count '(' and ')', keeping track of the largest magnitude of the count.
That's the depth of the tree, if you parsed it and traversed it.

This depends on the input strictly following the rules.
