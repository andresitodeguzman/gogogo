/*
Mo To Generator in Go!
Andresito de Guzman
*/
package main

// Import
import (
	"os"
	"fmt"
)

// Define Main
func main(){

	if len(os.Args) == 1 {
		fmt.Printf("Empty Input!")
	} else {
		word := os.Args[1]
		construct := word + " mo to!"
		fmt.Printf(construct)
	}

}