package trees

import (
	"fmt"
	"time"
)

type Notice struct {
	Spawn    string
	Date     time.Time
	Quantity float32
}

type Tree struct {
	high  *Tree
	level int
	value Notice
	low   *Tree
}

func Add(article Notice) *Tree {
	var arbor *Tree
	for level := 1; level <= 10; level++ {
		arbor = insert(1, article, arbor, arbor)
	}
	return arbor
}

func insert(level int, article Notice, arbor *Tree, tree *Tree) *Tree {
	if arbor == nil {
		return &Tree{nil, level, article, nil}
	}
	if arbor.high == nil {
		arbor.high = insert(level+1, article, arbor.high, arbor)
	} else {
		if arbor.low == nil {
			arbor.low = insert(level+1, article, arbor.low, arbor)
		} else {
			arbor.high = insert(level+1, article, arbor.high, arbor)
		}
	}
	return arbor
}

func complete(level int, arbor *Tree) {
	var plump = 1
	for {
		plump++
		if plump == level {
			break
		}
	}
}

func (arbor *Tree) String() string {
	if arbor == nil {
		return "()"
	}
	text := ""
	if arbor.high != nil {
		text += arbor.high.String() + " "
	}
	text += fmt.Sprint(arbor.level)
	if arbor.low != nil {
		text += arbor.low.String() + " "
	}
	return "(" + text + ")"

}

/*
func find(article Notice, arbor *Tree) *Tree {
	if arbor == nil {
		arbor = insert(article, arbor)
		return arbor
	}
	if arbor.left == nil {

	}
	arbor = find(article, arbor.left)
	arbor = find(article, arbor.right)
}
*/
