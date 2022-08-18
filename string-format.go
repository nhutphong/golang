
package main

import (
    "fmt"
    "os"
)

func format() {
    fmt.Println()
    fmt.Println("--------------------------------------------------------------------------------")
    fmt.Println()
}

type point struct {
    x, y int
}

func main() {

    name := "falcon"
    number := 122
    float := 455.67

    fmt.Printf("%s\n", name)
    fmt.Printf("%10s\n", name)

    fmt.Println("---------------------")

    fmt.Printf("%d\n", number)
    fmt.Printf("%7d\n", number)
    fmt.Printf("%07d\n", number)

    fmt.Println("---------------------")

    fmt.Printf("%10f\n", float)
    fmt.Printf("%11f\n", float)
    fmt.Printf("%12f\n", float)


    p := point{1, 2}
    fmt.Printf("struct1: %v\n", p)

    fmt.Printf("struct2: %+v\n", p)

    fmt.Printf("struct3: %#v\n", p)

    fmt.Printf("type: %T\n", p)

    fmt.Printf("bool: %t\n", true)

    fmt.Printf("int: %d\n", 123)

    fmt.Printf("bin: %b\n", 14)

    fmt.Printf("char: %c\n", 33)

    fmt.Printf("hex: %x\n", 456)

    fmt.Printf("float1: %f\n", 78.9)

    fmt.Printf("float2: %e\n", 123400000.0)
    fmt.Printf("float3: %E\n", 123400000.0)

    fmt.Printf("str1: %s\n", "\"string\"")

    fmt.Printf("str2: %q\n", "\"string\"")

    fmt.Printf("str3: %x\n", "hex this")

    fmt.Printf("pointer: %p\n", &p)

    fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

    fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

    fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

    fmt.Printf("width4: |%10s|%10s|\n", "right1", "right2")

    fmt.Printf("width5: |%-30s|%-30s|\n", "left1", "left2")

    s := fmt.Sprintf("sprintf: a %s", "string")
    fmt.Println(s)

    fmt.Fprintf(os.Stderr, "io: an %s\n", "error")

    fmt.Printf("%40s\n", "phong")

    format()

    byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
    str := string(byteSlice) // convert and join ve string
    fmt.Println("string(byteSlice) = ", str)

    // byteSlice := []byte{67, 97, 102, 195, 169} //decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
    
    format()

    runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
    str_two := string(runeSlice)
    fmt.Println("string(runeSlice) = ", str_two)
}