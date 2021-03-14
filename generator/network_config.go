package main

import "fmt"

type networkConfig struct {
	ip string
}

func (n networkConfig) generate(file string) {
	fmt.Printf("%s\n", n.ip)
}
