package main

import (
	"fmt"
	"github.com/jzelinskie/whirlpool"
)

func main() {
	w := whirlpool.New()
	text := []byte("This is an example.")
	w.Write(text)
	fmt.Printf("Plaintext:\t%s\n", text)
	fmt.Printf("Hash:\t\t%x\n", w.Sum(nil))

}
