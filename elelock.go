package vgo

import "fmt"

type Elelock struct {
	ck       *Bitvec64
	res      *Bitvec64
	close    *Bitvec64
	tenkey   *Bitvec64
	lock     *Bitvec64
	ke1      *Bitvec64
	ke2      *Bitvec64
	key      []*Bitvec64
	match    *Bitvec64
	key_enbl *Bitvec64
}

// 暗証番号
var SECRET_3 uint64 = 0b0101
var SECRET_2 uint64 = 0b1001
var SECRET_1 uint64 = 0b0110
var SECRET_0 uint64 = 0b0011

func NewElelock(in []chan uint64, out chan uint64) *Elelock {
	elelock := &Elelock{NewReg64(Bitmask64[1]),
		NewReg64(Bitmask64[1]),
		NewReg64(Bitmask64[1]),
		NewReg64(Bitmask64[10]),
		NewWire64(Bitmask64[1]),
		NewWire64(Bitmask64[1]),
		NewWire64(Bitmask64[1]),
		[]*Bitvec64{NewWire64(Bitmask64[4]), NewWire64(Bitmask64[4]), NewWire64(Bitmask64[4]), NewWire64(Bitmask64[4])},
		NewWire64(Bitmask64[1]),
		NewWire64(Bitmask64[1])}
	elelock.lock.value = 1
	go elelock.run(in, out)
	return elelock
}

func (elelock *Elelock) exec() {
	// elelock.match.value = (elelock.key[0].value == SECRET_0) && (elelock.key[1].value == SECRET_1) && (elelock.key[2].value == SECRET_2) && (elelock.key[2].value == SECRET_2)
	if (elelock.key[0].value == SECRET_0) && (elelock.key[1].value == SECRET_1) && (elelock.key[2].value == SECRET_2) && (elelock.key[3].value == SECRET_3) {
		elelock.match.value = 1
	} else {
		elelock.match.value = 0
	}
	elelock.key_enbl.value = (^elelock.ke2.value & elelock.ke1.value) & elelock.key_enbl.mask
}

func (elelock *Elelock) run(in []chan uint64, out chan uint64) {
	defer close(out)
	for {
		select {
		case ck, ok := <-in[0]:
			if ok {
				if elelock.ck.value < ck {
					fmt.Println("ckの立ち上がり")
					elelock.ck.value = ck & elelock.ck.mask
					// 値を更新
					if elelock.res.value == 1 {
						elelock.key[3].value = 0b1111
						elelock.key[2].value = 0b1111
						elelock.key[1].value = 0b1111
						elelock.key[0].value = 0b1111
						elelock.ke2.value = 0
						elelock.ke1.value = 0
						elelock.lock.value = 1
					} else if elelock.close.value == 1 {
						elelock.key[3].value = 0b1111
						elelock.key[2].value = 0b1111
						elelock.key[1].value = 0b1111
						elelock.key[0].value = 0b1111
						elelock.ke2.value = elelock.ke1.value
						elelock.ke1.value = elelock.tenkey.Reductionor().value
						elelock.lock.value = 1
					} else if elelock.key_enbl.value == 1 {
						elelock.key[3].value = elelock.key[2].value
						elelock.key[2].value = elelock.key[1].value
						elelock.key[1].value = elelock.key[0].value
						elelock.key[0].value = elelock.keyenc(elelock.tenkey).value
						fmt.Printf("key: %d %d %d %d\n", elelock.key[3].value, elelock.key[2].value, elelock.key[1].value, elelock.key[0].value)
						elelock.ke2.value = elelock.ke1.value
						elelock.ke1.value = elelock.tenkey.Reductionor().value
					} else if elelock.match.value == 1 {
						elelock.ke2.value = elelock.ke1.value
						elelock.ke1.value = elelock.tenkey.Reductionor().value
						elelock.lock.value = 0
					} else {
						elelock.ke2.value = elelock.ke1.value
						elelock.ke1.value = elelock.tenkey.Reductionor().value
					}
					elelock.exec()
					out <- elelock.lock.value
				} else {
					elelock.ck.value = ck & elelock.ck.mask
				}
			} else {
				return
			}
		case res, ok := <-in[1]:
			if ok {
				if elelock.res.value < res {
					fmt.Println("resの立ち上がり")
					elelock.res.value = res & elelock.res.mask
					// 値を更新
					if elelock.res.value == 1 {
						elelock.key[3].value = 0b1111
						elelock.key[2].value = 0b1111
						elelock.key[1].value = 0b1111
						elelock.key[0].value = 0b1111
						elelock.ke2.value = 0
						elelock.ke1.value = 0
						elelock.lock.value = 1
					} else if elelock.close.value == 1 {
						elelock.key[3].value = 0b1111
						elelock.key[2].value = 0b1111
						elelock.key[1].value = 0b1111
						elelock.key[0].value = 0b1111
						elelock.ke2.value = elelock.ke1.value
						elelock.ke1.value = elelock.tenkey.Reductionor().value
						elelock.lock.value = 1
					} else if elelock.key_enbl.value == 1 {
						elelock.key[3].value = elelock.key[2].value
						elelock.key[2].value = elelock.key[1].value
						elelock.key[1].value = elelock.key[0].value
						elelock.key[0].value = elelock.keyenc(elelock.tenkey).value
						fmt.Printf("key: %d %d %d %d\n", elelock.key[3].value, elelock.key[2].value, elelock.key[1].value, elelock.key[0].value)
						elelock.ke2.value = elelock.ke1.value
						elelock.ke1.value = elelock.tenkey.Reductionor().value
					} else if elelock.match.value == 1 {
						elelock.ke2.value = elelock.ke1.value
						elelock.ke1.value = elelock.tenkey.Reductionor().value
						elelock.lock.value = 0
					}
					elelock.exec()
					out <- elelock.lock.value
				} else {
					elelock.ck.value = res & elelock.ck.mask
				}
			} else {
				return
			}
		case close, ok := <-in[2]:
			if ok {
				// 値を更新
				elelock.close.value = close & elelock.close.mask
				elelock.exec()
				out <- elelock.lock.value
			} else {
				return
			}
		case tenkey, ok := <-in[3]:
			if ok {
				// 値を更新
				elelock.tenkey.value = tenkey & elelock.tenkey.mask
				elelock.exec()
				// out <- elelock.lock.value
			} else {
				return
			}
		}
	}
}

func (elelock *Elelock) keyenc(sw *Bitvec64) *Bitvec64 {
	keyenc := NewWire64(Bitmask64[4])
	// fmt.Printf("入力: %d\n", sw.value)
	switch sw.value {
	case 0b0000000001:
		keyenc.value = 0b0001
	case 0b0000000010:
		keyenc.value = 0b0010
	case 0b0000000100:
		keyenc.value = 0b0011
	case 0b0000001000:
		keyenc.value = 0b0100
	case 0b0000010000:
		keyenc.value = 0b0101
	case 0b0000100000:
		keyenc.value = 0b0110
	case 0b0001000000:
		keyenc.value = 0b0111
	case 0b0010000000:
		keyenc.value = 0b1000
	case 0b0100000000:
		keyenc.value = 0b1001
	case 0b1000000000:
		keyenc.value = 0b1010
	}
	return keyenc
}
