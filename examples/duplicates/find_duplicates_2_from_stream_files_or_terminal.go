// To test programm add the parameters in arguments:
// examples/duplicates/data/file_with_duplicates_1.txt examples/duplicates/data/file_with_duplicates_2.txt
package duplicates

import (
	"bufio"
	"fmt"
	"os"
)

func Duplicates2withFiles() {
	counts := make(map[string]int)
	var filesWithDuplcates []string

	files := os.Args[1:] //read names of files from arguments cmd

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, fileName := range files {
			file, openFileError := os.Open(fileName)
			if openFileError != nil {
				_, err := fmt.Fprintf(os.Stderr, "dup2: %v\n", openFileError)
				if err != nil {
					fmt.Println(err)
				}
				continue
			}

			hasDuplicates := countLines(file, counts)
			if hasDuplicates {
				filesWithDuplcates = append(filesWithDuplcates, fileName)
			}

			if err := file.Close(); err != nil {
				fmt.Println(err)
			}
		}
	}

	if filesWithDuplcates != nil {
		fmt.Println("Files with duplicates:")
	}
	for _, fileName := range filesWithDuplcates {
		fmt.Println(fileName)
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(file *os.File, counts map[string]int) bool {
	input := bufio.NewScanner(file)

	hasDuplicates := false

	for input.Scan() {
		err1 := input.Err()
		if err1 != nil {
			_, err := fmt.Fprintf(os.Stderr, "dup2: %v\n", err1)
			if err != nil {
				fmt.Println(err)
			}
		}

		counts[input.Text()]++
		hasDuplicates = true
	}

	return hasDuplicates
}
