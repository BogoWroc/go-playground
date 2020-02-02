package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "Some text yea\nNext line!"
	scanner := bufio.NewScanner(strings.NewReader(s))

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
