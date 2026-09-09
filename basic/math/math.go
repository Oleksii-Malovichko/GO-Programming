package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main(){
	fmt.Println("My favourite num:", rand.Intn(10));
	fmt.Printf("Here is the root %g !!!\n", math.Sqrt(7));
	fmt.Println(math.Pi) // тут мы экспортировали константу Pi из библиотеки math (экспортированные имена должны начинатся с большой буквы)
}

