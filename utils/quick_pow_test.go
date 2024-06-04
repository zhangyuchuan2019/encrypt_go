package utils

import "testing"

func TestQuickPow(t *testing.T) {
	x := int64(2)
	p := int64(46411)
	mod := int64(9998000099)
	result := QuickPow(x, p, mod)
	if result != 905611805 {
		t.Errorf("QuickPow(%d, %d, %d) != 905611805, result = %v", x, p, mod, result)
	}
}
