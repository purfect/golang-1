package main

import (
	"fmt"
	"golang.org/x/crypto/sha3"
)

type message struct {
	plaintext string
}

func main() {
	f1 := message{
		"Ich bin Root! Ich darf das!",
	}
	f2 := message{
		"Hello World!",
	}
	hashes := map[string][64]byte{
		f1.plaintext: sha3.Sum512([]byte(f1.plaintext)),
		f2.plaintext: sha3.Sum512([]byte(f2.plaintext)),
	}
	for k, v := range hashes {
		fmt.Printf("Plain:\t%s\n", k)
        fmt.Printf("Sha3::\t%x\n", v)

	}
}
