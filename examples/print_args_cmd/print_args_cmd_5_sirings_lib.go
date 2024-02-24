// Package internal Print args from cmd
package print_args_cmd

import (
	"fmt"
	"os"
	"strings"
)

func PrintGoArgsCmd5() {
	// advantage: simple
	fmt.Println("---SimpleWay----")
	SimpleWay()

	// disadvantages: allocations of strings, due to strings are immutable
	fmt.Println("---JoinWay----")
	JoinWay()

	// advantage: NO allocations of strings
	fmt.Println("---StringBuilderWay----")
	StringBuilderWay()

	// advantage: NO allocations of strings, code style like in documentation
	fmt.Println("---StringBuilderFprintfWay----")
	StringBuilderFprintfWay()
}

func SimpleWay() {
	fmt.Println(os.Args[1:])
}

func JoinWay() {
	fmt.Println(strings.Join(os.Args[1:], ","))
}

func StringBuilderWay() {
	var stringBuilder strings.Builder
	for _, currentString := range os.Args[1:] {
		stringBuilder.WriteString(currentString)
		stringBuilder.WriteString(",")
	}
	fmt.Println(stringBuilder.String())
}

func StringBuilderFprintfWay() {
	var stringBuilder strings.Builder
	stringBuilder.Grow(32)
	for _, currentString := range os.Args[1:] {
		_, err := fmt.Fprintf(&stringBuilder, "%s,", currentString)
		if err != nil {
			fmt.Println("ERROR: main. strings.Builder + fmt.Fprintf")
			return
		}
	}
	fmt.Println(stringBuilder.String())
}
