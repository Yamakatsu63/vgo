package vgo

import (
	"testing"
)

func TestAdder(t *testing.T) {
	signal1 := make(chan []Bitvec64)
	signal2 := make(chan Bitvec64)
	a := NewReg64(Bitmask64[3])
	b := NewReg64(Bitmask64[3])
	// q := NewWire64(Bitmask64[3])

	adder := Adder{
		// a:   a,
		// b:   b,
		// q:   q,
		in:  signal1,
		out: signal2,
	}

	go func() {
		defer close(signal1)
		a.Set(1)
		b.Set(0)
		//モジュールに信号を流す
		signal1 <- []Bitvec64{*a, *b}
		a.Set(5)
		signal1 <- []Bitvec64{*a, *b}
		b.Set(4)
		signal1 <- []Bitvec64{*a, *b}
		b.Set(7)
		signal1 <- []Bitvec64{*a, *b}
		a.Set(3)
		signal1 <- []Bitvec64{*a, *b}
	}()

	adder.Assign()
}
