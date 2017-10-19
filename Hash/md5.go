package main

import (
	"crypto/md5"
	"fmt"
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
	hashes := map[string][16]byte{
		f1.plaintext: md5.Sum([]byte(f1.plaintext)),
		f2.plaintext: md5.Sum([]byte(f2.plaintext)),
	}
	for k, v := range hashes {
		fmt.Printf("\nPlaintext:\t%s\n", k)
		fmt.Printf("MD5-Hash:\t%x\n\n", v)

	}
    fmt.Println("The size of an MD5 checksum in bytes: ", md5.Size)
    fmt.Println("The blocksize of MD5 in bytes: ", md5.BlockSize)
}
