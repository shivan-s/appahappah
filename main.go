package main

import (
	"html/template"
	"log"
	"os"
	fp "path/filepath"
	"strings"

	md "github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

const (
	publicDir  = "public"
	staticDir  = "static"
	contentDir = "content"
)

type Layout struct {
	Title  string
	Header template.HTML
	Main   template.HTML
	Footer template.HTML
}

func mdToHTML(f string) ([]byte, error) {
	content, err := os.ReadFile(fp.Join(contentDir, f))
	extension := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extension)
	html := md.ToHTML(content, p, nil)
	return html, err
}

func main() {
	header, err := mdToHTML("header.md")
	body, err := mdToHTML("index.md")
	footer, err := mdToHTML("footer.md")

	os.Mkdir(publicDir, 0755)
	f, err := os.Create(fp.Join(publicDir, "index.html"))
	t, err := template.ParseFiles(fp.Join(staticDir, "layout.html"))
	t.Execute(
		f,
		Layout{
			Title:  "Appahappah",
			Header: template.HTML(header),
			Main:   template.HTML(body),
			Footer: template.HTML(footer),
		},
	)
	f.Close()

	os.Mkdir(fp.Join(staticDir, "css"), 0755)
	css, err := os.ReadFile(fp.Join(staticDir, "css", "style.css"))
	os.WriteFile(fp.Join(publicDir, "css", "style.css"), css, 0644)

	os.Mkdir(fp.Join(publicDir, "img"), 0755)
	imgs, err := os.ReadDir(fp.Join(staticDir, "img"))
	for _, img := range imgs {
		if img.IsDir() == false {
			s := strings.Split(img.Name(), ".")
			ext := s[len(s)-1]
			if ext == "jpg" || ext == "jpeg" || ext == "png" {
				data, err := os.ReadFile(fp.Join(staticDir, "img", img.Name()))
				err = os.WriteFile(fp.Join(staticDir, "img", img.Name()), data, 0644)
				if err != nil {
					log.Fatal(err)
				}

			}
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}
