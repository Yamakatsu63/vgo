package vgo

import "fmt"

type Fulladder struct {
	in  chan []Bitvec64
	out chan []Bitvec64
}

func (fulladder *Fulladder) Assign() {
	a := NewReg64(Bitmask64[1])
	b := NewReg64(Bitmask64[1])
	cin := NewReg64(Bitmask64[1])
	q := NewReg64(Bitmask64[1])
	cout := NewReg64(Bitmask64[1])
	for {
		select {
		case n, ok := <-fulladder.in:
			if ok {
				a.value = n[0].value & a.mask
				b.value = n[1].value & b.mask
				cin.value = n[2].value & a.mask
				q.value = a.value ^ b.value ^ cin.value
				cout.value = (a.value & b.value) | (b.value & cin.value) | (cin.value & a.value)
				fmt.Printf("a: %d ", a.value)
				fmt.Printf("b: %d ", b.value)
				fmt.Printf("cin: %d ", cin.value)
				fmt.Printf("q: %d ", q.value)
				fmt.Printf("cout: %d\n", cout.value)
				// アウトチャンネルに信号を送信
				go func() {
					fulladder.out <- []Bitvec64{*q, *cout}
				}()
			} else {
				// チャンネルがクローズされると終了
				return
			}
		}
	}
}
