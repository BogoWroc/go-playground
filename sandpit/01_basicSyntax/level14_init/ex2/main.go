package main

//By including the blank identifier, the value of the import itself is discarded,
//so that only the side effect of the import comes through.
import (
	"fmt"
	_ "github.com/bogowroc/go-playground/sandpit/01_basicSyntax/level14_init/ex2/sideEffect"
)

//Unlike the main() function that can only be declared once, the init() function
//can be declared multiple times throughout a package. However, multiple init()s can make
//it difficult to know which one has priority over the others.
//In this section, we will show how to maintain control over multiple init() statements.
func main() {
	fmt.Println("main")

}
