package NodeList

import "fmt"

// NodeList represents a singly-linked list that holds
// values of any type.
type NodeList[T comparable] struct {
	Next *NodeList[T]
	Val  T
}

func (ll NodeList[T]) Find(val T) int {
	cur := ll
	i := 0
	for cur.Val != val && cur.Next != nil {
		cur = *(cur.Next)
		i++
	}
	if cur.Val == val {
		return i
	}
	return -1
}

func (ll *NodeList[T]) Append(value T) {
	currentNode := ll

	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	newVal := NodeList[T]{Val: value}
	currentNode.Next = &newVal
}

func (ll NodeList[T]) String() string {
	str := "["
	currentNode := ll

	for currentNode.Next != nil {
		str += fmt.Sprintf("%v, ", currentNode.Val)
		currentNode = *currentNode.Next
	}
	str += fmt.Sprintf("%v]", currentNode.Val)
	return str
}
