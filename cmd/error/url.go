package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	parse, e := url.Parse("https:// a b.com")
	if e != nil {
		fmt.Println(e)
		fmt.Printf("%#v", e)
		if e, ok := e.(*url.Error); ok {
			fmt.Println("Op:", e.Op)
			fmt.Println("URL:", e.URL)
			fmt.Println("Err:", e.Err)
		}
		os.Exit(1)

	} else {
		fmt.Println(parse)
	}
}
