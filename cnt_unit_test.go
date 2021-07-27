package vgo

import (
	"fmt"
	"testing"
)

func TestCntUnit(t *testing.T) {
	ck := make(chan uint64)
	res := make(chan uint64)
	en := make(chan uint64)
	q := make(chan uint64)
	ca := make(chan uint64)

	NewCntUnit([]chan uint64{ck, res, en}, []chan uint64{q, ca})

	go func() {
		defer close(ck)
		defer close(res)
		defer close(en)
		ck <- 0
		ck <- 1
		en <- 1
		ck <- 0
		ck <- 1
		ck <- 0
		ck <- 1
		ck <- 0
		ck <- 1
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
		case ca, ok := <-ca:
			if ok {
				fmt.Printf("ca: %d\n", ca)
			} else {
				// アウトプットチャンネルがクローズされたら終了
				return
			}
		}
	}
}
