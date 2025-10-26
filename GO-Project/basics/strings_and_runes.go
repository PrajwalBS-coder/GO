
// A Go string is a read-only slice of bytes. 
// The language and the standard library treat strings specially - as containers of text encoded in UTF-8. In other languages, strings are made of “characters”. 
// In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point. 
// This Go blog post is a good introduction to the topic.

package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {

    const s = "สวัสดี"

    fmt.Println("Len:", len(s))

    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
    fmt.Println()

    fmt.Println("Rune count:", utf8.RuneCountInString(s))

    for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }

    fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width

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


// s is a string assigned a literal value representing the word “hello” in the Thai language. 
// Go string literals are UTF-8 encoded text.
// Since strings are equivalent to []byte, this will produce the length of the raw bytes stored within.
// Indexing into a string produces the raw byte values at each index. 
// This loop generates the hex values of all the bytes that constitute the code points in s.
// To count how many runes are in a string, we can use the utf8 package. 
// Note that the run-time of RuneCountInString depends on the size of the string, because it has to decode each UTF-8 rune sequentially. 
// Some Thai characters are represented by UTF-8 code points that can span multiple bytes, so the result of this count may be surprising.
// A range loop handles strings specially and decodes each rune along with its offset in the string.

// We can achieve the same iteration by using the utf8.
// DecodeRuneInString function explicitly.
// This demonstrates passing a rune value to a function.
// Values enclosed in single quotes are rune literals. 
// We can compare a rune value to a rune literal directly.
// By convention, errors are the last return value and have type error, a built-in interface.
// errors.New constructs a basic error value with the given error message.
// A nil value in the error position indicates that there was no error.
// A sentinel error is a predeclared variable that is used to signify a specific error condition.
// We can wrap errors with higher-level errors to add context. 
// The simplest way to do this is with the %w verb in fmt.Errorf. 
// Wrapped errors create a logical chain (A wraps B, which wraps C, etc.) that can be queried with functions like errors.Is and errors.As.
// It’s idiomatic to use an inline error check in the if line.
// errors.Is checks that a given error (or any error in its chain) matches a specific error value. This is especially useful with wrapped or nested errors, allowing you to identify specific error types or sentinel errors in a chain of errors.