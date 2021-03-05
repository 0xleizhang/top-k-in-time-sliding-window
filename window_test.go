package top_k

import (
	"fmt"
	"github.com/bmizerany/assert"
	"github.com/seven4x/top-k/pkg"
	"testing"
)

func TestWindow_BuildSynopsis(t *testing.T) {
	win := NewWindow(3, 3, 120)

	syno := pkg.Synopsis{
		Start: parseTime("2021-03-05 12:00:00"),
		Top:   []pkg.Word{{Key: "a", Count: 50}, {"b", 30}},
	}

	win.Sliding(syno)

	syno = pkg.Synopsis{
		Start: parseTime("2021-03-05 12:01:00"),
		Top:   []pkg.Word{{Key: "a", Count: 10}, {Key: "f", Count: 20}, {"b", 30}},
	}
	win.Sliding(syno)

	syno = pkg.Synopsis{
		Start: parseTime("2021-03-05 12:03:01"),
		Top:   []pkg.Word{{Key: "e", Count: 20}, {"b", 35}, {Key: "f", Count: 3}},
	}
	win.Sliding(syno)

	syno = pkg.Synopsis{
		Start: parseTime("2021-03-05 12:04:02"),
		Top:   []pkg.Word{{Key: "k", Count: 26}, {"b", 30}},
	}
	win.Sliding(syno)
	top := win.BuildSynopsis()
	max := top.Top[0]
	for _, v := range top.Top {
		if v.Count > max.Count {
			max = v
		}
	}
	fmt.Printf("%v \n", top.Top)
	assert.Equal(t, "b", max.Key)

	b := win.IsOutbound(parseTime("2021-03-05 12:05:06"))
	assert.Equal(t, false, b)
	b  = win.IsOutbound(parseTime("2021-03-05 12:07:03"))
	assert.Equal(t, true, b)

}
