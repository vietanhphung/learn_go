package main

import (
	"fmt"
	"unicode/utf8"
)

// string is a slice of bytes, enclosed in " "
func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i]) //use %x for hexadecimal (UTF-8 encoded), %c for char
	}
}

func printChars(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

func printCharsRune(s string) {
	fmt.Printf("Characters: ")
	runes := []rune(s) // convert to a slice of runes
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func charsAndBytePosition(s string) {
	for index, rune := range s { // iterate over rune with just range{}
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func main() {

	//string is a slice of bytes, enclosed in " "
	name := "Hello World"
	fmt.Println(name)
	printBytes(name)

	name = "Señor"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	// return S e Ã ± o r, wrong
	//UTF-8 encoding a code point can occupy more than 1 byte
	//Unicode code point of ñ is U+00F1 and its UTF-8 encoding occupies 2 bytes c3 and b1.
	// use rune, builtin type, alias of int32, represents a Unicode code point in Go
	printCharsRune(name)
	fmt.Printf("\n")
	// byte index of char -> ñ occupies 2 bytes
	charsAndBytePosition(name)

	// string from slice of bytes
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9} // this is UTF-8 Encoded hex bytes, equivalently in decimal value {67, 97, 102, 195, 169}
	str := string(byteSlice)
	fmt.Println(str)

	// string from slice of runes
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str = string(runeSlice)
	fmt.Println(str)

	// length() return number of bytes, but a char can take more than 1 byte

	word1 := "Señor"
	fmt.Printf("String: %s\n", word1)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word1))
	//the RuneCountInString(s string) (n int) function of the utf8 package can be used to find the length of the string.
	// import "unicode/utf8"
	fmt.Printf("Number of bytes: %d\n", len(word1))

	fmt.Printf("\n")
	word2 := "Pets"
	fmt.Printf("String: %s\n", word2)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word2))
	fmt.Printf("Number of bytes: %d\n", len(word2))

	// String comparison -> == ; string concatenation -> +
	// result := fmt.Sprintf("%s %s", string1, string2) -> format specifier input for Sprintf

	// String in Go is immutable; to change content string -> convert to rune then convert back to rune
	// otherwise reassign is okay -> a := "abc" a:= "abc1"
}
