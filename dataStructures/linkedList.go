package datastructures

import (
	"fmt"
)

// The implimentation of a doubly linked list

type LinkedList struct {
	Head, Tail *Node
	count      uint // current number of items in the list
}

func (l *LinkedList) Count() uint { return l.count } // making count readonly

type Node struct {
	Last, Next *Node
	Data       string
}

func (list *LinkedList) AddTail(data string) (ok bool) {

	item := &Node{Data: data}

	if list.Head == nil || list.count == 0 {
		// fmt.Println("Head empty, making item head.")
		item.Last = nil
		item.Next = nil
		list.Head = item
		list.Tail = item
		list.count = 1
		// fmt.Printf("Head: %+v\t%p\n", *list.Head, list.Head)
		return true
	}

	// update tail
	item.Last = list.Tail
	list.Tail.Next = item
	list.Tail = item
	list.count++
	return true

}

func (list *LinkedList) AddHead(data string) (ok bool) {

	item := &Node{Data: data}

	if list.Head == nil || list.count == 0 {
		// fmt.Println("Head empty, making item head.")
		item.Last = nil
		item.Next = nil
		list.Head = item
		list.Tail = item
		list.count = 1
		// fmt.Printf("Head: %+v\t%p\n", *list.Head, list.Head)
		return true
	}

	// update Head
	item.Next = list.Head
	list.Head.Last = item
	list.Head = item
	list.count++
	return true
}

func (list *LinkedList) Insert(data string, pos uint) (ok bool) {
	// Add an item to a linked list,
	// remove next and last references of the item if the list is empty.
	item := &Node{Data: data}

	if pos > list.count {
		// fmt.Printf("position not matching list count %t\n", (pos > list.count))
		return false
	}

	// TODO: fix adding to the head and tail.

	if list.Head == nil || list.count == 0 {
		// fmt.Println("Head empty, making item head.")
		item.Last = nil
		item.Next = nil
		list.Head = item
		list.Tail = item
		list.count = 1
		// fmt.Printf("Head: %+v\t%p\n", *list.Head, list.Head)
		return true
	}

	if pos == 0 {
		// update Head
		item.Next = list.Head
		list.Head.Last = item
		list.Head = item
		list.count++
		return true
	}

	if pos == list.count {
		// update tail
		item.Last = list.Tail
		list.Tail.Next = item
		list.Tail = item
		list.count++
		return true
	}

	// fmt.Printf("Iterating list for %d items\n", pos-1)
	var tmp *Node
	if pos > list.count/2 {
		// fmt.Println("iterating from end")
		tmp = list.Tail
		for i := list.count; i > pos; i-- {
			if tmp.Last == nil {
				break
			}
			tmp = tmp.Last
		}

	} else {

		// fmt.Println("iterating from start")
		tmp = list.Head
		for i := 0; i < int(pos)-1; i++ {
			if tmp.Next == nil {
				break
			}
			tmp = tmp.Next
		}
	}

	// fmt.Println("setting the next item")
	if tmp.Next != nil {
		next := tmp.Next
		// replace the next and move it forward
		next.Last = item
		item.Next = next
	}
	tmp.Next = item
	item.Last = tmp

	list.count++

	if pos == list.count-1 {
		list.Tail = item
	}

	return true
}

func (l *LinkedList) Print() {
	if l == nil || l.count == 0 {
		fmt.Println("Empy List")
		return
	}
	tmp := l.Head

	if *tmp == (Node{}) {
		fmt.Println("Empy List")
		return
	}
	fmt.Printf("Count: %d\n", l.count)
	fmt.Printf("Head: %p\tTail: %p\n", l.Head, l.Tail)

	for i := 0; i < int(l.count); i++ {
		fmt.Printf("%d\t%p\t%+v\n", i, tmp, *tmp)
		if tmp.Next == nil {
			fmt.Println("End of List.")
			break
		}
		tmp = tmp.Next
	}
}
