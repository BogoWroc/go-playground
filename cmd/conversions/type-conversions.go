package main

import (
	"fmt"
	"strconv"
)

func main() {
	convertIntToStr()

	convertStrToInt()

	booleanConversion()

	// must return true
	fmt.Println(convertStrToBool("true"))
	fmt.Println(convertStrToBool("yes"))
	fmt.Println(convertStrToBool("1"))

	// must return false
	fmt.Println(convertStrToBool("false"))
	fmt.Println(convertStrToBool("no"))
	fmt.Println(convertStrToBool("0"))
}

func convertStrToBool(value string) bool {
	switch value {
	case "true", "yes", "1":
		return true
	case "false", "no", "0":
		return false
	default:
		return false
	}

}

func booleanConversion() {
	yesNo := "no"
	launch := (yesNo == "yes")
	fmt.Println("Ready for launch:", launch)
}

func convertStrToInt() {
	i, err := strconv.Atoi("10")
	if err != nil {
		fmt.Println("Yyy error")
	}
	fmt.Printf("%T\n", i)
}

func convertIntToStr() {
	// Integer to ascii
	var str = strconv.Itoa(10)
	fmt.Printf("%T\n", str)
}
