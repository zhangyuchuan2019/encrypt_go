package main

import (
	"fmt"
	"github.com/encrypt_go/rsa"
)

func main() {
	rsa := rsa.GetRSA(int64(1000000000))
	fmt.Println(rsa)

	test := int64(1234567890)

	enc := rsa.Encrypt(test)
	fmt.Println(enc)

	dec := rsa.Decrypt(enc)
	fmt.Println(dec)

}
