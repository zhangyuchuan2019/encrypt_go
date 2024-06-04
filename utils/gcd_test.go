package utils

import (
	"fmt"
	"testing"
)

func TestGCD(t *testing.T) {
	a := int64(986040)
	b := int64(439)
	x, y := GCD(a, b)
	fmt.Printf("GCD(%d, %d) = %d*%d + %d*%d\n", a, b, x, a, y, b)
}

func TestGetGCD(t *testing.T) {
	a := int64(70)
	b := int64(20)
	gcd := GetGCD(a, b)
	fmt.Printf("GCD(%d, %d) = %d\n", a, b, gcd)
	if gcd != 10 {
		t.Errorf("GCD(%d, %d) != 10", a, b)
	}
}
