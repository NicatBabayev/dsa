package main

import "fmt"

const ArraySize = 10

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

func (b *bucket) insert(node *bucketNode) {
	currentNode := b.head
	listLength := b.length
	if b.length == 0 {
		b.head = node
		b.length++
		return
	}

	for listLength > 0 {
		if currentNode.next == nil {

			b.head.next = node
			b.length++
		}
		currentNode = currentNode.next
		listLength--
	}
}

func (b *bucket) search(node *bucketNode) *bucketNode {
	res := bucketNode{}
	currentNode := b.head
	listLength := b.length
	for listLength > 0 {
		if currentNode.key == node.key {
			return currentNode
		}
		currentNode = currentNode.next
		listLength--
	}

	return &res
}

func (b *bucket) delete(node *bucketNode) {
	currentNode := b.head
	listLength := b.length

	if currentNode.key == node.key {
		b.head = b.head.next
		b.length--
		return
	}

	for listLength > 0 {
		if currentNode.next.key == node.key {
			if currentNode.next.next == nil {
				currentNode.next = nil
				b.length--
				return
			} else {
				currentNode.next = currentNode.next.next
				b.length--
				return
			}
		}
		currentNode = currentNode.next
		listLength--
	}
}

func (b *bucket) list() {
	currentNode := b.head
	listLength := b.length

	for listLength > 0 {
		fmt.Println("\t", currentNode.key)
		listLength--
	}
}

func (h *HashTable) insert(s string) {
	emptyBucket := bucketNode{}
	index := hashSum(s)
	bucket := h.array[index]
	node := bucketNode{key: s}
	present := bucket.search(&node)
	if *present == emptyBucket {
		bucket.insert(&node)
	}
}

func (h *HashTable) search(s string) bucketNode {
	res := bucketNode{}
	index := hashSum(s)
	bucket := h.array[index]
	node := bucketNode{key: s}
	present := bucket.search(&node)

	if *present != res {
		res = *present
	}
	return res
}
func (h *HashTable) delete(s string) {
	emptyBucket := bucketNode{}
	index := hashSum(s)
	bucket := h.array[index]
	node := bucketNode{key: s}

	present := bucket.search(&node)
	if *present != emptyBucket {
		bucket.delete(&node)
	}
}

func (h *HashTable) list() {
	for index, val := range h.array {
		if val.length != 0 {
			fmt.Println(index, "bucket values:")
			val.list()
		}
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
	hashTable.insert("Hellowa")
	hashTable.insert("Helloadsfa")
	hashTable.insert("Heasdfa")
	hashTable.insert("Eric")
	hashTable.insert("cirE")
	hashTable.delete("Eric")
	hashTable.list()
}
