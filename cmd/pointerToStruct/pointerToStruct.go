package main

import "fmt"

type myStruct struct {
	value int
}

func main() {
	var st *myStruct = &myStruct{10}
	fmt.Println(st)  //it displays information that st is a pointer to struct
	(*st).value = 20 //dereferencing
	//or thanks to compiler we can write the same without brackets
	st.value = 20
	fmt.Println(*st)
}
