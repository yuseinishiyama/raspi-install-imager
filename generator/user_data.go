package main

import "fmt"

type userData struct {
	user      string
	publicKey string
}

func (u userData) generate(file string) {
	fmt.Printf("%s\n", u.user)
}
