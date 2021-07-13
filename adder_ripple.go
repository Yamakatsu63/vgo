package vgo

// import "fmt"

// type AdderRipple struct {
// 	in  chan []Bitvec64
// 	out chan Bitvec64
// }

// func (adderRipple *AdderRipple) Assign() {
// 	a := NewReg64(Bitmask64[3])
// 	b := NewReg64(Bitmask64[3])
// 	q := NewReg64(Bitmask64[3])
// 	// cout := NewReg64(Bitmask64[3])
// 	signal1 := make(chan []Bitvec64)
// 	signal2 := make(chan []Bitvec64)
// 	signal3 := make(chan []Bitvec64)
// 	signal4 := make(chan []Bitvec64)
// 	add0 := Fulladder{
// 		in:  signal1,
// 		out: signal2,
// 	}
// 	add1 := Fulladder{
// 		in:  signal3,
// 		out: signal4,
// 	}
// 	for {
// 		select {
// 		case n, ok := <-adderRipple.in:
// 			if ok {
// 				a.value = n[0].value & a.mask
// 				b.value = n[1].value & b.mask
// 				// 各モジュールに接続

// 				// アウトチャンネルに信号を送信
// 				go func() {
// 					adderRipple.out <- *q
// 				}()
// 			} else {
// 				// チャンネルがクローズされると終了
// 				return
// 			}
// 		case n, ok := <-add0.out:
// 			if ok {
// 				fmt.Println(n)
// 			}
// 		case n, ok := <-add1.out:
// 			if ok {
// 				fmt.Println(n)
// 			}
// 		}
// 	}
// }
