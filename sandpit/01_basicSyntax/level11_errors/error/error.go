package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	err := errors.New("Some error message")
	fmt.Printf("%v %T\n", err, err)

	err2 := fmt.Errorf("error occurred at: %v", time.Now())
	fmt.Println("An error happened:", err2)

}
