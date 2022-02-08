package main

import "fmt"

type Tree struct {
	llink *Tree
	info  string
	rlink *Tree
}

func PreOrder(t *Tree) {
	if t == nil {
		return
	}
	fmt.Printf("%s ", t.info)
	PreOrder(t.llink)
	PreOrder(t.rlink)
}

func InOrder(t *Tree) {
	if t == nil {
		return
	}
	InOrder(t.llink)
	fmt.Printf("%s ", t.info)
	InOrder(t.rlink)
}

func PostOrder(t *Tree) {
	if t == nil {
		return
	}
	PostOrder(t.llink)
	PostOrder(t.rlink)
	fmt.Printf("%s ", t.info)
}
