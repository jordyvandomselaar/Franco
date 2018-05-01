package controllers

import (
	"html/template"
	"io"
	"log"
	"path/filepath"
	"strings"
)

// This will hold base functionality for controllers.
type Base struct{}

// Render a template. We use go's template library; https://golang.org/pkg/html/template/
// Templates should be loaded in top-level-first order, example: "layout", "home".
// Template paths are relative from templates/, you should separate directories
// with fullstops; home.index.
func (b Base) Render(writer io.Writer, data interface{}, templates ...string) {

	t, err := template.ParseFiles(
		parseTemplatePaths(templates...)...,
	)

	if err != nil {
		log.Fatal(err)
	}

	t.Execute(writer, nil)
}

// Parse dot-separated paths to normal paths relative from templates/.
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
