package anonymous

import "fmt"

func ExampleAnonStructWithFunc() {
	anonymousStructWithMethod := struct {
		firstName, lastName, patronymic string
		Greeting                        func()
	}{firstName: "Sergey", lastName: "Bondarenko", patronymic: "Sergeevich"}

	anonymousStructWithMethod.Greeting = func() {
		fmt.Println(anonymousStructWithMethod)
	}

	anonymousStructWithMethod.Greeting()
}
