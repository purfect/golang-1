package main

import (
	"flag"
	"fmt"
)

func main() {
	minusO := flag.String("o", "", "an string, -o=\"test test2\"")
	minusC := flag.Int("c", 0, "an int")
	flag.Parse()
	fmt.Println("-o ", *minusO)
	fmt.Println("-c ", *minusC)

	for index, val := range flag.Args() {
		fmt.Println(index, ": ", val)

	}

}
