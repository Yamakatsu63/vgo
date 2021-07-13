package vgo

import (
	"fmt"
	"testing"
)

func TestFullAdder(t *testing.T) {
	a := make(chan uint64)
	b := make(chan uint64)
	cin := make(chan uint64)
	q := make(chan uint64)
	cout := make(chan uint64)

	NewFullAdder([]chan uint64{a, b, cin}, []chan uint64{q, cout})

	go func() {
		defer close(a)
		defer close(b)
		defer close(cin)
		a <- 1
		b <- 1
		cin <- 1
		a <- 0
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
