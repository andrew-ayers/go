// from https://gist.github.com/bemasher/1777766
package myheap

type Heap struct {
	top  *Element
	size int
}

type Element struct {
	value interface{} // All types satisfy the empty interface, so we can store anything here.
	next  *Element
}

// Return the heap's length
func (h *Heap) Len() int {
	return h.size
}

// Push a new element onto the heap
func (h *Heap) Push(value interface{}) {
	h.top = &Element{value, h.top}
	h.size++
}

// Remove the top element from the heap and return it's value
// If the heap is empty, return nil
func (h *Heap) Pop() (value interface{}) {
	if h.size > 0 {
		value, h.top = h.top.value, h.top.next
		h.size--
		return
	}

	return nil
}
