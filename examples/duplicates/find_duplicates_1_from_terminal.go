package duplicates

import (
	"bufio"
	"fmt"
	"os"
)

func Duplicates1() {
	// creates empty with string key and int value
	counts := make(map[string]int)

	// input data from keyboard
	// quit - break input
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		if text == "quit" {
			break
		}
		fmt.Println("debug:", text)
		counts[text]++
		fmt.Println("debug:", counts)
	}

	// print duplicates
	fmt.Println("duplicates:")
	fmt.Println("value\tkey")
	fmt.Println("(cnt)\t(str)")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
