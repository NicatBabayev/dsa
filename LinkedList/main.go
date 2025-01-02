package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type SinglyLinkedList struct {
	head   *Node
	length int
}

func (l *SinglyLinkedList) listNodes() {
	currentNode := l.head
	listLength := l.length
	for listLength > 0 {
		fmt.Printf("Data of the Node: %d, Mem Address of the next node: %v\n", currentNode.data, &currentNode.next)
		currentNode = currentNode.next
		listLength--
	}
}

func (l *SinglyLinkedList) findByValue(data int) Node {
	currentNode := l.head
	listLength := l.length
	for listLength > 0 {
		if currentNode.data == data {
			return *currentNode
		}
		currentNode = currentNode.next
	}
	return Node{}
}

func (l *SinglyLinkedList) append(n Node) {
	currentNode := l.head
	listLength := l.length

	for listLength > 0 {
		if currentNode.next == nil {
			currentNode.next = &n
			l.length++
			return
		}
		currentNode = currentNode.next
		listLength--
	}
}

func (l *SinglyLinkedList) prepend(n Node) {
	second := l.head
	l.head = &n
	l.head.next = second
	l.length++
}

func (l *SinglyLinkedList) addAfterValue(addAfter int, node Node) {
	currentNode := l.head
	listLength := l.length
	for listLength > 0 {
		if currentNode.data == addAfter {
			nextNode := currentNode.next
			currentNode.next = &node
			currentNode.next.next = nextNode
			l.length++
			return
		}
		currentNode = currentNode.next
		listLength--
	}
}

func (l *SinglyLinkedList) addBeforeValue(addBefore int, node Node) {
	currentNode := l.head
	listLength := l.length
	for listLength > 0 {
		if currentNode.next.data == addBefore {
			nextNode := currentNode.next
			currentNode.next = &node
			currentNode.next.next = nextNode
			l.length++
			return
		}
		currentNode = currentNode.next
		listLength--
	}
}

func (l *SinglyLinkedList) deleteFromBeginning() {
	l.head = l.head.next
	l.length--
}

func (l *SinglyLinkedList) deleteFromEnd() {
	currentNode := l.head
	listLength := l.length
	for listLength > 0 {
		if currentNode.next.next == nil {
			currentNode.next = nil
			l.length--
			return
		}
		currentNode = currentNode.next
		listLength--
	}
}
func (l *SinglyLinkedList) deleteByValue(node Node) {
	currentNode := l.head
	listLength := l.length

	for listLength > 0 {
		if currentNode.next.data == node.data {
			if currentNode.next.next == nil {
				currentNode.next = nil
				l.length--
				return
			} else {
				currentNode.next = currentNode.next.next
				l.length--
				return
			}
		}
	}
	currentNode = currentNode.next
	listLength--
}

func main() {
	lList := SinglyLinkedList{}
	node1 := Node{data: 45}
	node2 := Node{data: 356}
	node3 := Node{data: 238}
	node4 := Node{data: 123213}
	node5 := Node{data: 11}
	node6 := Node{data: 77}
	lList.prepend(node1)
	lList.prepend(node2)
	lList.prepend(node3)
	lList.append(node4)
	lList.addAfterValue(356, node5)
	lList.addBeforeValue(356, node6)
	fmt.Println(lList)
	lList.deleteFromBeginning()
	lList.deleteFromEnd()
	lList.deleteByValue(node2)
	lList.listNodes()
}
