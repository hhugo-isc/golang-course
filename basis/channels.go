package basis

import (
	"fmt"

	"github.com/lhhernandez-isc/golang-basis/helpers"
)

const numPool = 10

func calculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func Channels() {
	intChan := make(chan int)

	defer close(intChan)

	go calculateValue(intChan)

	num := <-intChan
	fmt.Println(num)
}
