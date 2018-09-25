package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/**
json支持转义的类型：
boolean, string, float64, nil
map[string]T  pointer

通过[]interface{} 和 map[string]interface{}能够表示所有的json反序列化的数据
 */

type Address struct {
	Type string `json:"type"`
	City string `json:"city"`
	Country string `json:"country"`
	IsInternal bool `json:"isInternal"`
	Area float64 `json:"area"`
}

type VCard struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Address []Address `json:"address"`
	Remark string `json:"remark"`
}

var pa = Address{Type: "private", City: "ad", Country: "China", IsInternal: true, Area: 100}
var wa = Address{Type: "private", City: "ad", Country: "China", IsInternal: false, Area: 200}
var vc = VCard{FirstName: "jan", LastName: "ker", Address: []Address{pa, wa}, Remark: "good"}

func encodeJSON()  {
	byteData, err := json.Marshal(vc)
	if err != nil {
		fmt.Println(err)
		return
	}
	ioutil.WriteFile("./person.json", byteData, 0666)
}

func decodeJSON()  {
	content, err := ioutil.ReadFile("./person.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	byteData := []byte(content)
	vcard := new(VCard)
	json.Unmarshal(byteData, vcard)
	fmt.Println(*vcard)
}

func encodeFile()  {
	file, err := os.OpenFile("./person.json", os.O_WRONLY | os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(vc)
	if err != nil {
		fmt.Println(err)
	}
}

func decodeFile()  {
	file, err := os.Open("./person.json")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	decoder := json.NewDecoder(file)
	var val interface{}
	decoder.Decode(&val)
	fmt.Println(val)
}

//func handleJsonType()  {
//	jsonObject := decodeAndReturnJSON()
//	for key, value := range jsonObject {
//		switch value := value.(type) {
//		case string:
//			fmt.Println(key, "is a string", value)
//		case int:
//			fmt.Println(key, "is a int", value)
//		case []interface{}:
//			fmt.Println(key, "is a map", value)
//		default:
//			fmt.Println("unknown type")
//		}
//	}
//}


//type Car struct {
//	Name string `json:"id, omitempty"`
//	Type string `json:"type, omitempty"`
//	Wheel string `json:"wheel, omitempty"`
//}

func main()  {
	// encodeJSON()
	// decodeJSON()
	// encodeFile()
	decodeFile()
}