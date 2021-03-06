package main

import (
	"fmt"
)

func main() {
	tree := create(10)

	traverseTree(tree)
	fmt.Println()

	tree = insertTree(tree, -10)
	tree = insertTree(tree, -2)

	traverseTree(tree)

	table := make(map[int]*Node, SIZE)
	hash := &HashTable{Table: table, Size: SIZE}
	fmt.Println("Number of spaces: ", hash.Size)

	for i := 0; i < 120; i++ {
		insertHash(hash, i)
	}

	traverseHash(hash)

	if lookupHash(hash, 120) {
		fmt.Println("120 was found in hash")
	}

	if lookupHash(hash, 100) {
		fmt.Println("100 was found in hash")
	}
}
