package anonymous

import "fmt"

func ExampleStructWithFunc() {
	anonymousStructWithMethod := personStruct{"Sergey", "Bondarenko", "Sergeevich"}
	sayHello(anonymousStructWithMethod)
}

type personInterface interface {
	greeting()
}

type personStruct struct{ firstName, lastName, patronymic string }

func (s personStruct) greeting() {
	fmt.Println(s)
}

//func (s struct{ firstName, lastName, patronymic string }) greetingAnon() { fmt.Println(s) }

func sayHello(p personInterface) {
	p.greeting()
}
