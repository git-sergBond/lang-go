// Print args from cmd
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// way 1 - disadvantages: allocations of strings, due to strings are immutable
	fmt.Println(strings.Join(os.Args[1:], " "))

	// way 2 - advantage: NO allocations of strings
	var stringBuilder strings.Builder
	for _, currentString := range os.Args[1:] {
		stringBuilder.WriteString(currentString)
		stringBuilder.WriteString(" ")
	}
	fmt.Println(stringBuilder.String())
}
