package basis

import "fmt"

func Variables() {
	var whatToSay string
	var i int
	whatToSay = "Goodbye, cruel world!"
	fmt.Println(whatToSay)
	i = 7
	fmt.Println("I is set to", i)
	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println(whatWasSaid)
	fmt.Println(theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
