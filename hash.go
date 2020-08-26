package main

import "fmt"

// SIZE holds the number of buckents in the hash table
const SIZE = 15

// Node represents one node in one has bucket
type Node struct {
	Value int
	Next  *Node
}

// HashTable represents full table of buckets
type HashTable struct {
	Table map[int]*Node
	Size  int
}

// Calculate hash table bucket index
func hashFunction(i, size int) int {
	return (i % size)
}

// Insert new node in hash table
func insertHash(hash *HashTable, value int) int {
	index := hashFunction(value, hash.Size)
	element := &Node{Value: value, Next: hash.Table[index]}
	hash.Table[index] = element
	return index
}

// Traverse passed hash table
func traverseHash(hash *HashTable) {
	for i := range hash.Table {
		if hash.Table[i] != nil {
			node := hash.Table[i]

			for node != nil {
				fmt.Printf("%d -> ", node.Value)
				node = node.Next
			}

			fmt.Println()
		}
	}
}
