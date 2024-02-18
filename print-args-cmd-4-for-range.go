// Print args from cmd
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i, arg := range os.Args[1:] {
		fmt.Println("--> i=", i, "e=", os.Args[i], "arg=", arg)

		s += sep + arg
		sep = " "
	}
	fmt.Println("[", s, "]")
}
