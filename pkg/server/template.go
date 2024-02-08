package server

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"io/fs"

	"github.com/labstack/echo/v4"
)

//go:embed templates
var templateFS embed.FS

type EmbeddedTemplater struct {
	templates map[string]*template.Template
}

func (t *EmbeddedTemplater) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates[name].Execute(w, data)
}

var _ echo.Renderer = (*EmbeddedTemplater)(nil)

func NewEmbeddedTemplater() (*EmbeddedTemplater, error) {
	templates := map[string]*template.Template{}

	tmplFiles, err := fs.ReadDir(templateFS, "templates")
	if err != nil {
		return nil, fmt.Errorf("error reading templates directory: %w", err)
	}

	for _, tmpl := range tmplFiles {
		if tmpl.IsDir() {
			continue
		}

		tmplName := tmpl.Name()
		tmplPath := fmt.Sprintf("templates/%s", tmplName)

		parsed, err := template.New(tmplName).ParseFS(templateFS, tmplPath, "templates/layouts/*.gohtml")
		if err != nil {
			return nil, fmt.Errorf("error parsing template %s: %w", tmplName, err)
		}

		templates[tmplName] = parsed
	}

	return &EmbeddedTemplater{
		templates: templates,
	}, nil
}
