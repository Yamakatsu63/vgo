package vgo

import "fmt"

type Adder struct {
	// a, b, q *Bitvec64
	in  chan []Bitvec64
	out chan Bitvec64
}

func (adder *Adder) Assign() {
	a := NewReg64(Bitmask64[3])
	b := NewReg64(Bitmask64[3])
	q := NewReg64(Bitmask64[3])
	for {
		select {
		case n, ok := <-adder.in:
			if ok {
				a.value = n[0].value & a.mask
				b.value = n[1].value & b.mask
				q.value = (a.value + b.value) & q.mask
				fmt.Println(q.value)
				// アウトチャンネルに信号を送信
				go func() {
					adder.out <- *q
				}()
			} else {
				// チャンネルがクローズされると終了
				return
			}
		}
	}
}
