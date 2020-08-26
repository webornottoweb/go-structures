package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Tree represents
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func traverseTree(t *Tree) {
	if t == nil {
		return
	}

	traverseTree(t.Left)
	fmt.Print(t.Value, " ")
	traverseTree(t.Right)
}

func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())

	for i := 0; i < n*2; i++ {
		temp := rand.Intn(n * 2)
		t = insertTree(t, temp)
	}

	return t
}

func insertTree(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}

	if t.Value == v {
		return t
	}

	if t.Value > v {
		t.Left = insertTree(t.Left, v)
		return t
	}

	t.Right = insertTree(t.Right, v)
	return t
}
