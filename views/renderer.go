package views

import (
	"embed"
	"groupie-tracker/functions"
	"html/template"
	"io"
)

type ViewDataRenderer struct {
	templ *template.Template
}

var (
	//go:embed htmlTemplates/*
	viewTemplate embed.FS
)

var funcMap = template.FuncMap{
	"add":       functions.Addition,
	"join":      functions.SliceToString,
	"sortDates": functions.SortDates,
}

func NewViewDataRenderer() (*ViewDataRenderer, error) {
	// Initialize template
	templ, err := template.New("Artists").Funcs(funcMap).ParseFS(viewTemplate, "htmlTemplates/*.html")
	if err != nil {
		return nil, err
	}
	return &ViewDataRenderer{templ: templ}, nil
}

func (r *ViewDataRenderer) RenderViewData(w io.Writer, v ViewData) error {
	// Render view data
	if err := r.templ.Execute(w, v); err != nil {
		return err
	}
	return nil
}
