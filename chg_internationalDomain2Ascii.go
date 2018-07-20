package main

import (
	"fmt"
	"golang_org/x/net/idna"
)

func main() {
	src := "開発"
	ascii, err := idna.ToASCII(src)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s -> %s\n", src, ascii)
}
