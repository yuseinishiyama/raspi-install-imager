package main

import (
	"fmt"
	"os"
	"path"
	"text/template"
)

func generate(templatePath string, data interface{}, outputDir string) error {
	filename := path.Base(templatePath)

	template, err := template.New(filename).ParseFiles(templatePath)

	if err != nil {
		return fmt.Errorf("failed to parse %s. %v", templatePath, err)
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	f, err := os.Create(path.Join(outputDir, filename))
	if err != nil {
		return err
	}

	return template.Execute(f, data)
}
