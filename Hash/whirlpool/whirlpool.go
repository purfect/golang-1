package main

import (
	"fmt"
	"github.com/jzelinskie/whirlpool"
)

func main() {
	w := whirlpool.New()
	text := []byte("This is an example.")
	w.Write(text)
	fmt.Printf("\nPlaintext:\t%s\n", text)
	fmt.Printf("Hash:\t\t%x\n\n", w.Sum(nil))

}
