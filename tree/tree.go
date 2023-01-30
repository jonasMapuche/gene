package trees

import (
	"time"
)

type Notice struct {
	Spawn    string
	Date     time.Time
	Quantity float32
}

type Tree struct {
	left  *Tree
	value Notice
	right *Tree
}

var note *Tree

func Add(article Notice, arbor *Tree) *Tree {
	note = insert(article, arbor)
	return note
}

func insert(article Notice, arbor *Tree) *Tree {
	if arbor == nil {
		return &Tree{nil, article, nil}
	}
	return arbor
}
