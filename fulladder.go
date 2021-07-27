package vgo

import "fmt"

type FullAdder struct {
	a    *Bitvec64
	b    *Bitvec64
	cin  *Bitvec64
	q    *Bitvec64
	cout *Bitvec64
}

func NewFullAdder(in []chan uint64, out []chan uint64) *FullAdder {
	fullAdder := &FullAdder{NewReg64(Bitmask64[1]), NewReg64(Bitmask64[1]), NewReg64(Bitmask64[1]), NewWire64(Bitmask64[1]), NewWire64(Bitmask64[1])}
	go fullAdder.run(in, out)
	return fullAdder
}

func (fulladder *FullAdder) exec() {
	fulladder.q.value = (fulladder.a.value ^ fulladder.b.value ^ fulladder.cin.value) & fulladder.q.mask
	fulladder.cout.value = (fulladder.a.value & fulladder.b.value) | (fulladder.b.value & fulladder.cin.value) | (fulladder.cin.value & fulladder.a.value)
}

func (fulladder *FullAdder) run(in []chan uint64, out []chan uint64) {
	defer close(out[0])
	defer close(out[1])
	defer fmt.Println("end")
	for {
		select {
		case a, ok := <-in[0]:
			if ok {
				fmt.Printf("a in fulladder: %d\n", a)
				fulladder.a.value = a & fulladder.a.mask
				fulladder.exec()
				out[0] <- fulladder.q.value
				out[1] <- fulladder.cout.value
			} else {
				// チャンネルがクローズされると終了
				return
			}
		case b, ok := <-in[1]:
			if ok {
				fmt.Printf("b in fulladder: %d\n", b)
				fulladder.b.value = b & fulladder.b.mask
				fulladder.exec()
				out[0] <- fulladder.q.value
				out[1] <- fulladder.cout.value
			} else {
				// チャンネルがクローズされると終了
				return
			}
		case cin, ok := <-in[2]:
			if ok {
				fmt.Printf("cin in fulladder: %d\n", cin)
				fulladder.cin.value = cin & fulladder.cin.mask
				fulladder.exec()
				out[0] <- fulladder.q.value
				out[1] <- fulladder.cout.value
			} else {
				// チャンネルがクローズされると終了
				return
			}
		}
	}
}
