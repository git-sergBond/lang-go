// Package internal Print args from cmd
package print_args_cmd

import (
	"fmt"
	"os"
)

func printGoArgsCmd4() {
	var s, sep = "", ""
	var indent = 1
	for i, arg := range os.Args[indent:] {
		fmt.Println("--> i=", indent+i, "e=", os.Args[indent+i], "arg=", arg)

		s += sep + arg
		sep = " "
	}
	fmt.Println("[", s, "]")
}
