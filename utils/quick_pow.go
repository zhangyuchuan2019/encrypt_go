package utils

import "math/big"

func QuickPow[T Int](x, p, mod T) T {
	if p == 0 {
		return 1
	}
	if p == 1 {
		return x
	}

	res := big.NewInt(1)
	x %= mod
	num := big.NewInt(int64(x))
	for p > 0 {
		if p&1 == 1 {
			res = res.Mul(res, num).Mod(res, big.NewInt(int64(mod)))
		}
		num = num.Mul(num, num).Mod(num, big.NewInt(int64(mod)))
		p >>= 1
	}
	return T(res.Int64())

}
