package luke

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {

	var l LukeImpl = LukeImpl{Age: 14}

	if l.Age != 14 {
		t.Errorf("luke's age should be 14\n")
	} else {
		fmt.Printf("age works\n")
	}
}

func Test2(t *testing.T) {
	t.Fail()
}
