package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// instantiate scanner
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("\n\nInput something here: ") // prompt for user

	scanner.Scan() // get some text

	input := scanner.Text() // save it for later

	fmt.Printf("\nYou input: %s\n", input) // display it
}
