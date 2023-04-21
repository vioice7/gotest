package dog

import "testing"

func TestYears(t *testing.T) {

	x := Years(10)
	if x != 70 {
		t.Error("Exepcted 70 got", x)
	}

	x = Years(8)
	if x != 56 {
		t.Error("Exepcted 56 got", x)
	}

	x = Years(7)
	if x != 49 {
		t.Error("Exepcted 49 got", x)
	}

}
