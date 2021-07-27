package vgo

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	a := make(chan uint64)
	b := make(chan uint64)
	q := make(chan uint64)

	NewAdder([]chan uint64{a, b}, q)

	go func() {
		defer close(a)
		defer close(b)
		a <- 1
		b <- 2
		b <- 4
	}()

	for {
		select {
		case q, ok := <-q:
			if ok {
				fmt.Println(q)
			} else {
				// アウトプットチャンネルがクローズされたら終了
				return
			}
		}
	}
}
