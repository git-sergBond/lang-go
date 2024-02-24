// Package print_args_cmd
// Print args from cmd
package print_args_cmd

import (
	"fmt"
	"os"
)

func PrintGoArgsCmd1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
