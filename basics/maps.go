package main

import "fmt"

func main() {

	//map

	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["Key1"] = 1
	myMap["key2"] = 2
	fmt.Println(myMap)
	fmt.Println(myMap["Key1"]) //1
	fmt.Println(myMap["key1"]) //0

	delete(myMap, "Key1")
	fmt.Println(myMap) //map[key2:2]

	//to clear map

	clear(myMap) //map[]

	myMap["Key1"] = 1
	myMap["Key2"] = 2
	myMap["Key3"] = 3
	myMap["Key4"] = 4

	value, ok := myMap["Key2"]
	fmt.Println(value) //2
	fmt.Println(ok)    //true

	for k, v := range myMap {
		fmt.Println(k, v)
	}
}
