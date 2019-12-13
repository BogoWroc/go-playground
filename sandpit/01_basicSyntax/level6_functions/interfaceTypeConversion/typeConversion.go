package main

import "fmt"

type Writer interface {
	Write(s string)
}

type Reader interface {
	Read() string
}

type WriteRead interface {
	Writer
	Reader
}

type Syso struct {
	data string
}

func (s *Syso) Write(d string) {
	s.data = d
}

func (s Syso) Read() string {
	return s.data
}

func main() {
	var i WriteRead = &Syso{data: ""}
	i.Write("Test Data")
	r, ok := i.(Reader) // interface converting
	if ok {
		fmt.Println(r.Read())
	}

}
