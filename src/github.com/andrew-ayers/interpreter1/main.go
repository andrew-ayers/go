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
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-= LINKED LIST DEMO =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

	list := new(mylinkedlist.List)

	//fmt.Println(list)

	//fmt.Println(list.Len())

	for i := 0; i < 10; i++ {
		list.Add("yolo-" + strconv.Itoa(i))
	}

	//fmt.Println(list)

	//fmt.Println(list.Len())

	v := list.Head()

	for {
		if v == nil {
			break
		}

		prev := v.(*mylinkedlist.Element).Prev
		curr := v
		next := v.(*mylinkedlist.Element).Next

		fmt.Println("PREV:", prev, "CURR:", curr, "NEXT:", next)

		v = list.Next()
	}
}
