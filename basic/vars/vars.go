package main

import "fmt";

var c, python, java bool = true, false, true;

func main(){
	var i int; // по дефолту 0
	fmt.Println(i, c, python, java);

	var c, j = 5, "NO";
	fmt.Println(c, j);

	// если не хочется писать var, то можно исп := с неявным типо. Это доступно ТОЛЬКО внутри ф-ий
	k := 3;
	s := "string";
	fmt.Println(k, s);
}

