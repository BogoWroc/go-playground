package cmdline

import (
	"fmt"
	"os"
	"strings"
)

func PrintArgs() {
	args := os.Args
	fmt.Println("Command -> " + args[0])
	displayPassedArguments(args)
	// new better way
	fmt.Println("Has args -> " + strings.Join(args[1:], " "))
}

func displayPassedArguments(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("Has args -> " + s)
}
