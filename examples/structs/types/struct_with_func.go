package types

import "fmt"

// ExampleStructWithFunc
// polimorfizm: greeting method implemented for personStruct and adminStruct
func ExampleStructWithFunc() {
	person := personStruct{"Sergey", "Bondarenko", "Sergeevich"}
	admin := adminStruct{"SergeyAdmin", "Bondarenko", "Sergeevich", "ПАО Сбербанк Москва"}

	fmt.Println("=== usage in function ===")
	// usage in function
	sayHello(person)
	sayHello(admin)

	fmt.Println("===  convert struct to interface ===")
	// convert struct to interface - way 1
	var personIptr personInterface = &person
	personIptr.greeting()
	// convert struct to interface - way 1
	var adminI personInterface = &admin
	adminI.greeting()

	fmt.Println("===  implicit call interface ===")
	// implicit call interface
	person.greeting()
	admin.greeting()
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
