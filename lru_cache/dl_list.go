package lru_cache

import "fmt"

type NodeData struct {
	Key   string
	Value *string
}

type Node struct {
	Data *NodeData
	Prev *Node
	Next *Node
}

func NewNode(key, data *string) *Node {
	var node Node

	if key == nil && data == nil {
		node.Data = nil
	} else {
		node.Data = &NodeData{
			Key:   *key,
			Value: data,
		}
	}

	node.Prev = nil
	node.Next = nil

	return &node
}

type DoublyLinkedList struct {
	Header  *Node
	Trailer *Node
	Size    int
}

func NewDoublyLinkedList(firstNode *Node) *DoublyLinkedList {

	var list DoublyLinkedList

	// empty nodes as header and trailer
	list.Header = NewNode(nil, nil)
	list.Trailer = NewNode(nil, nil)

	list.Header.Prev = nil
	list.Trailer.Next = nil

	if firstNode == nil {
		list.Header.Next = list.Trailer
		list.Trailer.Prev = list.Header
		return &list
	}

	list.Header.Next = firstNode
	list.Trailer.Prev = firstNode
	list.Size++

	return &list
}

func (list *DoublyLinkedList) InsertAtHead(key, value *string) *Node {

	// new node with the data
	node := NewNode(key, value)
	pastHeadRef := list.Header.Next

	list.Size++

	// if the list was empty at the time of insertion
	if pastHeadRef == list.Trailer {
		list.Header.Next = node
		list.Trailer.Prev = node
		node.Prev = list.Header
		node.Next = list.Trailer

		return node
	}

	node.Prev = list.Header
	node.Next = pastHeadRef

	list.Header.Next = node
	pastHeadRef.Prev = node

	return node
}

func (list *DoublyLinkedList) RemoveAtTail() (*Node, error) {

	nodeRef := list.Trailer.Prev

	// trying to delete an empty list results in an error
	if nodeRef == list.Header {
		return nil, fmt.Errorf("cannot delete an empty linked list")
	}

	nodeRef.Prev.Next = list.Trailer
	list.Trailer.Prev = nodeRef.Prev

	list.Size--

	return nodeRef, nil
}

func (list *DoublyLinkedList) PrintList() {

	head := list.Header.Next

	fmt.Printf("Printing LinkedList: ")
	for head != list.Trailer {
		fmt.Printf("{%s: %s} ", head.Data.Key, *head.Data.Value)
		head = head.Next
	}
	fmt.Printf("\n")
}

// given a reference to a node inside the list, reallocate it to the beginning
func (list *DoublyLinkedList) ReInsertAtHead(nodeRef *Node) {

	prev := nodeRef.Prev
	next := nodeRef.Next

	// if the element is already first or single-element list have no effect
	if prev == list.Header || (prev == list.Header && next == list.Trailer) {
		return
	}

	nodeRef.Prev = list.Header
	nodeRef.Next = list.Header.Next

	list.Header.Next.Prev = nodeRef
	list.Header.Next = nodeRef

	prev.Next = next
	next.Prev = prev

}
