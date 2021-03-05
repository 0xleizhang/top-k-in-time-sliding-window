package pkg

import (
	"container/heap"
	"fmt"
)

//堆
type MinHeap []Word

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i].Count < h[j].Count }

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(word interface{}) {
	*h = append(*h, word.(Word))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MinHeap) Top() interface{} {
	return h[0]
}

func (h *MinHeap) LimitPush(word Word, limit int) {
	i := h.Search(word)
	l := h.Len()

	if l < limit { //未超
		if i == -1 { //未找到
			heap.Push(h, word)
		} else {
			(*h)[i] = word
			heap.Fix(h, i)
		}
	} else {
		if i == -1 {
			top := h.Top().(Word)
			if word.Count > top.Count {
				(*h)[0] = word
				heap.Fix(h, 0)
			}
		} else {
			(*h)[i] = word
			heap.Fix(h, i)
		}
	}

}

func (h *MinHeap) Search(word Word) int {
	for i, w := range *h {
		if w.Key == word.Key {
			return i
		}
	}
	return -1
}

func (h *MinHeap) Print() {
	fmt.Printf("%v \n", *h)
}
