package basis

import (
	"fmt"
	"sort"
)

type userMap struct {
	FirstName string
	LastName  string
}

func Maps() {
	myMap := make(map[string]string)
	myMap["mydog"] = "kiper"
	myMap["mydog2"] = "cbum"

	fmt.Println(myMap["mydog"])
	fmt.Println(myMap["mydog2"])

	myMap2 := make(map[string]userMap)
	me := userMap{
		FirstName: "Hugo",
		LastName:  "Hernandez",
	}

	myMap2["me"] = me
	fmt.Println(myMap2["me"])
}

func Slices() {
	var mySlice []int
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	sort.Ints(mySlice)
	fmt.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(numbers)
	fmt.Println(numbers[0:2])
	fmt.Println(numbers[6:7])

	names := []string{"One", "seven", "fish", "cat"}
	fmt.Println(names)
}
