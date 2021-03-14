package main

type networkConfig struct {
	Addresses   []string
	Nameservers Nameservers
}

type Nameservers struct {
	Addresses []string
	Search    []string
}
