package top_k

import (
	"bufio"
	"github.com/bmizerany/assert"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestMyHandler_GetTop(t *testing.T) {
	file, err := os.Open("./data/topk-mock-data2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	handler, err := NewHandlerCustom(func(config *Config) {
		config.TopKeepSize = []int{3, 3, 3}
		config.WindowTimeSize = []int{60, 3 * 60} // 1min 3min
		config.SubWindowTimeSize = []int{10, 60}  // 10s,1min

	})
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, handler != nil)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "--print--") {
			handler.Print()
			continue
		}
		kv := strings.Split(scanner.Text(), ",")
		handler.Consume(strings.Trim(kv[1], " "), parseTime(kv[0]))
	}
	handler.Print()
}

func TestGeneratorDate(t *testing.T) {
	file, err := os.Create("./data/topk-mock-data2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.Sync()
	start := parseTime("2021-03-05 09:00:01")
	w := bufio.NewWriter(file)
	for i := 0; i < 60*60; i++ {
		w.WriteString(start.Format(TimeStampLayout + ",贵州茅台\n"))
		d, _ := time.ParseDuration("1s")
		start = start.Add(d)
	}
	w.Flush()

}
