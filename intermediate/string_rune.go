package main

import "fmt"

func main() {

	message := "Hello World"
	fmt.Println(message)
	message1 := "Hello, \n world"
	fmt.Println(message1)
	fmt.Println(message[0])

	str1 := "Apple"
	str2 := "app"
	fmt.Println(str1 < str2) //true because 65 < 97

	var ch rune = 'a'
	fmt.Println(ch) //97

}
