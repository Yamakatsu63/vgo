package vgo

import "fmt"

type CntUnit struct {
	ck  *Bitvec64
	res *Bitvec64
	en  *Bitvec64
	q   *Bitvec64
	ca  *Bitvec64
}

func NewCntUnit(in []chan uint64, out []chan uint64) *CntUnit {
	cntUnit := &CntUnit{NewReg64(Bitmask64[1]), NewReg64(Bitmask64[1]), NewReg64(Bitmask64[1]), NewWire64(Bitmask64[1]), NewWire64(Bitmask64[1])}
	go cntUnit.run(in, out)
	return cntUnit
}

func (cntUnit *CntUnit) exec() {
	cntUnit.ca.value = (cntUnit.en.value & cntUnit.q.value) & cntUnit.q.mask
}

func (cntUnit *CntUnit) run(in []chan uint64, out []chan uint64) {
	defer close(out[0])
	defer close(out[1])
	for {
		select {
		case ck, ok := <-in[0]:
			if ok {
				if cntUnit.ck.value < ck {
					fmt.Println("ckの立ち上がり")
					cntUnit.ck.value = ck & cntUnit.ck.mask
					if cntUnit.res.value == 1 {
						cntUnit.q.value = 0
					} else if cntUnit.en.value == 1 {
						cntUnit.q.value = ^cntUnit.q.value & cntUnit.q.mask
					}
					// 値を更新
					cntUnit.exec()
					out[0] <- cntUnit.q.value
					out[1] <- cntUnit.ca.value
				} else {
					cntUnit.ck.value = ck & cntUnit.ck.mask
				}
			} else {
				return
			}
		case res, ok := <-in[1]:
			if ok {
				if cntUnit.res.value < res {
					fmt.Println("resの立ち上がり")
					cntUnit.res.value = res & cntUnit.res.mask
					if cntUnit.res.value == 1 {
						cntUnit.q.value = 0
					} else if cntUnit.en.value == 1 {
						cntUnit.q.value = ^cntUnit.q.value & cntUnit.q.mask
					}
					// 値を更新
					cntUnit.exec()
					out[0] <- cntUnit.q.value
					out[1] <- cntUnit.ca.value
				} else {
					cntUnit.ck.value = res & cntUnit.ck.mask
				}
			} else {
				return
			}
		case en, ok := <-in[2]:
			if ok {
				// 値を更新
				cntUnit.en.value = en & cntUnit.en.mask
				cntUnit.exec()
				out[0] <- cntUnit.q.value
				out[1] <- cntUnit.ca.value
			} else {
				return
			}
		}
	}
}
