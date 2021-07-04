package vgo

import "fmt"

type Adder struct {
	a, b, q Bitvec64
	in      []<-chan uint64 // 読み取り専用のチャンネルの配列
	out     chan<- uint64   // 書き込み専用のチャンネル
}

func (b *Bitvec64) Assign(ports []Bitvec64) {
	var x, y uint64
	for {
		select {
		case n, ok := <-b.in[0]:
			if ok {
				x = n
				b.value = (x + y) & b.mask
				fmt.Println(b.value)
			} else {
				// チャンネルがクローズされると終了
				return
			}
		case n, ok := <-b.in[1]:
			if ok {
				y = n
				b.value = (x + y) & b.mask
				fmt.Println(b.value)
			} else {
				// チャンネルがクローズされると終了
				return
			}
			// default:
			// 	b.value = (x + y) & b.mask
		}
	}
}
