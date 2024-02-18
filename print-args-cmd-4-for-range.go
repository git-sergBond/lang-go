// Print args from cmd
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	var indent = 1
	for i, arg := range os.Args[indent:] {
		fmt.Println("--> i=", indent+i, "e=", os.Args[indent+i], "arg=", arg)

		s += sep + arg
		sep = " "
	}
	fmt.Println("[", s, "]")
}
