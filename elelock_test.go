package vgo

import (
	"fmt"
	"testing"
)

func TestElelock(t *testing.T) {
	ck := make(chan uint64)
	res := make(chan uint64)
	clos := make(chan uint64)
	tenkey := make(chan uint64)
	lock := make(chan uint64)

	NewElelock([]chan uint64{ck, res, clos, tenkey}, lock)

	go func() {
		defer close(ck)
		defer close(res)
		defer close(clos)
		defer close(tenkey)
		ck <- 0
		ck <- 1
		ck <- 0
		tenkey <- 0b0000010000
		ck <- 1
		ck <- 0
		tenkey <- 0b0000010000
		ck <- 1
		ck <- 0
		tenkey <- 0b0000010000
		ck <- 1
		ck <- 0
		tenkey <- 0
		ck <- 1
		ck <- 0
		tenkey <- 0
		ck <- 1
		ck <- 0
		tenkey <- 0
		ck <- 1
		ck <- 0
		tenkey <- 0b0100000000
		ck <- 1
		ck <- 0
		tenkey <- 0b0100000000
		ck <- 1
		ck <- 0
		tenkey <- 0b0100000000
		ck <- 1
		ck <- 0
		tenkey <- 0
		ck <- 1
		ck <- 0
		tenkey <- 0
		ck <- 1
		ck <- 0
		tenkey <- 0
		ck <- 1
		ck <- 0
		tenkey <- 0b0000100000
		ck <- 1
		ck <- 0
		tenkey <- 0b0000100000
		ck <- 1
		ck <- 0
		tenkey <- 0b0000100000
		ck <- 1
		ck <- 0
		tenkey <- 0
		ck <- 1
		ck <- 0
		tenkey <- 0
		ck <- 1
		ck <- 0
		tenkey <- 0b0000000100
		ck <- 1
		ck <- 0
		tenkey <- 0b0000000100
		ck <- 1
		ck <- 0
		tenkey <- 0b0000000100
		ck <- 1
		ck <- 0
		tenkey <- 0
		ck <- 1
		res <- 1
	}()

	for {
		select {
		case lock, ok := <-lock:
			if ok {
				fmt.Printf("lock: %d\n", lock)
			} else {
				// アウトプットチャンネルがクローズされたら終了
				return
			}
		}
	}
}
