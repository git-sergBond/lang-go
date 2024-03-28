package types

import "fmt"

// ExampleStructWithFunc
// polimorfizm: greeting method implemented for personStruct and adminStruct
func ExampleStructWithFunc() {
	person := personStruct{"Sergey", "Bondarenko", "Sergeevich"}
	admin := adminStruct{"SergeyAdmin", "Bondarenko", "Sergeevich", "ПАО Сбербанк Москва"}

	fmt.Println("=== usage in function ===")
	sayHello(person)
	sayHello(admin)

	fmt.Println("===  convert struct to interface ===")
	// way 1 - ptr
	var personIptr personInterface = &person
	personIptr.greeting()
	// way 2 - value
	var adminI role = admin
	adminI.getRole()

	fmt.Println("===  direct call ===")
	person.greeting()
	admin.greeting()

	fmt.Println("===  for with interface ===")
	// way 1 - array literal
	personArray := [...]personInterface{person, admin}
	fmt.Printf("%T", personArray)
	for _, p := range personArray {
		p.greeting()
	}

	// way 2 - slice literal
	personSlice := []personInterface{person, admin}
	fmt.Printf("%T", personSlice)
	for _, p := range personSlice {
		p.greeting()
	}
}

type personInterface interface {
	greeting()
	role // include interface
}

type role interface {
	getRole() string
}

type personStruct struct{ firstName, lastName, patronymic string }

type adminStruct struct{ firstName, lastName, patronymic, branchName string }

func (s personStruct) greeting() {
	fmt.Println("user", s)
}

func (s adminStruct) greeting() {
	fmt.Println("admin: ", s)
}

func (s personStruct) getRole() string {
	fmt.Println("ROLE: user")
	return "user"
}

func (s adminStruct) getRole() string {
	fmt.Println("ROLE: admin")
	return "admin"
}

func sayHello(p personInterface) {
	p.greeting()
	p.getRole()
}
