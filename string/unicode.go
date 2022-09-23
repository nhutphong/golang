package main

import (
	"fmt"
	"unicode"
)

func main() {
	// constant with mixed type runes
	const mixed = "\b5Ὂg̀9! ℃ᾭG"
	for _, c := range mixed {
		fmt.Printf("For %q:\n", c)
		if unicode.IsControl(c) {
			fmt.Println("\tis control rune")
		}
		if unicode.IsDigit(c) {
			fmt.Println("\tis digit rune")
		}
		if unicode.IsGraphic(c) {
			fmt.Println("\tis graphic rune")
		}
		if unicode.IsLetter(c) {
			fmt.Println("\tis letter rune")
		}
		if unicode.IsLower(c) {
			fmt.Println("\tis lower case rune")
		}
		if unicode.IsMark(c) {
			fmt.Println("\tis mark rune")
		}
		if unicode.IsNumber(c) {
			fmt.Println("\tis number rune")
		}
		if unicode.IsPrint(c) {
			fmt.Println("\tis printable rune")
		}
		if !unicode.IsPrint(c) {
			fmt.Println("\tis not printable rune")
		}
		if unicode.IsPunct(c) {
			fmt.Println("\tis punct rune")
		}
		if unicode.IsSpace(c) {
			fmt.Println("\tis space rune")
		}
		if unicode.IsSymbol(c) {
			fmt.Println("\tis symbol rune")
		}
		if unicode.IsTitle(c) {
			fmt.Println("\tis title case rune")
		}
		if unicode.IsUpper(c) {
			fmt.Println("\tis upper case rune")
		}
	}

	// fmt.Printf("%#v", unicode.Mc)
}

const NOTE string = `

func In(r rune, ranges ...*RangeTable) bool
func Is(rangeTab *RangeTable, r rune) bool
func IsControl(r rune) bool


func IsGraphic(r rune) bool
func IsMark(r rune) bool


func IsOneOf(ranges []*RangeTable, r rune) bool
func SimpleFold(r rune) rune

func IsDigit(r rune) bool
func IsLetter(r rune) bool
func IsLower(r rune) bool
func IsNumber(r rune) bool
func IsPrint(r rune) bool
func IsPunct(r rune) bool
func IsSpace(r rune) bool
func IsSymbol(r rune) bool
func IsTitle(r rune) bool
func IsUpper(r rune) bool

func To(_case int, r rune) rune
func ToLower(r rune) rune
func ToTitle(r rune) rune
func ToUpper(r rune) rune


 `