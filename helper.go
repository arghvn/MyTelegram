package main

// for unicode/utf8

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// s is a string assigned a literal value representing the word “hello” in the Thai language.
	//  Go string literals are UTF-8 encoded text.
	const s = "สวัสดี"

	fmt.Println("Len:", len(s))

	// Since strings are equivalent to []byte, this will produce the length of the raw bytes stored within.Indexing into a string produces
	// the raw byte values at each index. This loop generates the hex values of all the bytes that constitute the code points in s.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	// To count how many runes are in a string, we can use the utf8 package.
	// Note that the run-time of RuneCountInString dependes on the size of the string,
	// because it has to decode each UTF-8 rune sequentially. Some Thai characters are represented by multiple UTF-8 code points.
	// A range loop handles strings specially and decodes each rune along with its offset in the string.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		// *
		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

// output :

// Len: 18
// e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5
// Rune count: 6
// U+0E2A 'ส' starts at 0
// U+0E27 'ว' starts at 3
// U+0E31 'ั' starts at 6
// U+0E2A 'ส' starts at 9
// U+0E14 'ด' starts at 12
// U+0E35 'ี' starts at 15

// Using DecodeRuneInString
// U+0E2A 'ส' starts at 0
// found so sua
// U+0E27 'ว' starts at 3
// U+0E31 'ั' starts at 6
// U+0E2A 'ส' starts at 9
// found so sua
// U+0E14 'ด' starts at 12
// U+0E35 'ี' starts at 15

// *
// We can achieve the same iteration by using the utf8.DecodeRuneInString function explicitly.
// This demonstrates passing a rune value to a function.Values enclosed in single quotes are rune literals.
//  We can compare a rune value to a rune literal directly.
