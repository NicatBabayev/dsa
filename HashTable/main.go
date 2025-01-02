package main

import "fmt"

const ArraySize = 100

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head   *bucketNode
	length int
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func hashSum(s string) int {
	var res int
	for _, i := range s {
		res += int(i)
	}
	return res % ArraySize
}

func (b *bucket) insert(bucketNode *bucketNode) {
	currentNode := b.head
	listLength := b.length
	for listLength > 0 {
		if currentNode.next == nil {
			currentNode.next = bucketNode
			b.length++
		}
		currentNode = currentNode.next
		listLength--
	}
}

func (b *bucket) search(bucketNode *bucketNode) *bucketNode {
	res := bucketNode
	currentNode := b.head
	listLength := b.length
	for listLength > 0 {
		if currentNode.key == bucketNode.key {
			return currentNode
		}
		currentNode = currentNode.next
		listLength--
	}

	return res
}

func (b *bucket) delete(bucketNode *bucketNode) {
	currentNode := b.head
	listLength := b.length
	for listLength > 0 {
		if currentNode.next.key == bucketNode.key {
			if currentNode.next == nil {
				currentNode.next = nil
				b.length--
			} else {
				currentNode.next = currentNode.next.next
				b.length--
			}
		}
		currentNode = currentNode.next
		listLength--
	}
}

func (h *HashTable) insert(s string) {
	index := hashSum(s)
	bucket := h.array[index]
	bucketNode := bucketNode{key: s}

	if present := bucket.search(&bucketNode); present == nil {
		bucket.insert(&bucketNode)
	}

}
func (h *HashTable) search(s string) bucketNode {
	res := bucketNode{}
	index := hashSum(s)
	bucket := h.array[index]
	bucketNode := bucketNode{key: s}

	if present := bucket.search(&bucketNode); present != nil {
		res = *present
	}
	return res
}
func (h *HashTable) delete(s string) {
	index := hashSum(s)
	bucket := h.array[index]
	bucketNode := bucketNode{key: s}

	if present := bucket.search(&bucketNode); present != nil {
		bucket.delete(&bucketNode)
	}
}

func (h *HashTable) list() {
	for index, val := range h.array {
		fmt.Println(index, ":", val)
	}
}

func InitHashTable() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	hashTable := InitHashTable()
	hashTable.insert("Hello")
	hashTable.list()
}
