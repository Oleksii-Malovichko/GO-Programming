package main

import "fmt"

// классический вид ф-ии
func add(x int, y int) int {
	return x + y;
}

// сокращенный вид
func add1(x, y int) int {
	return x + y;
}

// возвращение неск значений
func swap(x, y string) (string, string) {
	return y, x;
}

/* 
func имя(ПАРАМЕТРЫ=вход) (ВОЗВРАЩАЕМЫЕ_ЗНАЧЕНИЯ=выход) {
} */

// неявное возвращение тех же переменных (выход), такое лучше исп только в коротких ф-ях
func split(sum int) (x, y int) {
	x = sum * 4 / 9;
	y = sum - x;
	return // тут возвращается x и y по умолчанию
}

func main(){
	// normal funcs
	fmt.Println("The result:", add(5, 5));
	fmt.Printf("The result1: %d\n\n", add1(5, 5));

	// more values in return
	a, b := swap("Hello", "world");
	fmt.Println(a, b);

	// неявный возврат переменных
	fmt.Println(split(17));
}