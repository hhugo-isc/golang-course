package helpers

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
}

func TestDivision(t *testing.T) {
	for _, test := range tests {
		got, err := Divide(test.dividend, test.divisor)
		if test.isErr {
			if err == nil {
				t.Error("Expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error but gone one")
			}
		}

		if got != test.expected {
			t.Errorf("Expected %f but got %f", test.expected, got)
		}
	}
}

// func TestDivide(t *testing.T) {
// 	_, err := Divide(10.0, 1.0)
// 	if err != nil {
// 		t.Error("Got an error when we should not have")
// 	}
// }

// func TestBadDivide(t *testing.T) {
// 	_, err := Divide(10.0, 0.0)
// 	if err == nil {
// 		t.Error("Did not get an error when we should have")
// 	}
// }
