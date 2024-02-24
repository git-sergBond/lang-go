// Package internal Print args from cmd
package print_args_cmd

import (
	"fmt"
	"os"
)

func printGoArgsCmd2() {
	var s, sep string

	var i = 1 // var i int = 1
	for i < len(os.Args) {
		s += sep + os.Args[i]
		sep = " "
		i++
	}
	fmt.Println(s)
}
