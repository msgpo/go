package main

import "fmt"

// Reverse function
func reverse(xs []int) []int {
	ys := make([]int, len(xs))
	for k, v := range xs {
		ys[len(xs)-(k+1)] = v
	}
	return ys
}

/*
func stringReverse(sx []string) []string {
	sy := make([]string, len(sx))
	for p, q := range sx {
		sy[len(sx)-(p+1)] = q
	}
	return sy
}
*/

// Another reverse function
func revert(xr []int) []int {
	if len(xr) <= 1 {
		return xr
	}
	yr := []int{xr[len(xr)-1]}
	yr = append(yr, revert(xr[:len(xr)-1])...)
	return yr
}

// Call of the reverse function
func main() {
	zs := []int{1, 2, 3, 4, 5}
	fmt.Println(reverse(zs))

//	sz := []string{a, b, c, d, e, f}
//	fmt.Println(stringReverse(sz))

	zz := []int{1, 2, 3, 4, 5}
	fmt.Println(revert(zz))

}
