package basis

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func ReadingJson() {
	myJson := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "Black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false
		}
	]`

	var unmarshaled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshaled)
	if err != nil {
		fmt.Println("Error unmarshaling json", err)
	}

	fmt.Printf("Unmarshaled %v\n", unmarshaled)
}

func WritingJson() {
	var mySlice []Person
	var m1 Person
	var m2 Person

	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	m2.HairColor = "black"
	m2.HasDog = false

	mySlice = append(mySlice, m1)
	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", " ")
	if err != nil {
		fmt.Println("Error marshalling", err)
	}

	fmt.Println(string(newJson))

}
