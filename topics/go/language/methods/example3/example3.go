// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare function variables.
package main

import "fmt"

// data is a struct to bind methods to.
type data struct {
	name string
	age  int
}

// displayName provides a pretty print view of the name.
func (d data) displayName() {
	fmt.Println("My Name Is", d.name)
}

// setAge sets the age and displays the value.
func (d *data) setAge(age int) {
	d.age = age
	fmt.Println(d.name, "Is Age", d.age)
}

func main() {

	// Declare a variable of type data.
	d := data{
		name: "Bill",
	}

	fmt.Println("Proper Calls to Methods:")

	// How we actually call methods in Go.
	d.displayName()
	d.setAge(45)

	fmt.Println("\nWhat the Compiler is Doing:")

	// This is what Go is doing underneath.
	// ----- ruben: this is only for us to see how is done internally, we must never use this way to call methods:
	data.displayName(d)
	(*data).setAge(&d, 45)

	// =========================================================================

	fmt.Println("\nCall Value Receiver Methods with Variable:")

	// Declare a function variable for the method bound to the d variable.
	// The function variable will get its own copy of d because the method
	// is using a value receiver.
	f1 := d.displayName

	// Call the method via the variable.
	f1()

	// Change the value of d.
	d.name = "Joan"

	// Call the method via the variable. We don't see the change.
	// ----- ruben
	// Since the function used in the variable f1 is defined as value semantics,
	// then f1 is using its own copy of d, that is why it prints ‘Bill’ and not ‘Joan’.
	// ----- ruben
	f1()

	// =========================================================================

	fmt.Println("\nCall Pointer Receiver Method with Variable:")

	// Declare a function variable for the method bound to the d variable.
	// The function variable will get the address of d because the method
	// is using a pointer receiver.
	f2 := d.setAge

	// Call the method via the variable.
	f2(45)

	// Change the value of d.
	// ----- ruben
	// for f2 it uses pointer receiver, so it does NOT have its own copy,
	// it points to d and updates it to ‘Sammy’:
	// ----- ruben
	d.name = "Sammy"

	// Call the method via the variable. We see the change.
	f2(45)
}
