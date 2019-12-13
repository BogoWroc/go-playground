package main

import "fmt"

func main() {
	fmt.Println("start main")
	m1()
	fmt.Println("start end")
}

func m1() {
	fmt.Println("start m1")
	defer m2()
	m3()
	fmt.Println("end m1")
}

func m2() {
	fmt.Println("deffered m2")
}

func m3() {
	fmt.Println("m3")
}
