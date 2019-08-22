package base

import (
	"bufio"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (n int, err error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (w *WordCounter) Write(p []byte) (n int, err error) {
	*w = WordCounter(countWord(p))
	return countWord(p), nil
}

func countWord(p []byte) int {
	sum := 0
	words, _, _ := bufio.ScanWords(p, true)
	if words == 0 {
		sum = 0
	}
	if words >= len(p) {
		sum = 1
	} else {
		sum = countWord(p[words:]) + 1
	}
	return sum
}
