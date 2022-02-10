package main

import "fmt"

func main() {
	// REGULAR TREES
	t0 := &Tree{info: "t0"}
	t1 := &Tree{info: "t1"}
	t2 := &Tree{info: "t2"}
	t3 := &Tree{info: "t3"}
	t4 := &Tree{info: "t4"}
	t5 := &Tree{info: "t5"}

	// regular links

	t0.llink = t1
	t0.rlink = t2
	t1.llink = t3
	t1.rlink = t4
	t4.llink = t5

	fmt.Println("pre-order:")
	PreOrder(t0)
	fmt.Print("\n\n")

	fmt.Println("in-order:")
	InOrder(t0)
	fmt.Print("\n\n")

	fmt.Println("post-order:")
	PostOrder(t0)
	fmt.Print("\n\n")

	// THREADED TREES

	tthead := &ThreadedTree{info: "h"}
	tt0 := &ThreadedTree{info: "t0"}
	tt1 := &ThreadedTree{info: "t1"}
	tt2 := &ThreadedTree{info: "t2"}
	tt3 := &ThreadedTree{info: "t3"}
	tt4 := &ThreadedTree{info: "t4"}
	tt5 := &ThreadedTree{info: "t5"}

	// definition of head
	tthead.rtag = false
	tthead.rlink = tthead

	// regular links
	tthead.llink = tt0
	tt0.llink = tt1
	tt0.rlink = tt2

	tt1.llink = tt3
	tt1.rlink = tt4

	tt4.llink = tt5

	// threaded links
	tt3.rtag = true
	tt3.rlink = tt1

	tt5.rtag = true
	tt5.rlink = tt4

	tt4.rtag = true
	tt4.rlink = tt0

	tt2.rtag = true
	tt2.rlink = tthead

	fmt.Println("threaded traversal:")
	ThreadedTraversal(tthead)
	fmt.Print("\n\n")
}
