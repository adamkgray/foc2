package main

import "fmt"

func main() {
	t2l := &Tree{info: "t2l"}
	t2r := &Tree{info: "t2r"}
	t := &Tree{info: "root", llink: t2l, rlink: t2r}

	fmt.Println("pre-order:")
	PreOrder(t)
	fmt.Print("\n")

	fmt.Println("in-order:")
	InOrder(t)
	fmt.Print("\n")

	fmt.Println("post-order:")
	PostOrder(t)
	fmt.Print("\n")
}
