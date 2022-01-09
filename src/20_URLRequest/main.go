package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456hb"

func main() {

	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	//parsing

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=austinUser",
	}

	anotherURL := partOfUrl.String()
	fmt.Println(anotherURL)

}
