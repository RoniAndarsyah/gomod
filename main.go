package main

import "fmt"

var i string

func main() {
	// fmt.Println("Hallo World")
	go Gomog()
	go Asal()
	Kamal()
	fmt.Scan(&i)

}
