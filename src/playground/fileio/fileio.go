package main

import (
	//"bufio"
	"fmt"
	//"io"
	"io/ioutil"
	//"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// read in the file
	dat, err := ioutil.ReadFile("data_1.txt")

	check(err)

	fmt.Print(string(dat))

	// write out the file (copy)
	err = ioutil.WriteFile("data_2.txt", dat, 0644)
	check(err)
}
