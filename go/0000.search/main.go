package main

import "fmt"

func main() {
	ns := []int{1, 2, 3, 3, 3, 3, 3, 5, 6, 8, 9, 13, 14, 17, 21}

	// for idx, v := range ns {
	// 	idxl := bsl(ns, v)
	// 	idxr := bsr(ns, v)
	// 	fmt.Printf("%d: %d: %d == %d\n", idx, v, idxl, idxr)
	// }

	// idx := bsrr(ns, 3)
	// var v int
	// if idx < 0 {
	// 	v = -1
	// } else {
	// 	v = ns[idx]
	// }
	// fmt.Printf("idx: %d => %d\n", idx, v)

	v := bsl(ns, 5)
	fmt.Printf("v: %d\n", v)
}
