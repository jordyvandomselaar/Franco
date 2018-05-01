package controllers

import (
	"html/template"
	"io"
	"log"
	"path/filepath"
	"strings"
)

type Base struct{}

func (b Base) Render(writer io.Writer, data interface{}, templates ...string) {

	t, err := template.ParseFiles(
		parseTemplatePaths(templates...)...,
	)

	if err != nil {
		log.Fatal(err)
	}

	t.Execute(writer, nil)
}

func parseTemplatePaths(t ...string) []string {
	var pt []string

	for _, item := range t {
		pathPieces := append([]string{
			"app",
			"templates",
		}, strings.Split(item, ".")...)

		path := filepath.Join(pathPieces...) + ".gohtml"

		pt = append(pt, path)
	}

	return pt
}
