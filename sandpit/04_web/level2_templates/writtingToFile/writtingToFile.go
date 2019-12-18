package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	text := "Hello"
	t := `
		<html>
			<body>` + text + `</body>
		</html>
	`
	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal("unable to create index.html file", err)
	}
	defer file.Close()

	io.Copy(file, strings.NewReader(t))
}
