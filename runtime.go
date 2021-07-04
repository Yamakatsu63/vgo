package vgo

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
	in    []<-chan uint64 // 読み取り専用のチャンネルの配列
	out   chan<- uint64   // 書き込み専用のチャンネル
}

func NewWire64(mask uint64, in []<-chan uint64) *Bitvec64 {
	return &Bitvec64{
		value: 0x0,
		mask:  mask,
		in:    in,
	}
}

func NewReg64(mask uint64, out chan<- uint64) *Bitvec64 {
	return &Bitvec64{
		value: 0xffffffffffffffff & mask,
		mask:  mask,
		out:   out,
	}
}

func (b *Bitvec64) Set(x uint64) {
	b.value = x & b.mask
	b.out <- b.value
}

func (b *Bitvec64) Add(x *Bitvec64) *Bitvec64 {
	b.value += x.value & b.mask
	b.value &= b.mask
	b.undef |= x.undef
	return b
}
