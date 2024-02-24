// Print args from cmd
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// way 0
	fmt.Println("-------")
	fmt.Println("way 0")
	fmt.Println(os.Args[1:])
	fmt.Println("-------")

	// way 1 - disadvantages: allocations of strings, due to strings are immutable
	fmt.Println("way 1")
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("-------")

	// way 2 - advantage: NO allocations of strings
	fmt.Println("way 2")
	var stringBuilder strings.Builder
	for _, currentString := range os.Args[1:] {
		stringBuilder.WriteString(currentString)
		stringBuilder.WriteString(" ")
	}
	fmt.Println(stringBuilder.String())
	fmt.Println("-------")

	// way 3 - advantage: NO allocations of strings, code style like in documentation
	fmt.Println("way 3")
	var stringBuilder2 strings.Builder
	stringBuilder2.Grow(1000)
	for _, currentString := range os.Args[1:] {
		_, err := fmt.Fprintf(&stringBuilder2, "%s ", currentString)
		if err != nil {
			return
		}
	}
	fmt.Println(stringBuilder2.String())
	fmt.Println("-------")
}
