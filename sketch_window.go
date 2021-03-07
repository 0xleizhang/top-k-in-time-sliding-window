package top_k

import (
	"container/heap"
	"github.com/seven4x/top-k/pkg"
	cms "github.com/shenwei356/countminsketch"
	"time"
)

// 秒级窗口，超过窗口生成窗口概要
type SketchWindow struct {
	sketch      *cms.CountMinSketch
	minHeap     *pkg.MinHeap
	start       time.Time
	end         time.Time
	topKeepSize int
	timeSize    int //秒
}

func newSketchWindow(topKeepSize, timeSize int) *SketchWindow {
	epsilon, delta := 0.0001, 0.9999
	sketch, _ := cms.NewWithEstimates(epsilon, delta)

	h := &pkg.MinHeap{}
	heap.Init(h)

	window := &SketchWindow{
		sketch:      sketch,
		minHeap:     h,
		start:       time.Time{},
		topKeepSize: topKeepSize,
		timeSize:    timeSize,
	}

	return window
}

func (w *SketchWindow) Push(word string, timestamp time.Time) {
	if w.start.IsZero() {
		w.start = timestamp
	} else {
		w.end = timestamp
	}
	w.sketch.UpdateString(word, 1)
	count := int(w.sketch.EstimateString(word))
	pair := pkg.Word{
		Key:   word,
		Count: count,
	}
	h := w.minHeap
	h.LimitPush(pair, w.topKeepSize)
}

func (w *SketchWindow) IsFull(timestamp time.Time) bool {
	if w.start.IsZero() {
		return false
	}
	d := timestamp.Sub(w.start)
	return int(d.Seconds()) > w.timeSize
}

func (w *SketchWindow) BuildSynopsis() pkg.Synopsis {
	s := pkg.Synopsis{}
	s.Start = w.start
	s.End = w.end
	h := w.minHeap
	s.Top = *h
	return s
}

func (w *SketchWindow) Clear() {
	epsilon, delta := 0.0001, 0.9999
	sketch, _ := cms.NewWithEstimates(epsilon, delta)
	h := &pkg.MinHeap{}
	heap.Init(h)
	w.start = time.Time{}
	w.end = time.Time{}
	w.minHeap = h
	w.sketch = sketch
}
