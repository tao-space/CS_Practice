package condition

import "testing"

func TestIfMultiSec(t *testing.T) {
	if a := 2 == 1; a {
		t.Log("1==1")
	} else {
		t.Log("1 != 1")
	}
}

func TestSwitchMultiCase(t *testing.T) {
	// type 1
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("type1: Even")
		case 1, 3:
			t.Log("type1: Old")
		default:
			t.Log("type1: it is not 0-3")
		}
	}

	// type 2
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("type2: Even")
		case i%2 != 0:
			t.Log("type2: Old")
		default:
			t.Log("type2: it is not 0-3")
		}
	}
}
