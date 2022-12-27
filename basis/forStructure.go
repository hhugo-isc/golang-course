package basis

import (
	"fmt"
)

func ForStructure() {
	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	for i, animal := range animals {
		fmt.Println(i, animal)
	}
}

func ForRangeStructure() {
	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	for i, animal := range animals {
		fmt.Println(i, animal)
	}
}

func ForRangeStructureInMap() {
	animals := make(map[string]string)
	animals["dog"] = "kipper"
	animals["cat"] = "perry"

	for animalType, animal := range animals {
		fmt.Println(animalType, animal)
	}
}

func ForRangeOverString() {
	var firstLine = "Once upon a midnight dreary"
	for i, l := range firstLine {
		fmt.Println(i, l)
	}
}

func ForRangeOverStruct() {
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User

	users = append(users, User{"Mary", "Smith", "msmith@example.com", 30})
	users = append(users, User{"Mary", "Smith", "msmith@example.com", 30})
	users = append(users, User{"Mary", "Smith", "msmith@example.com", 30})

	for i, user := range users {
		fmt.Println(i, user)
	}

}
