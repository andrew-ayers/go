package main

import (
	"fmt"
	"github.com/andrew-ayers/helpers/myheap"
	"github.com/andrew-ayers/helpers/mylinkedlist"
	"strconv"
)

func main() {
	demoHeap()
	demoLinkedList()
}

func demoHeap() {
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-= HEAP DEMO =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

	heap := new(myheap.Heap)

	heap.Push(1234)
	heap.Push("Things")
	heap.Push("and")
	heap.Push("Stuff")

	for heap.Len() > 0 {
		// We have to do a type assertion because we get back a variable of type
		// interface{} while the underlying type is a string.
		//fmt.Printf("%s ", heap.Pop().(string))
		fmt.Println(heap.Pop())
	}

	fmt.Println()
}

func demoLinkedList() {
	// Initialize the linked list
	list := new(mylinkedlist.List)

	for i := 0; i < 10; i++ {
		list.Add("yolo-" + strconv.Itoa(i))
	}

	// Iterate over list and display, from head to tail
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-= LINKED LIST DEMO - HEAD2TAIL =-=-=-=-=-=-=-=-=-=-=-=-=-=")

	v := list.Head()

	for {
		if v == nil {
			break
		}

		prev := "\t"
		if v.(*mylinkedlist.Element).Prev != nil {
			prev = v.(*mylinkedlist.Element).Prev.Value.(string)
		}

		//prev := v.(*mylinkedlist.Element).Prev != nil ? v.(*mylinkedlist.Element).Prev.Value : nil
		curr := v.(*mylinkedlist.Element).Value

		next := "\t"
		if v.(*mylinkedlist.Element).Next != nil {
			next = v.(*mylinkedlist.Element).Next.Value.(string)
		}

		fmt.Printf("PREV: %s\tCURR: %s\tNEXT: %s\n", prev, curr, next)

		v = list.Next()
	}

	// Iterate over list and display, from tail to head
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-= LINKED LIST DEMO - TAIL2HEAD =-=-=-=-=-=-=-=-=-=-=-=-=-=")

	v = list.Tail()

	for {
		if v == nil {
			break
		}

		prev := "\t"
		if v.(*mylinkedlist.Element).Prev != nil {
			prev = v.(*mylinkedlist.Element).Prev.Value.(string)
		}

		//prev := v.(*mylinkedlist.Element).Prev != nil ? v.(*mylinkedlist.Element).Prev.Value : nil
		curr := v.(*mylinkedlist.Element).Value

		next := "\t"
		if v.(*mylinkedlist.Element).Next != nil {
			next = v.(*mylinkedlist.Element).Next.Value.(string)
		}

		fmt.Printf("NEXT: %s\tCURR: %s\tPREV: %s\n", next, curr, prev)

		v = list.Previous()
	}

	v = list.Tail()
	fmt.Println(list.Position())

	v = list.Head()
	fmt.Println(list.Position())
}
