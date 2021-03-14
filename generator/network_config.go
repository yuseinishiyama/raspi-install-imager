package main

import (
	"fmt"
	"os"
	"path"
	"text/template"
)

type networkConfig struct {
	IP string
}

func (n networkConfig) generate(dir string) error {
	templatePath := "templates/network-config"
	filename := path.Base(templatePath)

	template, err := template.New(filename).ParseFiles(templatePath)

	if err != nil {
		return fmt.Errorf("failed to parse %s. %v", templatePath, err)
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	f, err := os.Create(path.Join(dir, filename))
	if err != nil {
		return err
	}

	return template.Execute(f, n)
}
