package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const (
	ConfigFile = "config.yml"
)

type config struct {
	Hosts []struct {
		Name   string
		IP     string
		Master bool
	}
}

func main() {
	conf := config{}
	bytes, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		log.Fatalf("failed to read %s. %v", ConfigFile, err)
	}

	if err := yaml.Unmarshal(bytes, &conf); err != nil {
		log.Fatalf("failed to unmarshal %s. %v", ConfigFile, err)
	}

	fmt.Printf("%+v", conf.Hosts[0])
}
