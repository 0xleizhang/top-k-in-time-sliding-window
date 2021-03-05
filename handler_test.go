package top_k

import (
	"bufio"
	"github.com/bmizerany/assert"
	"log"
	"os"
	"strings"
	"testing"
)

func TestMyHandler_GetTop(t *testing.T) {
	file, err := os.Open("./data/topk-mock-data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	handler, err := NewHandlerCustom(func(config *Config) {
		config.TopKeepSize = []int{3, 3, 3}
		config.WindowTimeSize = []int{5 * 60, 60 * 60} // 5min 1h
		config.SubWindowTimeSize = []int{60, 5 * 60}   // 1min,5min

	})
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, handler != nil)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "--print--" {
			handler.Print()
			continue
		}
		kv := strings.Split(scanner.Text(), ",")
		handler.Consume(strings.Trim(kv[1]," "), parseTime(kv[0]))
	}
	handler.Print()
}
