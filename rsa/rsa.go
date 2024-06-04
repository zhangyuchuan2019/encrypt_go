package rsa

import "github.com/encrypt_go/utils"

type RSA[T utils.Int] struct {
	N, phi T
	e, d   T
}

func GetRSA[T utils.Int](n T) *RSA[T] {
	rsa := &RSA[T]{
		N:   n,
		phi: 0,
		e:   0,
		d:   0,
	}
	prims := utils.GetPrime(n)
	p := prims[len(prims)-1]
	q := prims[len(prims)-2]
	rsa.N = p * q
	rsa.phi = (p - 1) * (q - 1)
	rsa.e = prims[len(prims)/2]
	x, _ := utils.GCD(rsa.e, rsa.phi)
	rsa.d = x%rsa.phi + rsa.phi
	return rsa
}

func (r *RSA[T]) Encrypt(m T) T {
	return utils.QuickPow(m, r.e, r.N)
}

func (r *RSA[T]) Decrypt(c T) T {
	return utils.QuickPow(c, r.d, r.N)
}
