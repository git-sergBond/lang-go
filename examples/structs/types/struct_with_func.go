package types

import "fmt"

// ExampleStructWithFunc
// polimorfizm: greeting method implemented for personStruct and adminStruct
func ExampleStructWithFunc() {
	person := personStruct{"Sergey", "Bondarenko", "Sergeevich"}
	admin := adminStruct{"SergeyAdmin", "Bondarenko", "Sergeevich", "ПАО Сбербанк Москва"}

	// usage in function
	sayHello(person)
	sayHello(admin)

	// convert struct to interface - way 1
	var personIptr personInterface = &person
	personIptr.greeting()
	// convert struct to interface - way 1
	var adminI personInterface = &admin
	adminI.greeting()
}

type personInterface interface {
	greeting()
}

type personStruct struct{ firstName, lastName, patronymic string }

type adminStruct struct{ firstName, lastName, patronymic, branchName string }

func (s personStruct) greeting() {
	fmt.Println("user", s)
}

func (s adminStruct) greeting() {
	fmt.Println("admin: ", s)
}

func sayHello(p personInterface) {
	p.greeting()
}
