package methods

import "fmt"

type library []string

func (l library) print() {
	for _, val := range l {
		fmt.Println(val)
	}
}

func MethodOnType() {
	lib := library{"book1", "book2", "book3"}
	lib.print()
}
