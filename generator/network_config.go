package main

type networkConfig struct {
	Addresses   []string
	Gateway4    string
	Nameservers Nameservers
}

type Nameservers struct {
	Addresses []string
	Search    []string
}
