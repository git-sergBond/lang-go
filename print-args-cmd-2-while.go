// Print args from cmd
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	var i = 1 // var i int = 1
	for i < len(os.Args) {
		s += sep + os.Args[i]
		sep = " "
		i++
	}
	fmt.Println(s)
}
