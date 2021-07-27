package vgo

import (
	"fmt"
	"testing"
)

func TestSaikoro(t *testing.T) {
	ck := make(chan uint64)
	res := make(chan uint64)
	enable := make(chan uint64)
	lamp := make(chan uint64)

	NewSaikoro([]chan uint64{ck, res, enable}, lamp)

	go func() {
		defer close(ck)
		defer close(res)
		defer close(enable)
		ck <- 0
		ck <- 1
		enable <- 1
		ck <- 0
		ck <- 1 //8
		ck <- 0
		ck <- 1 //65
		ck <- 0
		ck <- 1 //28
		ck <- 0
		ck <- 1 //85
		ck <- 0
		ck <- 1 //93
		ck <- 0
		ck <- 1 //119
		ck <- 0
		ck <- 1 //8
		ck <- 0
		ck <- 1  //65
		res <- 1 //8
	}()

	for {
		select {
		case lamp, ok := <-lamp:
			if ok {
				fmt.Printf("lamp: %d\n", lamp)
			} else {
				// アウトプットチャンネルがクローズされたら終了
				return
			}
		}
	}
}
