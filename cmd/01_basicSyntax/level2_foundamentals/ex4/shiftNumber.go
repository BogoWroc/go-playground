package main

import "fmt"

func main() {
	number := 1
	displayNumberInManyFormats(number)

	s := number << 1
	displayNumberInManyFormats(s)

}

func displayNumberInManyFormats(number int) {
	fmt.Printf("Decimal: %d, Binary: %[1]b, Hex: %#[1]x\n", number)
}
