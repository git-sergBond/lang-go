package anonymous

import "fmt"

func ExampleStructInArg() {
	printAnonymousStruct(struct {
		x,
		y int
	}{
		x: 1,
		y: -1,
	})

	printAnonymousStructArray([]struct {
		x,
		y int
	}{
		{x: 0, y: 0},
		{x: 1, y: 1},
	})

	fmt.Println(returnNestedStruct("A", 5, 10))
}

func printAnonymousStruct(point struct{ x, y int }) {
	fmt.Println(point)
}

func printAnonymousStructArray(arrayPoint []struct{ x, y int }) {
	fmt.Println(arrayPoint)
}

func returnNestedStruct(name string, x, y int) struct {
	name  string
	point struct{ x, y int }
} {
	return struct {
		name  string
		point struct{ x, y int }
	}{
		name: name,
		point: struct{ x, y int }{
			x: x,
			y: y,
		},
	}
}
