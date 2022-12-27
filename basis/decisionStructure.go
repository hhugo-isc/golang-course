package basis

import "fmt"

func DesicionStruct() {
	isTrue := true

	if isTrue {
		fmt.Println("Is true is", isTrue)
	} else {
		fmt.Println("Is true is", isTrue)
	}

}

func SwitchStruct() {

	animal := "cat"
	switch animal {
	case "cat":
		fmt.Println("Fish is set to animal")
	case "dog":
		fmt.Println("Fish is set to animal")
	case "fish":
		fmt.Println("Fish is set to animal")
	default:
		fmt.Println("Animal is something else")
	}

}
