
package main

import "fmt"


func main(){

    // var demoMap map[string]int //cach1
    dict := make(map[string]int) //cach2

    dict["old"] = 27
    value, ok := dict["old"]
    fmt.Println("value: ",value ,"ok: ", ok)   // Lấy giá trị của phần tử có key = "old", kết quả là 42

    delete(dict, "old")
    fmt.Println(dict["old"])   // Phần tử có key = "old" đã bị xóa, kết quả là 0 (zero value của int)


    // Tiếp theo ví dụ trên
    value, ok := dict["old"]
    fmt.Println("Giá trị của phần tử là: ", value)   // value = 0
    fmt.Println("Kiểm tra phần tử có tồn tại hay không: ", ok)   // ok = false

}