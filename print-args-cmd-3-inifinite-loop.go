// Print args from cmd
package main

import (
	"fmt"
	"os"
)

func main() {
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
