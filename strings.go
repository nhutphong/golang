package main

import (
	"fmt"
	"strings"
	// "unicode"
)

func main() {
	stringsRepeat := strings.Repeat("*", 80)
	fmt.Println(stringsRepeat)

}

/*

func Clone(s string) string



func Compare(a, b string) int
	-1 if a < b ; 0 if a == b ; 1 if a > b 
	fmt.Println(strings.Compare("abc", "acb")) // -1 ; so sanh theo cap pair index, 00 11 22

func EqualFold(s, t string) bool
	// cũng là so sánh
	fmt.Println(strings.EqualFold("Go", "go")) //true 
	fmt.Println(strings.EqualFold("AB", "ab")) // true because comparison uses simple case-folding
	fmt.Println(strings.EqualFold("ß", "ss"))  // false because comparison does not use full case-folding
}

func Contains(s, substr string) bool
	// chứa substr = chars 
	fmt.Println(strings.Contains("seafood", "")) // true 
	fmt.Println(strings.Contains("", ""))		 // true 

func ContainsAny(s, chars string) bool
	// chỉ cần chứa 1 char la ok
	fmt.Println(strings.ContainsAny("fail", "ui")) // true 
	fmt.Println(strings.ContainsAny("ure", "ui")) // true 
	fmt.Println(strings.ContainsAny("foo", "")) // false  
	fmt.Println(strings.ContainsAny("", ""))	// false 

func ContainsRune(s string, r rune) bool
	fmt.Println(strings.ContainsRune("aardvark", 97)) //true 
	fmt.Println(strings.ContainsRune("timeout", 97)) //false NOT 'a'

func HasPrefix(s, prefix string) bool // == str.startswith(s, prefix)
func HasSuffix(s, suffix string) bool // == str.endswith(s , suffix)



//return int
func Count(s, substr string) int
	fmt.Println(strings.Count("cheese", "e")) //3 
	fmt.Println(strings.Count("five", "")) //5  before & after each rune

// str.partition(), str.rpartition() python 
func Cut(s, sep string) (before, after string, found bool)
	before, after, found := strings.Cut(s, sep)
	Cut("Gopher", "Go") = "", "pher", true
	Cut("Gopher", "ph") = "Go", "er", true
	Cut("Gopher", "er") = "Goph", "", true
	Cut("Gopher", "Badger") = "Gopher", "", false

// left to right
func Index(s, substr string) int
	// -1 if NOT found

func IndexAny(s, chars string) int
	fmt.Println(strings.IndexAny("chicken", "aeiouy")) // 2 vi co 'e'
	fmt.Println(strings.IndexAny("crwth", "aeiouy")) // -1 not found

func IndexByte(s string, c byte) int
	fmt.Println(strings.IndexByte("golang", 'g'))
	fmt.Println(strings.IndexByte("gophers", 'h'))
	fmt.Println(strings.IndexByte("golang", 'x')) // -1 NOT found

func IndexFunc(s string, f func(rune) bool) int
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f))
	fmt.Println(strings.IndexFunc("Hello, world", f))

func IndexRune(s string, r rune) int
	fmt.Println(strings.IndexRune("chicken", 107)) // 4 ; ord('k') == 107
	fmt.Println(strings.IndexRune("chicken", 'd')) // -1

// right to left
func LastIndex(s, substr string) int
func LastIndexAny(s, chars string) int
func LastIndexByte(s string, c byte) int
func LastIndexFunc(s string, f func(rune) bool) int
	fmt.Println(strings.LastIndexFunc("go 123", unicode.IsNumber)) //5
	fmt.Println(strings.LastIndexFunc("123 go", unicode.IsNumber)) //2
	fmt.Println(strings.LastIndexFunc("go", unicode.IsNumber)) //-1 NOT found




//return []string
//đơn giản không cần dùng đến package: regexp
func Fields(s string) []string
	// cắt full spaces 
	slice = strings.Fields("  foo bar  baz   ") //["foo" "bar" "baz"]

// use if muốn cắt nhiều "abcd" thì "stringsss" nếu chứa các chars "cdab" thì cắt hết 
func FieldsFunc(s string, f func(rune) bool) []string
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo1;bar2,baz3...", f)) //["foo" "bar" "baz"]

// cắt theo substr="dung" phải đúng "dung"
sep=','
func Split(s, sep string) []string
	cắt full 

func SplitN(s, sep string, n int) []string
	cat voi n items

func SplitAfter(s, sep string) []string
	cắt full, sau ',' 

func SplitAfterN(s, sep string, n int) []string
	cắt = n items, sau ',' ; muốn arr có 5 item thì n=5



//return string
//
func Join(elems []string, sep string) string

func Map(mapping func(rune) rune, s string) string
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
	// "Gjnf oevyyvt naq gur fyvgul tbcure..."


func Repeat(s string, count int) string
	strings.Repeat("na", 2) // "nana"

func Replace(s, old, new string, n int) string
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2)) // replaces 2 lần 
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) //-1 replacé full 

func ReplaceAll(s, old, new string) string
	fmt.Println(strings.ReplaceAll("oink oink oink", "oin", "moo")) // replaces full 

func ToLower(s string) string
func ToTitle(s string) string
func ToUpper(s string) string

// use with package: "unicode"
func ToLowerSpecial(c unicode.SpecialCase, s string) string
func ToTitleSpecial(c unicode.SpecialCase, s string) string
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))
func ToUpperSpecial(c unicode.SpecialCase, s string) string

func ToValidUTF8(s, replacement string) string


func TrimLeft(s, cutset string) string
func Trim(s, cutset string) string
	// removed đầu và cuối (left right)
	// dùng từng char in cutset để removed
	fmt.Print(strings.Trim("¡¡¡Hello,! Gophers!!!", "!¡")) // "Hello,! Gophers"
func TrimRight(s, cutset string) string
func TrimSpace(s string) string
	// removed /t /n spaces ... ở left right

func TrimLeftFunc(s string, f func(rune) bool) string
	fmt.Print(strings.TrimLeftFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
func TrimFunc(s string, f func(rune) bool) string
	fmt.Print(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
func TrimRightFunc(s string, f func(rune) bool) string

// phải đúng cả cụm prefix mới removed 
func TrimPrefix(s, prefix string) string
func TrimSuffix(s, suffix string) string

*/