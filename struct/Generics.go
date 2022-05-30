package main

import "fmt"

// K is comparable we can compare values of this type with the == and != operators , V can be any
// this is generic function
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

//list as generic type - represented as Linked List
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) push(value T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: value}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: value}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "One", 2: "Two", 3: "Three"}
	fmt.Println("Map of Keys:", MapKeys(m))

	llist := List[string]{}
	llist.push("One")
	llist.push("Two")
	llist.push("Three")
	fmt.Println("linkedList -->", llist.GetAll())
}
