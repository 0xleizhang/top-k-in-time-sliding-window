package top_k

import (
	"errors"
	"fmt"
	"github.com/seven4x/top-k/pkg"
	"time"
)

type (
	Handler interface {
		Consume(topic string, timestamp time.Time)
		GetTop(k int) []string
	}
	Config struct {
		TopKeepSize       []int
		SubWindowTimeSize []int
		WindowTimeSize    []int
	}
)

const (
	TimeStampLayout = "2006-01-02 15:04:05"
)

type MyHandler struct {
	config       Config
	sketchWindow *SketchWindow
	window       []*Window
}

func newDefaultConfig() Config {

	return Config{
		TopKeepSize:       []int{100, 100, 100},
		SubWindowTimeSize: []int{60, 30 * 60},        //第一个窗口1分钟滑动一次 第二个窗口30分钟滑动一次
		WindowTimeSize:    []int{1800, 24 * 60 * 60}, //第一个窗口大小30分钟，第二个窗口大小24小时
	}
}

func NewHandler() (*MyHandler, error) {
	return NewHandlerCustom(func(config *Config) {
		return
	})
}

func validateConfig(conf Config) error {
	if len(conf.SubWindowTimeSize) != len(conf.WindowTimeSize) || len(conf.SubWindowTimeSize) < 1 {
		return errors.New("error，check config")
	}
	if conf.SubWindowTimeSize[0] <= 0 {
		return errors.New("error,BasicTimeWindowSize >0 suggest 60")
	}
	if len(conf.TopKeepSize)-1 != len(conf.WindowTimeSize) {
		return errors.New("error,check TopKeepSize  [sketchWin,win1,win2]")
	}
	for i := range conf.TopKeepSize {
		if conf.TopKeepSize[i] <= 0 {
			return errors.New("config error,TopKeepSize must >0")
		}
	}
	return nil
}

func NewHandlerCustom(setting func(config *Config)) (*MyHandler, error) {
	conf := newDefaultConfig()
	setting(&conf)
	err := validateConfig(conf)
	if err != nil {
		return nil, err
	}
	handler := &MyHandler{config: conf}
	handler.sketchWindow = newSketchWindow(conf.TopKeepSize[0], conf.SubWindowTimeSize[0])
	expandSize := len(conf.WindowTimeSize)
	window := make([]*Window, expandSize, expandSize)
	for i := range conf.WindowTimeSize {
		qsize := conf.WindowTimeSize[i] / conf.SubWindowTimeSize[i]
		win := NewWindow(qsize, conf.TopKeepSize[i+1], conf.WindowTimeSize[i])
		window[i] = win
	}
	handler.window = window
	return handler, nil
}

func (h *MyHandler) Consume(topic string, timestamp time.Time) {
	sWin := h.sketchWindow
	if sWin.IsFull(timestamp) {
		synopsis := sWin.BuildSynopsis()
		h.refresh(synopsis, 0)
		sWin.Clear()
	}
	sWin.Push(topic, timestamp)
}

/**
滑动，小带大
*/
func (h *MyHandler) refresh(synopsis pkg.Synopsis, position int) {
	winSize := len(h.window)
	if position >= winSize {
		return
	}
	current := h.window[position]
	if position == winSize-1 { //最后一个窗口不用生成摘要，如果前后两个数据超过了时间窗口需要清空再添加
		if current.IsOutbound(synopsis.Start) {
			current.Clear()
		}
		current.Sliding(synopsis)
	} else {
		if current.IsOutbound(synopsis.Start) {
			current.Clear()
			current.Sliding(synopsis)
		} else if current.IsFull(synopsis.Start) {
			syn := current.BuildSynopsis()
			current.Clear()
			current.Sliding(synopsis)
			h.refresh(syn, position+1)
		} else {
			//正常更多是这
			current.Sliding(synopsis)
		}
	}

}

func (h *MyHandler) GetTop(k int) []pkg.Word {
	for i := len(h.window) - 1; i >= 0; i-- {
		if h.window[i].queue.IsEmpty() {
			continue
		}
		return *(h.window[i].minheap)
	}
	return nil
}

func (h *MyHandler) Print() {
	fmt.Print("-----------------------\n")
	fmt.Printf("sketchWindow in %ds, %s~%s,top: \n",
		h.config.SubWindowTimeSize[0],
		h.sketchWindow.start.Format(TimeStampLayout),
		h.sketchWindow.end.Format(TimeStampLayout),
	)

	h.sketchWindow.minHeap.Print()
	for i, c := range h.config.WindowTimeSize {
		if h.window[i].queue.IsEmpty() {
			fmt.Printf("subWinSize %d ,in %ds 数据不足为空\n",
				c,
				h.config.SubWindowTimeSize[i])
			continue
		}
		fmt.Printf("subWinSize%d,in %ds,%s~%s,buffer %d top:\n",
			c,
			h.config.SubWindowTimeSize[i],
			h.window[i].queue.Front().(pkg.Synopsis).Start.Format(TimeStampLayout),
			h.window[i].queue.Rear().(pkg.Synopsis).End.Format(TimeStampLayout),
			h.window[i].queue.Size())

		h.window[i].minheap.Print()
	}
	fmt.Print("-----------------------\n")
}
