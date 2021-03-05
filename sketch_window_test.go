package top_k

import (
	"github.com/bmizerany/assert"
	"testing"
	"time"
)

func TestSketchWindow_BuildSynopsis(t *testing.T) {
	win := newSketchWindow(3, 60)
	win.Push("a", parseTime("2021-03-05 14:59:01"))
	win.Push("a", parseTime("2021-03-05 14:59:02"))
	win.Push("b", parseTime("2021-03-05 14:59:02"))
	win.Push("b", parseTime("2021-03-05 14:59:03"))
	win.Push("b", parseTime("2021-03-05 14:59:04"))
	win.Push("f", parseTime("2021-03-05 14:59:04"))
	win.Push("d", parseTime("2021-03-05 14:59:04"))

	assert.Equal(t, 3, int(win.sketch.EstimateString("b")))
	assert.Equal(t,true,win.IsFull(parseTime("2021-03-05 15:00:04")))
	snno := win.BuildSynopsis()
	assert.Equal(t,"b",snno.Top[2].Key)
	win.Clear()
	assert.Equal(t,true,win.start.IsZero())
}

func parseTime(str string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", str)
	return t
}
