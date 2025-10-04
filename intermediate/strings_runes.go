package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func stringsAndRunes() {
	message := "Hello\n\t\rworld" // t-tab, r-carriage return, n-newline
	rawMessage := `Hello\nWorld`  // escape sequences will not work here

	fmt.Println(message, rawMessage)
	fmt.Printf("length of message: %v\n", len(message))
	fmt.Printf("the first character in message is %v", message[0])

	greeting := "hello"
	name := "adam"
	fmt.Println(greeting + name)

	str1 := "apple"
	str2 := "banana"

	// does not compare length of strings rather the ASCII values of each character in both string
	fmt.Println(str1 < str2) // lexical graphical comparison- ASCII of 'a' < ASCII of 'b'

	for i, char := range message {
		fmt.Printf("character at index %d is %c\n", i, char)
		fmt.Printf("%x\n", char) // returns character as hexadecimal values
	}

	fmt.Println(
		"Rune count",
		utf8.RuneCountInString(greeting),
	) // counts the number of utf8 characters in the string

	// rune is int32 -- runes represent characters in golang
	// runes support any utf8 texts from any language
	ch := 'a'
	fmt.Println(ch)

	runeStr := string(ch) // converts runes character into string
	fmt.Println(runeStr)
	fmt.Printf(
		"Type of runeStr is %T\n",
		runeStr,
	) // this formatting verb %T gives us the type of a variable
}
