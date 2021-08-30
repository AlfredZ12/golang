package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	suma()
}

func suma() {

	fmt.Println(albums)
}

var albums = []alf{
	{1, "Esther", 28, true},
	{1, "Raquel", 19, false},
	{1, "Yessica", 24, false}}

type alf struct {
	id    int    `structtagjson:"id"`
	name  string `structtagjson:"name"`
	age   int    `structtagjson:"age"`
	state bool   `structtagjson:"state"`
}
