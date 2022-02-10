package main

import "fmt"

type Tree struct {
	llink *Tree
	info  string
	rlink *Tree
}

type ThreadedTree struct {
	llink *ThreadedTree
	//ltag  bool // (threaded only) denotes if left pointer points to predecessor (in-order)
	info  string
	rlink *ThreadedTree
	rtag  bool // (threaded only) denotes if right pointer points to successor (in-order)
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

func ThreadedTraversal(t *ThreadedTree) {
	// seek leftmost
	for t.llink != nil {
		t = t.llink
	}
	for t != nil && !t.isHead() {
		fmt.Printf("%s ", t.info)
		if t.rtag {
			t = t.rlink
		} else {
			// seek leftmost of right link
			t = t.rlink
			for t.llink != nil {
				t = t.llink
			}
		}
	}

}

func (t *ThreadedTree) isHead() bool {
	// By convention, HEAD.RLINK = HEAD and HEAD.RTAG = 0
	return !t.rtag && t.rlink == t
}

func (t *Tree) InsertRight(p *Tree) {
	if t.rlink == nil {
		t.rlink = p
	} else {
		rlink := t.rlink
		t.rlink = p
		p.rlink = rlink
	}
}

func (t *Tree) InsertLeft(p *Tree) {
	if t.llink == nil {
		t.rlink = p
	} else {
		llink := t.llink
		t.llink = p
		p.rlink = llink
	}
}
