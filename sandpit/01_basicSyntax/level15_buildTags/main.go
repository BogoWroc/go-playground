package main

import "fmt"

var features = []string{
	"Free Feature #1",
	"Free Feature #2",
}

// https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags
// go build -tags pro ./...
//go build -tags "pro enterprise" ./...
func main() {
	for _, f := range features {
		fmt.Println(">", f)
	}
}
