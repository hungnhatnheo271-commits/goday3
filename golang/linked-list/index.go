package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Prepend(value int) {
	newNode := &Node{data: value, next: ll.head}
	ll.head = newNode
}

func (ll *LinkedList) Append(value int) {
	newNode := &Node{data: value}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *LinkedList) PrintList() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}

	// 10 -> 20 -> nil
	fmt.Println("nil")
}

func (ll *LinkedList) Search(value int) bool {
	current := ll.head
	for current.next != nil {
		if current.data == value {
			return true
		}
		current = current.next
	}
	return false
}

func (ll *LinkedList) Delete(value int) {
	if ll.head == nil {
		return
	}
	if ll.head.data == value {
		ll.head = ll.head.next
		return
	}
	current := ll.head
	for current.next != nil && current.next.data != value {
		current = current.next
	}
	if current.next != nil {
		current.next = current.next.next
	}
}

func main() {
	ll := LinkedList{}
	ll.Append(10)
	ll.Append(20)
	ll.Append(30)
	ll.Append(40)
	ll.Append(50)
	ll.PrintList()
	ll.Delete(20)
	ll.PrintList()
	ll.Prepend(5)
	ll.PrintList()
	fmt.Println("Searching 10...", ll.Search(10)) // true
	fmt.Println("Searching 30...", ll.Search(30)) // false
	ll.InsertAt(5,999)
	ll.PrintList()
}

// EXCERCISE: 1. Viết method InsertAt() để chèn vào bất kỳ vị trí nào mình muốn
func (ll *LinkedList) InsertAt(position int, value int) {
	// Tạo new node

	// Tìm position
	// Giả sử có List: 10 -> 20 -> 30 -> 40 -> 50 -> nil
	// Chèn số 9999 vào vị trí 3 -> 10 -> 20 -> 9999 -> 30 -> 40 -> 50 -> nil
	newNode := &Node{data: value}

	// nếu chèn đầu
	if position == 1 {
		newNode.next = ll.head
		ll.head = newNode
		return
	}

	current := ll.head
	index := 1

	for current != nil && index < position-1 {
		current = current.next
		index++
	}

	if current == nil {
		fmt.Println("Position out of range")
		return
	}

	newNode.next = current.next
	current.next = newNode
}

// EXCERCISE: 2. Thay đổi toàn bộ data type thành Generic Types !
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	val  T
	next *element[T]
}
