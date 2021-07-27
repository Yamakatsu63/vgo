package vgo

import "fmt"

type Saikoro struct {
	ck     *Bitvec64
	res    *Bitvec64
	enable *Bitvec64
	lamp   *Bitvec64
	cnt    *Bitvec64
}

func NewSaikoro(in []chan uint64, out chan uint64) *Saikoro {
	saikoro := &Saikoro{NewReg64(Bitmask64[1]), NewReg64(Bitmask64[1]), NewReg64(Bitmask64[1]), NewWire64(Bitmask64[7]), NewWire64(Bitmask64[3])}
	go saikoro.run(in, out)
	return saikoro
}

func (saikoro *Saikoro) exec() {
	saikoro.lamp.value = saikoro.dec().value & saikoro.lamp.mask
}

func (saikoro *Saikoro) run(in []chan uint64, out chan uint64) {
	defer close(out)
	for {
		select {
		case ck, ok := <-in[0]:
			if ok {
				if saikoro.ck.value < ck {
					fmt.Println("ckの立ち上がり")
					saikoro.ck.value = ck & saikoro.ck.mask
					// 値を更新
					if saikoro.res.value == 1 {
						saikoro.cnt.value = 1
					} else if saikoro.enable.value == 1 {
						if saikoro.cnt.value == 6 {
							saikoro.cnt.value = 1
						} else {
							saikoro.cnt.value++
						}
					}
					saikoro.exec()
					out <- saikoro.lamp.value
				} else {
					saikoro.ck.value = ck & saikoro.ck.mask
				}
			} else {
				return
			}
		case res, ok := <-in[1]:
			if ok {
				if saikoro.res.value < res {
					fmt.Println("resの立ち上がり")
					saikoro.res.value = res & saikoro.res.mask
					// 値を更新
					if saikoro.res.value == 1 {
						saikoro.cnt.value = 1
					} else if saikoro.enable.value == 1 {
						if saikoro.cnt.value == 6 {
							saikoro.cnt.value = 1
						} else {
							saikoro.cnt.value++
						}
					}
					saikoro.exec()
					out <- saikoro.lamp.value
				} else {
					saikoro.ck.value = res & saikoro.ck.mask
				}
			} else {
				return
			}
		case enable, ok := <-in[2]:
			if ok {
				// 値を更新
				saikoro.enable.value = enable & saikoro.enable.mask
				saikoro.exec()
				out <- saikoro.lamp.value
			} else {
				return
			}
		}
	}
}

func (saikoro *Saikoro) dec() *Bitvec64 {
	dec := NewWire64(Bitmask64[7])
	switch saikoro.cnt.value {
	case 0b001:
		dec.value = 0b0001000
	case 0b010:
		dec.value = 0b1000001
	case 0b011:
		dec.value = 0b0011100
	case 0b100:
		dec.value = 0b1010101
	case 0b101:
		dec.value = 0b1011101
	case 0b110:
		dec.value = 0b1110111
	}
	return dec
}
