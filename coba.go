package main

import (
	"fmt"

	"github.com/aiteung/atapi"
	"github.com/aiteung/musik"
	"github.com/mitchellh/mapstructure"
)

func Kamal() {
	fmt.Println("Selamat datang pak kamal")
}
func Odading() {
	fmt.Println(musik.GetIPaddress())

}

func Cuanki() {
	var data Data
	myData := atapi.Get("https://dog.ceo/api/breeds/image/random")
	// data = message.Data
	mapstructure.Decode(myData, &data)

	fmt.Println(data.Status)

}

type Data struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}
