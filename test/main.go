package main

import (
	"fmt"
	"unsafe"
)

type Database struct {
	Host     string
	User     string
	Password string
	Port     int
}

func main() {
	d := Database{
		Host:     "",
		User:     "",
		Password: "",
		Port:     0,
	}
	fmt.Println(unsafe.Sizeof(d))
}
