package numbers

import "testing"

func TestAdd(t *testing.T) {
	result := add(2, 2)
	if result != 4 {
		t.Errorf("Addition(2, 2) = %d; want 4", result)
	}
}
