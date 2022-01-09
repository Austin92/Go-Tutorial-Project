package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to Json Data lesson")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"ReactJS Bootcamp", 199, "LearnCodeOnline.in", "def123", []string{"full-stack", "js"}},
		{"ReactJS Bootcamp", 699, "LearnCodeOnline.in", "ghi123", nil},
	}

	//package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {

	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
        "Price": 299,
        "website": "LearnCodeOnline.in",
        "_": "abc123",
        "tags": ["web-dev","js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	//some cases where we can add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}

}
