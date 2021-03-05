package top_k

import (
	"container/heap"
	"github.com/seven4x/top-k/pkg"
	"time"
)

type Window struct {
	queue       *pkg.CircularQueue
	minheap     *pkg.MinHeap
	topKeepSize int
	timeSize    int //秒
}

func NewWindow(qsize, topKeepSize, timeSize int) *Window {
	h := &pkg.MinHeap{}
	heap.Init(h)
	return &Window{
		queue:       pkg.NewCircularQueue(qsize),
		minheap:     h,
		topKeepSize: topKeepSize,
		timeSize:    timeSize,
	}
}

func (w *Window) IsFull(timestamp time.Time) bool {
	if w.queue.IsEmpty() {
		return false
	}
	front := w.queue.Front()
	d := timestamp.Sub(front.(pkg.Synopsis).Start)
	isFull := int(d.Seconds()) > w.timeSize
	return isFull
}

func (w *Window) IsOutbound(timestamp time.Time) bool {
	rear := w.queue.Rear()
	if rear == nil {
		return false
	}
	d := timestamp.Sub(rear.(pkg.Synopsis).Start)
	return int(d.Seconds()) > w.timeSize
}

//满则先出再进
func (w *Window) Sliding(syn pkg.Synopsis) {
	isFull := w.queue.IsFull()
	if isFull {
		head := w.queue.DeQueue()
		words := head.(pkg.Synopsis).Top
		for _, word := range words {
			w.out(word)
		}
	}
	w.queue.EnQueue(syn)
	for _, word := range syn.Top {
		w.in(word)
	}

}

// 查找word，更新加 重新排序
func (w *Window) in(word pkg.Word) {
	p := w.minheap.Search(word)
	if p > 0 {
		(*w.minheap)[p].Count += word.Count
		heap.Fix(w.minheap, p)
	} else {
		if w.minheap.Len() < w.topKeepSize {
			heap.Push(w.minheap, word)
		} else {
			top := w.minheap.Top().(pkg.Word)
			if word.Count > top.Count {
				(*w.minheap)[0] = word
				heap.Fix(w.minheap, 0)
			}
		}
	}
}

func (w *Window) out(word pkg.Word) {
	p := w.minheap.Search(word)
	if p < 0 {
		return
	}
	(*w.minheap)[p].Count -= word.Count
	heap.Fix(w.minheap, p)
}

func (w *Window) BuildSynopsis() pkg.Synopsis {
	s := pkg.Synopsis{}
	front := w.queue.Front()
	s.Start = front.(pkg.Synopsis).Start
	s.Top = *(w.minheap)
	return s
}

func (w *Window) Clear() {
	w.queue.Clear()
	h := &pkg.MinHeap{}
	heap.Init(h)
	w.minheap = h
}
