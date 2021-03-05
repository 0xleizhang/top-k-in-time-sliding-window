package pkg

import "time"

type Word struct {
	Key   string
	Count int
}

//子窗口摘要
type Synopsis struct {
	Start time.Time
	End   time.Time
	Top   []Word
}
