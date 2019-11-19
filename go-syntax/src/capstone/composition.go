package capstone

import "fmt"

type Engine struct {
	Capacity float32
}

func (e Engine) Start() {
	fmt.Println("Start the engine ...")
}

type Car struct {
	Name string
	Engine
}
