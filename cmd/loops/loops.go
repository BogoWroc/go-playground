package main

import "fmt"

func main() {
	forLoop()
	whileLoop()
	doWhileLoop()
	labelsInLoop()
	forEachLoop()
}

func forLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

}

func whileLoop() {

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func doWhileLoop() {
	i := 0
	for {
		fmt.Println("do-while loop")
		i--
		if i <= 0 {
			break
		}
	}
}

func labelsInLoop() {
Label:
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, j)
			if i*j > 4 {
				break Label
			}

		}
	}
}

/*
range can be used for looping over collections: arrays, slices, maps, strings, channels
*/
func forEachLoop() {
	s := "Text"
	for k, v := range s {
		fmt.Println(k, string(v))
	}

	c := []string{"a", "b", "c"}
	for _, v := range c {
		fmt.Println(v)

	}
}
