package anonymous

import "fmt"

func ExampleAnonStructWithFunc() {
	anonymousStructWithMethod := struct {
		id                              int
		firstName, lastName, patronymic string
		Greeting                        func(string) int
	}{
		id:         1,
		firstName:  "Sergey",
		lastName:   "Bondarenko",
		patronymic: "Sergeevich",
	}

	anonymousStructWithMethod.Greeting = func(prefix string) int {
		fmt.Println(prefix, anonymousStructWithMethod)
		return anonymousStructWithMethod.id
	}

	fmt.Println(anonymousStructWithMethod.Greeting("Hello: "))
}
