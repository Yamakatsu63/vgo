package vgo

import "fmt"

type AdderRipple struct {
	a    *Bitvec64
	b    *Bitvec64
	q    *Bitvec64
	cout *Bitvec64
}

func NewAdderRipple(in []chan uint64, out chan uint64) *AdderRipple {
	adderRipple := &AdderRipple{NewReg64(Bitmask64[4]), NewReg64(Bitmask64[4]), NewReg64(Bitmask64[4]), NewWire64(Bitmask64[4])}
	go adderRipple.run(in, out)
	return adderRipple
}

func (adderRipple *AdderRipple) run(in []chan uint64, out chan uint64) {
	a0 := make(chan uint64)
	b0 := make(chan uint64)
	cin0 := make(chan uint64)
	q0 := make(chan uint64)
	cout0 := make(chan uint64)
	a1 := make(chan uint64)
	b1 := make(chan uint64)
	cin1 := make(chan uint64)
	q1 := make(chan uint64)
	cout1 := make(chan uint64)
	a2 := make(chan uint64)
	b2 := make(chan uint64)
	cin2 := make(chan uint64)
	q2 := make(chan uint64)
	cout2 := make(chan uint64)
	a3 := make(chan uint64)
	b3 := make(chan uint64)
	cin3 := make(chan uint64)
	q3 := make(chan uint64)
	cout3 := make(chan uint64)
	defer close(out)
	defer close(a0)
	defer close(b0)
	defer close(a1)
	defer close(b1)
	defer close(a2)
	defer close(b2)
	defer close(a3)
	defer close(b3)

	NewFullAdder([]chan uint64{a0, b0, cin0}, []chan uint64{q0, cout0})
	NewFullAdder([]chan uint64{a1, b1, cin1}, []chan uint64{q1, cout1})
	NewFullAdder([]chan uint64{a2, b2, cin2}, []chan uint64{q2, cout2})
	NewFullAdder([]chan uint64{a3, b3, cin3}, []chan uint64{q3, cout3})

	for {
		select {
		case a, ok := <-in[0]:
			if ok {
				fmt.Println(a)
				adderRipple.a.value = a & adderRipple.a.mask
				// adderRipple.a.Get(0)
				a0 <- adderRipple.a.Get(0)
				a1 <- adderRipple.a.Get(1)
				a2 <- adderRipple.a.Get(2)
				a3 <- adderRipple.a.Get(3)
			} else {
				// チャンネルがクローズされると終了
				return
			}
		case b, ok := <-in[1]:
			if ok {
				// fmt.Println(b)
				adderRipple.b.value = b & adderRipple.b.mask
				b0 <- adderRipple.b.Get(0)
				b1 <- adderRipple.b.Get(1)
				b2 <- adderRipple.b.Get(2)
				b3 <- adderRipple.b.Get(3)
			} else {
				// チャンネルがクローズされると終了
				return
			}
		case q0, ok := <-q0:
			if ok {
				fmt.Println(q0)
				// adderRipple.qを更新する
				// out <- q0
			} else {
				// チャンネルがクローズされると終了
				return
			}
		case cout0, ok := <-cout0:
			if ok {
				fmt.Println(cout0)
				// cin1 <- cout0
			} else {
				// チャンネルがクローズされると終了
				return
			}
		case q1, ok := <-q1:
			if ok {
				fmt.Println(q1)
				// out <- q1
			} else {
				// チャンネルがクローズされると終了
				return
			}
		case cout1, ok := <-cout1:
			if ok {
				fmt.Println(cout1)
				cin2 <- cout1
			} else {
				// チャンネルがクローズされると終了
				return
			}
		case q2, ok := <-q2:
			if ok {
				fmt.Println(q2)
				// out <- q2
			} else {
				// チャンネルがクローズされると終了
				return
			}
		case cout2, ok := <-cout2:
			if ok {
				fmt.Println(cout2)
				cin3 <- cout2
			} else {
				// チャンネルがクローズされると終了
				return
			}
		case q3, ok := <-q3:
			if ok {
				fmt.Println(q3)
				out <- q3
			} else {
				// チャンネルがクローズされると終了
				return
			}
		case cout3, ok := <-cout3:
			if ok {
				fmt.Println(cout3)

			} else {
				// チャンネルがクローズされると終了
				return
			}
		}
	}
}
