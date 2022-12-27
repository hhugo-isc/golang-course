package main

import (
	"fmt"

	"github.com/lhhernandez-isc/golang-basis/helpers"
)

func main() {
	// basis.Variables()
	// basis.Pointers()
	// basis.Structs()
	// basis.StructsWithFuncs()
	// basis.Maps()
	// basis.Slices()
	// basis.DesicionStruct()
	// basis.SwitchStruct()
	// basis.ForRangeStructure()
	// basis.ForRangeStructureInMap()
	// basis.ForRangeOverString()
	// basis.ForRangeOverStruct()
	// basis.Interfaces()
	// basis.Channels()
	// basis.ReadingJson()
	// basis.WritingJson()
	result, err := helpers.Divide(100.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Result of division is", result)
}
