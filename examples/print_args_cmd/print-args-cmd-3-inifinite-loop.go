// Package print_args_cmd
// Print args from cmd
package print_args_cmd

import (
	"fmt"
	"os"
)

func printGoArgsCmd3() {
	var s, sep string

	var i = 1 // var i int = 1
	for {
		if !(i < len(os.Args)) {
			break
		}
		s += sep + os.Args[i]
		sep = " "
		i++
	}
	fmt.Println(s)
}
