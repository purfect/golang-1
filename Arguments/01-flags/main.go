package main

import (
	"flag"
	"fmt"
)

func main() {
	minusO := flag.String("o", "", "an string, -o=\"test test2\"")
	minusC := flag.Int("C", 0, "an int")
	flag.Parse()
	fmt.Println("-o ", *minusO)
	fmt.Println("-C ", *minusC)

    for index, val := range flag.Args() {
        fmt.Println(index, ": ", val)
    }

}
