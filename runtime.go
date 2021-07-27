package vgo

import "fmt"

var Bitmask64 []uint64

func init() {
	var x uint64 = 0xffffffffffffffff
	Bitmask64 = make([]uint64, 65, 65)
	for i := 64; i >= 0; i-- {
		Bitmask64[i] = x
		x >>= 1
	}
}

type Bitvec64 struct {
	value uint64
	mask  uint64
	undef uint64
}

func NewWire64(mask uint64) *Bitvec64 {
	return &Bitvec64{
		value: 0x0,
		mask:  mask,
		undef: 0xffffffffffffffff & mask,
	}
}

func NewReg64(mask uint64) *Bitvec64 {
	return &Bitvec64{
		value: 0x0,
		mask:  mask,
		undef: 0xffffffffffffffff & mask,
	}
}

func (b *Bitvec64) Set(x uint64) {
	b.value = x & b.mask
	// b.out <- b.value
}

func (b *Bitvec64) Get(index int) uint64 {
	fmt.Printf("get: %d\n", uint64(int(b.value)>>index)&1)
	return uint64(int(b.value)>>index) & 1
}

func (b *Bitvec64) Add(x *Bitvec64) *Bitvec64 {
	b.value += x.value & b.mask
	b.value &= b.mask
	b.undef |= x.undef
	return b
}
