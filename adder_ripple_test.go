package vgo

import (
	"fmt"
	"testing"
	"time"
)

func TestAdderRipple(t *testing.T) {
	a := make(chan uint64)
	b := make(chan uint64)
	q := make(chan uint64)
	cout := make(chan uint64)

	NewAdderRipple([]chan uint64{a, b}, q)

	go func() {
		defer close(a)
		defer close(b)
		a <- 5
		time.Sleep(1000 * time.Millisecond)
		a <- 0
		time.Sleep(1000 * time.Millisecond)
		b <- 1
	}()

	for {
		select {
		case q, ok := <-q:
			if ok {
				fmt.Printf("q: %d\n", q)
			} else {
				// アウトプットチャンネルがクローズされたら終了
				return
			}
		case cout, ok := <-cout:
			if ok {
				fmt.Printf("cout: %d\n", cout)
			} else {
				// アウトプットチャンネルがクローズされたら終了
				return
			}
		}
	}
}
