package main

import "fmt"

type customErr struct {
	msg string
}

func New(msg string) error {
	return &customErr{msg}
}

func (c *customErr) Error() string {
	return c.msg
}

func foo(err error) {
	fmt.Println(err.Error())
}

func main() {
	foo(New("some error"))
}
