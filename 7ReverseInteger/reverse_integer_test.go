package reverseinteger

import (
	"fmt"
	"testing"
)

func TestReverseInteger(t *testing.T) {

	res := reverse(123)
	if res != 321 {
		t.Error(fmt.Sprintf("%d", res) + " != 321")
	}

	res = reverse(2147483326)
	if res != 0 {
		t.Error(fmt.Sprintf("%d", res) + " != 0")
	}

	res = reverse(2147483302)
	if res != 2033847412 {
		t.Error(fmt.Sprintf("%d", res) + " != 2033847412")
	}

	res = reverse(-2147483648)
	if res != 0 {
		t.Error(fmt.Sprintf("%d", res) + " != -8463847412")
	}
}
