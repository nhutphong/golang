package main

import (
	"fmt"
	"log" // use for debug
	"time"
	"encoding/json"

	"phong/tricks"
)


// type Person struct {
//     name string
//     age  int
//     // NOT exported to json , vì first char là lowercase 
// }

// type Person struct {
//     Name string
//     Age  int
//	   // Name Age,  first char là uppercase, nen exported to json ok 
// }

type Person struct {
    Name string `json:"name"`
    Age  int 	`json:"age"`
}


type FruitBasket struct {
    Name    string
    Fruit   []string
    Id      int64  `json:"ref"`
    private string // An unexported field is not encoded.
    Created time.Time
}

func main() {
	p := Person{"Alice", 22}
	jsonData, _ := json.Marshal(p)
	fmt.Println(string(jsonData	))


	tricks.Format("json.Marshal(basket); struct to json")
	basket := FruitBasket{
    Name:    "Standard",
    Fruit:   []string{"Apple", "Banana", "Orange"},
    Id:      999,
    private: "Second-rate",
    Created: time.Now(),
	}

	var jsonDataTwo []byte
	// jsonDataTwo, err := json.Marshal(basket) // basic
	jsonDataTwo, err := json.MarshalIndent(basket, "", "    ")
	if err != nil {
	    log.Println(err)
	}
	fmt.Println(string(jsonDataTwo))


	tricks.Format("json.Marshal(basket); json to struct")
	jsonDataThree := []byte(`
	{
	    "Name": "Standard",
	    "Fruit": [
	        "Apple",
	        "Banana",
	        "Orange"
	    ],
	    "ref": 999,
	    "Created": "2018-04-09T23:00:00Z"
	}`)

	var basketThree FruitBasket
	errThree := json.Unmarshal(jsonDataThree, &basketThree)
	if errThree != nil {
	    log.Println(errThree)
	}
	fmt.Println(basketThree.Name, basketThree.Fruit, basketThree.Id)
	fmt.Println(basketThree.Created)

	tricks.Format("*")

	jsonData4 := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)
	var info interface{}
	json.Unmarshal(jsonData4, &info)

	// get structData if type=map[string]interface{} va err=nil
	// structData, err := info.(map[string]interface{})
	structData := info.(map[string]interface{})

	fmt.Println(structData)


	for k, v := range structData {
	    switch v := v.(type) {
	    case string:
	        fmt.Println(k, v, "(string)")
	    case float64:
	        fmt.Println(k, v, "(float64)")
	    case []interface{}:
	        fmt.Println(k, "(array):")
	        for i, u := range v {
	            fmt.Println("    ", i, u)
	        }
	    default:
	        fmt.Println(k, v, "(unknown)")
	    }
	}

}