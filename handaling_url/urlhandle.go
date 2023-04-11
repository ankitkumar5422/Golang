package main


import (
	"fmt"
	"net/url"
)

const myurl string = "https://go.dev/doc/articles/wiki/"


func main(){
	fmt.Println("Welcome to handling Urls in golang")
	fmt.Println(myurl)

	//parsing
	result, _ :=url.Parse(myurl)

	//fmt.Println(result.Scheme)
	//fmt.Println(result.Host)
	//fmt.Println(result.Path)
	//fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Print("The type of query param are: &t\n",qparams)

	fmt.Println(qparams["coursename"])

	for  _, val := range qparams{
		fmt.Println("param is :",val)
	}
	


}