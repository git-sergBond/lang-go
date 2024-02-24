// Package internal Print args from cmd
package print_args_cmd

import (
	"fmt"
	"os"
)

func printGoArgsCmd1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
