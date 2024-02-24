// Print args from cmd
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// advantage: simple
	fmt.Println("---simpleWay----")
	simpleWay()

	// disadvantages: allocations of strings, due to strings are immutable
	fmt.Println("---joinWay----")
	joinWay()

	// advantage: NO allocations of strings
	fmt.Println("---stringBuilderWay----")
	stringBuilderWay()

	// advantage: NO allocations of strings, code style like in documentation
	fmt.Println("---stringBuilderFprintfWay----")
	stringBuilderFprintfWay()
}

func simpleWay() {
	fmt.Println(os.Args[1:])
}

func joinWay() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func stringBuilderWay() {
	var stringBuilder strings.Builder
	for _, currentString := range os.Args[1:] {
		stringBuilder.WriteString(currentString)
		stringBuilder.WriteString(" ")
	}
	fmt.Println(stringBuilder.String())
}

func stringBuilderFprintfWay() {
	var stringBuilder strings.Builder
	stringBuilder.Grow(32)
	for _, currentString := range os.Args[1:] {
		_, err := fmt.Fprintf(&stringBuilder, "%s ", currentString)
		if err != nil {
			fmt.Println("ERROR: main. strings.Builder + fmt.Fprintf")
			return
		}
	}
	fmt.Println(stringBuilder.String())
}
