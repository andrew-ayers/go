package mylinkedlist

import (
	//"fmt"
	"reflect"
)

type List struct {
	head   *Element
	curr   *Element
	tail   *Element
	length int
}

type Element struct {
	Value interface{} // All types satisfy the empty interface, so we can store anything here.
	Prev  *Element
	Next  *Element
}

/*
func (l *List) Init(elements int, value interface{}) int {
}
*/

func (l *List) Next() (value interface{}) {
	if l.length > 0 && l.curr.Next != nil {
		l.curr = l.curr.Next

		value = l.curr

		return
	}

	return nil
}

func (l *List) Previous() (value interface{}) {
	if l.length > 0 && l.curr.Prev != nil {
		l.curr = l.curr.Prev

		value = l.curr

		return
	}

	return nil
}

func (l *List) Head() (value interface{}) {
	if l.length > 0 {
		l.curr = l.head

		value = l.curr

		return
	}

	return nil
}

func (l *List) Tail() (value interface{}) {
	if l.length > 0 {
		l.curr = l.tail

		value = l.curr

		return
	}

	return nil
}

// Returns the value of the element at the current location
func (l *List) Current() (value interface{}) {
	if l.length > 0 {
		value = l.curr

		return
	}

	return nil
}

// Return the list's length
func (l *List) Len() int {
	return l.length
}

// Add an entry to the end of the list
func (l *List) Add(value interface{}) int {
	if l.length == 0 {
		l.curr = &Element{value, nil, nil}
		l.head = l.curr
		l.tail = l.curr
	} else {
		l.tail = &Element{value, l.tail, nil}
		l.curr.Next = l.tail
		l.curr = l.tail
	}

	l.length++

	return l.length
}

// ****Insert an entry at the current location
func (l *List) Insert(value interface{}) int {
	return 0
}

// Update the entry at the cursor location
func (l *List) Update(value interface{}) (ret interface{}) {
	l.curr = &Element{value, l.curr.Prev, l.curr.Next}

	if l.length > 0 {
		ret = l.curr
		return
	}

	return nil
}

// ****Delete the entry at the cursor location
func (l *List) Delete() int {
	if l.length > 0 {
	}

	return 0
}

/*
func (l *List) Reverse() int {
}

func (l *List) Shuffle(rounds int) int {
}

func (l *List) Fill(value interface{}) int {
}
*/

func (l *List) Position() int {
	pos := 0

	check := l.curr

	if l.length > 0 {
		v := l.Head()

		for {
			if v == nil {
				break
			}

			if reflect.DeepEqual(check, v) == true {
				l.curr = check

				return pos
			}

			v = l.Next()

			pos++
		}
	}

	return -1
}
