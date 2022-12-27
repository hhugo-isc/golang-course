package basis

import "log"

func Pointers() {
	// 2 POINTERS
	var myString string = "Green"
	log.Println("myString is set to", myString)
	changeUsingPointers(&myString)
	log.Println("After func call, myString is set to", myString)

}

func changeUsingPointers(s *string) {
	newValue := "Red"
	*s = newValue
}
