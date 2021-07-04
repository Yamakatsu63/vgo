package vgo

// type Fulladder struct {
// 	a, b, q Bitvec64
// }

// func (b *Bitvec64) assign() {
// 	var x, y, z uint64
// 	for {
// 		select {
// 		case n, ok := <-b.in[0]:
// 			if ok {
// 				x = n
// 			} else {
// 				// チャンネルがクローズされると終了
// 				return
// 			}
// 		case n, ok := <-b.in[1]:
// 			if ok {
// 				y = n
// 			} else {
// 				// チャンネルがクローズされると終了
// 				return
// 			}
// 		case n, ok := <-b.in[2]:
// 			if ok {
// 				z = n
// 			} else {
// 				// チャンネルがクローズされると終了
// 				return
// 			}
// 		default:
// 			b.value = (x ^ y ^ z) & b.mask
// 		}
// 	}
// }
