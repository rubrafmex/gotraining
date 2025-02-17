// Ruben: These changes were made in the video of safaribooks but not in this code, so I added it.

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how the concrete value assigned to
// the interface is what is stored inside the interface.
package main

import "fmt"

// printer displays information.
type printer interface {
	print()
}

type user struct {
	name string
}

// print displays the user's name.
func (u user) print() {
	fmt.Printf("User Name: %s\n", u.name)
}

func main() {

	u := user{"PIXMA TR4520"}

	// Add the values and pointers to slice of
	// printer interface values.
	printers := []printer{
		// Store a copy of the user value in the interface value.
		u,
		// Store a copy of the address of the user value in the interface value.
		&u,
	}

	// Change the name field on the user value.
	u.name = "PROGRAF PRO-1000"

	// Iterate over the slice and call
	// print against the copied interface value.
	for _, p := range printers {
		p.print()
	}

	// When we store a value, the interface value has its own
	// copy of the value. Changes to the original value will
	// not be seen.

	// When we store a pointer, the interface value has its own
	// copy of the address. Changes to the original value will
	// be seen.
}
