package basis

import (
	"fmt"
	"time"
)

// var firstName string
// var lastName string
// var phoneNumber string
// var age int
// var virthDate time.Time

type user struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func Structs() {
	user1 := user{
		FirstName:   "Trevor",
		LastName:    "Sawler",
		PhoneNumber: "123",
	}
	fmt.Println("User FirstName", user1.FirstName)
	fmt.Println("User  ", user1.LastName)
	fmt.Println("User ", user1.PhoneNumber)
	fmt.Println("User ", user1.Age)
	fmt.Println("User ", user1.BirthDate)
}

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func StructsWithFuncs() {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "May",
	}

	fmt.Println("MyVar is set to", myVar.printFirstName())
	fmt.Println("MyVar2 is set to", myVar2.printFirstName())
}
