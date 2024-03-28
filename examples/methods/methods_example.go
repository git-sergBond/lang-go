package methods

import "fmt"

type library []string

func (l library) print() {
	for _, val := range l {
		fmt.Println(val)
	}
}

// Ex1
func MethodOnType() {
	lib := library{"book1", "book2", "book3"}
	lib.print()
}

//===================================

type t1 struct {
	t int
}

func (tx t1) funcT1(t int) {
	tx.t = t
	fmt.Println("non mutable f1:", tx.t)
}
func (tx *t1) funcT2(t int) {
	tx.t = t
	fmt.Println("mutable f2:", tx.t)
}

// Ex2
func MethodMutability() {
	_t1 := t1{t: 0}

	_t1.funcT1(1)
	fmt.Println("non mutable result:", _t1)

	(&_t1).funcT1(2)
	fmt.Println("non mutable result:", _t1)

	_t1.funcT2(3)
	fmt.Println("mutable result:", _t1)

	(&_t1).funcT2(4)
	fmt.Println("mutable result:", _t1)
}
