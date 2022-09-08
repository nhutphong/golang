package main
 
import (
    "fmt"
    "math/rand"
    "time"
)
 
func CalculateValue(c chan int, num int) {
    value := rand.Intn(10)
    fmt.Println("Calculated Random Value:%d - %d", value, num)
    time.Sleep(1000 * time.Millisecond)
    c <- value
    fmt.Println("This executes regardless as the send is now non-blocking")
}
 
func main() {
    fmt.Println("Go Channel Tutorial")
 
    valueChannel := make(chan int, 3)
    defer close(valueChannel)
 
    go CalculateValue(valueChannel,1)
    go CalculateValue(valueChannel,2)
    go CalculateValue(valueChannel,3)
    go CalculateValue(valueChannel,4)
    go CalculateValue(valueChannel,5)
 
    one := <-valueChannel
    fmt.Println("main", one)
    two := <-valueChannel
    fmt.Println("main", two)
    // three := <-valueChannel
    // fmt.Println("main", three)
    // four := <-valueChannel
    // fmt.Println("main", four)
    // five := <-valueChannel
    // fmt.Println("main", five)
 
    time.Sleep(time.Second)
}