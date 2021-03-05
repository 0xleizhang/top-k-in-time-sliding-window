package pkg

import (
	"container/heap"
	"fmt"
	"github.com/bmizerany/assert"
	"testing"
)

func TestMinHeap(t *testing.T) {
	h := &MinHeap{}
	heap.Init(h)
	push("e2", 77, h)
	push("a", 20, h)
	push("b", 30, h)
	push("c", 40, h)
	push("c", 50, h)
	push("e", 23, h)
	push("f", 21, h)
	push("f3", 5, h)

	assert.Equal(t, 8, h.Len())
	assert.Equal(t, 5, heap.Pop(h).(Word).Count)

	i := h.Search(Word{
		Key:   "a",
		Count: 0,
	})
	assert.Equal(t, 0, i)
	heap.Remove(h, 0)
	i = h.Search(Word{Key: "a"})
	assert.Equal(t, -1, i)
	p := h.Search(Word{Key: "e"})
	(*h)[p] = Word{"e", 1000}
	heap.Fix(h, p)
	p2 := h.Search(Word{Key: "e"})
	assert.NotEqual(t, p, p2)
}
func push(key string, count int, h *MinHeap) {
	w := Word{
		Key:   key,
		Count: count,
	}
	heap.Push(h, w)
}

func TestMinHeapUser(t *testing.T) {
	h := MinHeap{}
	heap.Init(&h)
	heap.Push(&h, Word{
		Key:   "22",
		Count: 2,
	})
	top := h.Top()
	assert.Equal(t, "22", top.(Word).Key)
	assert.Equal(t, 1, h.Len())
	(h)[0].Count = 33
	assert.Equal(t, 33, heap.Pop(&h).(Word).Count)
}

func TestPrint(t *testing.T) {
	s := []Word{Word{"a", 2}, {"b", 33}}

	fmt.Printf("%v", s)
	fmt.Printf("%v", s)

}
