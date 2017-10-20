package main

import (
	"crypto/md5"
	"fmt"

	"github.com/jzelinskie/whirlpool"
	"golang.org/x/crypto/sha3"
)

func main() {
	s := "Passwort123"
	fmt.Printf("Plain:\t\t%s\n", s)
	fmt.Printf("MD5:\t\t%x\n", convert2md5(s))
	fmt.Printf("SHA3:\t\t%x\n", convert2sha3(s))
	fmt.Printf("Whirlpool:\t%x\n", convert2whirlpool(s))

}

func convert2md5(plain string) [16]byte {
	md5hash := md5.Sum([]byte(plain))
	return md5hash

}
func convert2sha3(plain string) [64]byte {
	sha3hash := sha3.Sum512([]byte(plain))
	return sha3hash

}
func convert2whirlpool(plain string) []byte {
	w := whirlpool.New()
	whirlpoolhash := []byte(plain)
	w.Write(whirlpoolhash)
	return w.Sum(nil)

}
