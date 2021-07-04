package vgo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAdder(t *testing.T) {
	signal1 := make(chan uint64)
	signal2 := make(chan uint64)
	a := NewReg64(Bitmask64[3], signal1)
	b := NewReg64(Bitmask64[3], signal2)
	q := NewWire64(Bitmask64[3], []<-chan uint64{signal1, signal2})

	go func() {
		defer close(a.out)
		defer close(b.out)
		a.Set(1)
		time.Sleep(1 * time.Nanosecond)
		// x := 1 + 4
		assert.Equal(t, int(q.value), 1)
		a.Set(5)
		time.Sleep(1 * time.Nanosecond)
		assert.Equal(t, int(q.value), 5)
		b.Set(4)
		time.Sleep(1 * time.Nanosecond)
		assert.Equal(t, int(q.value), 1)
		b.Set(7)
		time.Sleep(1 * time.Nanosecond)
		assert.Equal(t, int(q.value), 4)
		a.Set(3)
		time.Sleep(1 * time.Nanosecond)
		assert.Equal(t, int(q.value), 2)
		// fmt.Println(x)
	}()

	q.Assign([]Bitvec64{*a, *b})
}
