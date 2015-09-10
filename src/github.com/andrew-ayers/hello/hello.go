package main

import (
	"flag"
	"fmt"
	"playground/helloutil"
)

func main() {
	namePtr := flag.String("name", "yo", "a string")

	fmt.Println(namePtr)

	//namePtr

	flag.Parse()

	//&namePtr = &namePtr + 1
	// = "bitch"

	fmt.Printf("%s\n", helloutil.SayHello(*namePtr))
}
