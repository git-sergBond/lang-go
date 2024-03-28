package methods

import "fmt"

type library []string

func (l library) print() {
	for _, val := range l {
		fmt.Println(val)
	}
}

func (l *library) add(book string) {
	*l = append(*l, book)
}

// Ex1
func MethodOnType() {
	lib := library{"book1", "book2", "book3"}

	lib.add("perfect code")
	lib.print()
}

//===================================

type t1 struct {
	t int
}

func (tx t1) immutableSetter(t int) {
	tx.t = t
	fmt.Println("immutableSetter:", tx.t)
}
func (tx *t1) mutableSetter(t int) {
	tx.t = t
	fmt.Println("mutableSetter:", tx.t)
}

func (tx *t1) mutableIncrement1() {
	tx.t = tx.t + 1
	fmt.Println("mutableIncrement1:", tx.t)
}

func (tx *t1) mutableIncrement2() {
	(*tx).t = (*tx).t + 1
	fmt.Println("mutableIncrement2:", tx.t)
}

// Ex2
func MethodMutability() {
	_t1 := t1{t: 0}

	_t1.immutableSetter(1)
	fmt.Println("1 result:", _t1)

	(&_t1).immutableSetter(2)
	fmt.Println("2 result:", _t1)

	_t1.mutableSetter(3)
	fmt.Println("3 result:", _t1)

	(&_t1).mutableSetter(4)
	fmt.Println("4 result:", _t1)

	t1Link := &_t1

	t1Link.mutableIncrement1()
	fmt.Println("5 result:", _t1)

	t1Link.mutableIncrement2()
	fmt.Println("6 result:", _t1)
}
