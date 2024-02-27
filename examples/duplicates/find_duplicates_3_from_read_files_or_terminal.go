// To test programm add the parameters in arguments:
// examples/duplicates/data/file_with_duplicates_1.txt examples/duplicates/data/file_with_duplicates_2.txt
package duplicates

import (
	"fmt"
	"os"
	"strings"
)

func Duplicates3withFiles() {
	counts := map[string]int{}

	for _, fileName := range os.Args[1:] {
		data, err := os.ReadFile(fileName)

		if err != nil {
			_, err2 := fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			if err2 != nil {
				fmt.Println(err)
			}
		}

		lines := strings.Split(string(data), "\n")

		for _, line := range lines {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
