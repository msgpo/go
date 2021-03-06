/*

Range Continued - range-continued.go

You can skip or value by assigning to _

If you only want the index, drop the ", value" entirely.

*/

package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d > ", value)
		fmt.Printf("%q\n", value)
	}
}
