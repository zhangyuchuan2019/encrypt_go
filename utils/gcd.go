package utils

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func GCD[T Int](a, b T) (T, T) {
	if b == 0 {
		return 1, 0
	}
	x, y := GCD(b, a%b)
	return y, x - (a/b)*y
}

func GetGCD[T Int](a, b T) T {
	if b == 0 {
		return a
	}
	x, y := GCD(a, b)
	return x*a + y*b
}
