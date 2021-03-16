package generate

import (
	"fmt"
	"os"
	"path"
	"text/template"
)

func Generate(data templating, outputDir string) error {
	template, err := template.New(data.Name()).Parse(data.Template())

	if err != nil {
		return fmt.Errorf("failed to parse template %q. %v", data.Name(), err)
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	f, err := os.Create(path.Join(outputDir, data.Name()))
	if err != nil {
		return err
	}

	return template.Execute(f, data)
}
