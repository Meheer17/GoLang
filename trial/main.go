// TYPES
// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // alias for uint8
// rune // alias for int32
//      // represents a Unicode code point
// float32 float64
// complex64 complex128

package main

import (
	"fmt"
	"go/types"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var i, j int = 1, 2

func a() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}


func main() {
	fmt.Println(split(17))
	a()
	types.Print()
}
