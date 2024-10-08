package main

import (
	"fmt"
)

func main() {
	// defer fun1()
	fun1()
	fun2()

}

func fun1() {
	fmt.Println("fun 1")

	defer fmt.Println("line 2")

	fmt.Println("line 3")

	defer fmt.Println("line 4")

}

func fun2() {
	fmt.Println("fun 2")

}
