package _switch

import (
	"fmt"
	"time"
)

func Switch1() {
	switch day := time.Now().Weekday(); day { // assign var before semicolon
	case time.Saturday, time.Sunday: // multiple values
		fmt.Println("weekends -", day)
	default:
		fmt.Println("weekdays -", day)
	}
}

func Switch2() {
	x := 0
	y := 1
	switch { // switch without arg
	case x > 1, y > 0: // replacement logical or operator
		fmt.Println("x > 1 || y > 0")
		fallthrough // test default behaviour case operator for C-like languages
	case false:
		fmt.Println("fallthrough test")
	default:
		fmt.Println("x", x, "y", y)
	}
}

type kilometer float64

func Switch3() {
	var data interface{}

	data = kilometer(123.2323)
	fmt.Printf("%T\n", data)

	switch currentType := data.(type) {
	case string:
		fmt.Println("string")
	case float64:
		fmt.Println("float64")
	case kilometer:
		fmt.Println("kilometer")
	default:
		fmt.Printf("%T", currentType)
	}

}
