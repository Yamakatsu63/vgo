package vgo

type Adder struct {
	a *Bitvec64
	b *Bitvec64
	q *Bitvec64
}

func NewAdder(in []chan uint64, out chan uint64) *Adder {
	adder := &Adder{NewReg64(Bitmask64[3]), NewReg64(Bitmask64[3]), NewWire64(Bitmask64[3])}
	go adder.run(in, out)
	return adder
}

func (adder *Adder) exec() {
	adder.q.value = (adder.a.value + adder.b.value) & adder.q.mask
}

func (adder *Adder) run(in []chan uint64, out chan uint64) {
	defer close(out)
	for {
		select {
		case a, ok := <-in[0]:
			if ok {
				// 値を更新
				adder.a.value = a & adder.a.mask
				adder.exec()
				out <- adder.q.value
			} else {
				return
			}
		case b, ok := <-in[1]:
			if ok {
				// 値を更新
				adder.b.value = b & adder.b.mask
				adder.exec()
				out <- adder.q.value
			} else {
				return
			}
		}
	}
}
