package main

import "fmt"

func main() {
	lesson2_4_4()
}

func lesson2_4_1() {
	//Напиши программу, которая читает два целых числа (каждое с новой строки) и выводит их сумму.
	var a int
	var b int
	fmt.Scan(&a, &b)
	fmt.Println(a + b)
}

func lesson2_4_2() {
	//Напиши программу, которая читает два целых числа a и b и выводит результат побитовой операции a & b.
	var a int
	var b int
	fmt.Scan(&a, &b)
	fmt.Println(a & b)
}

func lesson2_4_3() {
	//Напиши программу, которая читает два целых числа a и b и выводит результат побитовой операции a ^ b (исключающее ИЛИ).
	var a int
	var b int
	fmt.Scan(&a, &b)
	fmt.Println(a ^ b)
}

func lesson2_4_4() {
	//Напиши программу, которая читает целое число a и целое число n, затем выводит два результата: a << n и a >> n
	var a int
	var n int
	fmt.Scan(&a, &n)
	fmt.Println(a << n)
	fmt.Println(a >> n)
}
