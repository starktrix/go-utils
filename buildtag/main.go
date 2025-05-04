package main

import "fmt"

// go build -tags dev .
//go:build !dev // build this unless its dev
// +build !dev
func main() {
	fmt.Println("Wakey wakey!!!")
	fmt.Println(Foo())
}