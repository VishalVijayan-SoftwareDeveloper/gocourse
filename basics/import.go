package main

import "fmt"
import "net/http"

func main() {

	fmt.Println("Hello Import a standard library")

	resp, err := http.Get("www.google.com")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	fmt.Println("Http status:", resp.Status)
}
