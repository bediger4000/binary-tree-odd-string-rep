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
I wrote an interface named `TreeNode`:

    type TreeNode interface {
        Left() TreeNode
        Right() TreeNode
        Depth() int
        Print(io.Writer)
        Graph(rune, io.Writer)
        Name(rune) string
    }

I wrote two structs, `InteriorNode` and `LeafNode`,
pointers to which fit that interface.
`InteriorNode` instances have `left` and `right` child elements of type `TreeNode`,
`LeafNode` instances do not have child nodes.
Instances of `*LeafNode` and `*InteriorNode` fit the `TreeNode` interface.
Since only `InteriorNode` has child elements, I had to use functions like `Left() TreeNode`
in the interface, and write them for both pointer types,
even though only `*InteriorNode` really uses them.

The only tough part was thinking of using `func Name(rune) string`.
I wanted to distinguish a node's children, left and right,
by name in the graphviz output, so I had to pass an 'L' or an 'R' to `func Name(rune) string`
for each struct fitting the TreeNode interface.

### Building and running

    $ go build problem357a.go 
    $ ./problem357a '((((00)0)0)0)' unbalanced.dot
    ((((00)0)0)0)
    Depth 4
    $ dot -Tpng -o unbalanced.png unbalanced.dot

This code parses the string on the command line (`((((00)0)0)0)`),
traverses it to print out a string representation for verification,
calculates the maximum depth of the tree, and prints that out.

If a second command line argument exists,
the code treats that like a file name.
It opens the file so named for writing,
and traverses the parsed tree,
this time to create a [DOT language](https://graphviz.gitlab.io/_pages/doc/info/lang.html)
representation for the `dot` program from [graphviz](https://www.graphviz.org) 


The example peculiar string representation of "((((00)0)0)0)" displays like this:

![unbalanced tree image](unbalanceda.png?mode=raw)

### Odd thing I learned

You have to have a struct element to get the Go compiler (go version go1.12 linux/amd64)
to create *different* `LeafNode` structs each time the code allocates one,
which the [graphviz](http://www.graphviz.org) output depends on to make visual sense.
The compiler is smart enough to return the same pointer
if you use `new(LeafNode)` or `&LeafNode{}` to do the allocation.

If you do this:

    type LeafNode struct {
        zork int
    }

the compiler will always put in code to allocate a new `LeafNode` instance.

## Iteration 2

I felt that the Iteration 1 code was a little tricky to get the
depth calculation correct.
I re-wrote the program to use a more C-style single struct for
the entire tree:

    type TreeNode struct {
        left  *TreeNode
        right *TreeNode
    }

`nil` values for `left` and `right` struct elements mean "leaf nodes".

The program has about 20 lines less than the more traditionally "object oriented" program.

### Building and running

    $ go build problem357b.go 
    $ ./problem357b '((((00)0)0)0)' unbalanced.dot
    ((((00)0)0)0)
    Depth 4
    $ dot -Tpng -o unbalanced.png unbalanced.dot

This program has the same story about a 2nd, filename argument on command line.
It will put `graphviz` DOT language output in that file.
I will note that between Go's `os.Create()` and `os.Open()` functions,
about 90% of what the programmer wants to do happens without fuss.

Only one `Depth()`, `printtree()` and `graphtree()` exist in this code.
With only a single type comprising a tree, I had to write less code.
Each function does have to account for `nil` values of `left` and `right` elements.

## Iteration 3

You can count '(' and ')', keeping track of the largest magnitude of the count.
That's the depth of the tree, as if you parsed it and traversed it.
This is linear in the number of symbols making up the peculiar tree representation,
and far less error-prone. It depends on the input strictly following the rules.

I wonder if you get extra points in real interviews for flashes of insight?
They don't come easy, and are improbable in stressful situtations (like interviews).

## Lines of Code

|File | line count |
|----------|----------:|
|problem357a.go|150|
|problem357b.go|122|
|problem357c.go|46|

In this case, I think the line count *underestimates* the cognitive complexity
of the object oriented version,
and *overestimates* the cognitive complexity of counting matching pairs of parens.
The insight needed to realize that you can just count parens makes it far harder overall.
