package main

/* Базовые типы GO
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // псевдонім для uint8

rune // псевдонім для int32
     // представляє кодову точку Unicode

float32 float64

complex64 complex128
*/

import (
	"fmt"
	"math/cmplx"
)
// переменные можно оглашать также, как и import
var (
	var1 bool = false;
	var2 uint64 = 1 << 64 - 1;
	var3 complex128 = cmplx.Sqrt(-5 + 12i)
)

// Zero Values. Переменные, которым не присваевается значение получают по дефолту "нулевое значение".
// (0 -> int, false -> bool, "" -> string)

func main(){
	fmt.Printf("Type: %T Value: %v\n", var1, var1); // Type: bool Value: false
	fmt.Printf("Type: %T Value: %v\n", var2, var2); // Type: uint64 Value: 18446744073709551615
	fmt.Printf("Type: %T Value: %v\n", var3, var3); // Type: complex128 Value: (2+3i)

	// Zero Values
	var i int;
	var f float64;
	var b bool;
	var s string;
	var s1 string = "hey";
	fmt.Printf("%v %v %v %q %q\n", i, f, b, s, s1);
}