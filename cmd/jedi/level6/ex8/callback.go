package main

import "fmt"

func callback(text string) {
	fmt.Println(text)
}

func sayHello(data string, callback func(text string)) {
	callback(data)
}

func main() {
	sayHello("Yo", callback)
}
