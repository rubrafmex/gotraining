// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

/*
	https://blog.golang.org/strings

	Go source code is always UTF-8.
	A string holds arbitrary bytes.
	A string literal, absent byte-level escapes, always holds valid UTF-8 sequences.
	Those sequences represent Unicode code points, called runes.
	No guarantee is made in Go that characters in strings are normalized.

	----------------------------------------------------------------------------

	Multiple runes can represent different characters:

	The lower case grave-accented letter à is a character, and it's also a code
	point (U+00E0), but it has other representations.

	We can use the "combining" grave accent code point, U+0300, and attach it to
	the lower case letter a, U+0061, to create the same character à.

	In general, a character may be represented by a number of different sequences
	of code points (runes), and therefore different sequences of UTF-8 bytes.
*/

// Sample program to show how strings have a UTF-8 encoded byte array.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// ----- ruben

	// strings in go are UTF-8 based
	// UTF-8 is a 3 layer character set:

	// on top you have characters
	// in the middle you have codepoints (32 bit or 4 byte value)
	// at the bottom you have bytes

	// Explanation:
	// The idea is that a codepoint is anywhere from 1 to 4 bytes
	// and a character is anywhere from 1 to multiple codepoints.

	// ----- ruben

	// Declare a string with both chinese and english characters.
	s := "世界 means world"

	// UTFMax is 4 -- up to 4 bytes per encoded rune.
	var buf [utf8.UTFMax]byte

	// Iterate over the string.
	for i, r := range s {

		// Ruben notes:
		//
		// When ranging over a string we iterate codepoint by codepoint:
		// i is the index of the codepoint
		// r is going to give us back in that position the codepoint that we just iterated over
		// r represents the type rune.
		// In go 'rune' IS NOT A TYPE, IT IS AN ALIAS for int32
		// the same happens to type 'byte' it is also an alias for uint8
		//
		// r represents rune, so our 32 bit or 4 byte value
		//
		// See the following to see what is being printed to console:

		fmt.Println(i, r)

		// In fmt.Printf:
		// %q is Quoted character (for Character (quoted, Unicode))
		// %d is Base 10 (for Integer)
		fmt.Printf("%q, %d\n", r, utf8.RuneLen(r))

		// Capture the number of bytes for this rune.
		rl := utf8.RuneLen(r)

		// Calculate the slice offset for the bytes associated
		// with this rune.
		si := i + rl

		// Copy of rune from the string to our buffer.
		copy(buf[:], s[i:si])

		// Display the details.
		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n\n", i, r, r, buf[:rl])
	}
}
