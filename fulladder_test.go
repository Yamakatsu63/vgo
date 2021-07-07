package vgo

import (
	"testing"
)

func TestFullAdder(t *testing.T) {
	signal1 := make(chan []Bitvec64)
	signal2 := make(chan []Bitvec64)
	a := NewReg64(Bitmask64[1])
	b := NewReg64(Bitmask64[1])
	cin := NewWire64(Bitmask64[1])

	fulladder := Fulladder{
		in:  signal1,
		out: signal2,
	}

	go func() {
		defer close(signal1)
		a.Set(0)
		b.Set(0)
		cin.Set(0)
		//モジュールに信号を流す
		signal1 <- []Bitvec64{*a, *b, *cin}
		a.Set(1)
		b.Set(0)
		cin.Set(0)
		signal1 <- []Bitvec64{*a, *b, *cin}
		a.Set(1)
		b.Set(0)
		cin.Set(1)
		signal1 <- []Bitvec64{*a, *b, *cin}
		a.Set(1)
		b.Set(1)
		cin.Set(0)
		signal1 <- []Bitvec64{*a, *b, *cin}
		a.Set(1)
		b.Set(1)
		cin.Set(1)
		signal1 <- []Bitvec64{*a, *b, *cin}
	}()

	fulladder.Assign()
}
