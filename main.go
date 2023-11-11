package main

import (
	"html/template"
	"os"

	md "github.com/gomarkdown/markdown"
)

type Layout struct {
	Title   string
	Content template.HTML
}

func main() {
	content, _ := os.ReadFile("content/index.md")
	html := md.ToHTML(content, nil, nil)

	os.Mkdir("public", 0755)
	f, _ := os.Create("public/index.html")
	t, _ := template.ParseFiles("static/layout.html")
	t.Execute(f, Layout{Title: "Appahappah", Content: template.HTML(html)})
	f.Close()

	os.Mkdir("public/css", 0755)
	css, _ := os.ReadFile("static/css/style.css")
	os.WriteFile("public/css/style.css", css, 0644)
}
