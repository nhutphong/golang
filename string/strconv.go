package main

import (
	"fmt"
	"strconv"

	"phong/tricks"
)


func main() {
	fmt.Println()
	tricks.Format("*")
	i, err := strconv.Atoi("-42") // string to int
	if err == nil {
		fmt.Println("string to int:", i)
	}

	s := strconv.Itoa(-42) // int to string
	fmt.Println("int to string: ", s)


}

const NOTE string = `
	// rune use quote '1 char' or '\u0100'
	// string use double quote "words"


	// string to types="bool float64 int unit"
	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseFloat("3.1415", 64)
	i, err := strconv.ParseInt("-42", 10, 64)
	u, err := strconv.ParseUint("42", 10, 64)

	types to string
	s := strconv.FormatBool(true)
	s := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s := strconv.FormatInt(-42, 16)
	s := strconv.FormatUint(42, 16)


	// base int = 2 8 10 16 ; he so numeric
	// append types to []byte
	func AppendBool(dst []byte, b bool) []byte
	func AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte
	func AppendInt(dst []byte, i int64, base int) []byte

	func AppendQuote(dst []byte, s string) []byte
		b := []byte("quote:")
		b = strconv.AppendQuote(b, `"Fran & Freddie's Diner"`)
		fmt.Println(string(b))
		// quote:"\"Fran & Freddie's Diner\""

	func AppendQuoteRune(dst []byte, r rune) []byte
	func AppendQuoteRuneToASCII(dst []byte, r rune) []byte
	func AppendQuoteRuneToGraphic(dst []byte, r rune) []byte
	func AppendQuoteToASCII(dst []byte, s string) []byte
	func AppendQuoteToGraphic(dst []byte, s string) []byte
	func AppendUint(dst []byte, i uint64, base int) []byte


	//
	func Quote(s string) string
	func QuoteToASCII(s string) string
	func QuoteToGraphic(s string) string
	func QuotedPrefix(s string) (string, error)
	func Unquote(s string) (string, error)
	func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)

	// r rune = dung quote ''
	func QuoteRune(r rune) string
	func QuoteRuneToASCII(r rune) string
	func QuoteRuneToGraphic(r rune) string

`