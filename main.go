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

func mdToHTML(f string) []byte {
	content, err := os.ReadFile(fp.Join(contentDir, f))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	extension := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extension)
	html := md.ToHTML(content, p, nil)
	return html
}

func createDir(p string) {
	os.Mkdir(p, 0755)
}

func main() {
	header := mdToHTML("header.md")
	body := mdToHTML("index.md")
	footer := mdToHTML("footer.md")

	createDir(publicDir)
	htmlF, err := os.Create(fp.Join(publicDir, "index.html"))
	t, err := template.ParseFiles(fp.Join(staticDir, "layout.html"))
	t.Execute(
		htmlF,
		Layout{
			Title:  "Appahappah",
			Header: template.HTML(header),
			Main:   template.HTML(body),
			Footer: template.HTML(footer),
		},
	)
	defer htmlF.Close()

	createDir(fp.Join(publicDir, "css"))
	css, err := os.ReadFile(fp.Join(staticDir, "css", "style.css"))
	cssF, err := os.Create(fp.Join(publicDir, "css", "style.css"))
	_, err = cssF.Write(css)
	defer cssF.Close()

	createDir(fp.Join(publicDir, "img"))
	imgs, err := os.ReadDir(fp.Join(staticDir, "img"))
	for _, img := range imgs {
		if img.IsDir() == false {
			s := strings.Split(img.Name(), ".")
			ext := s[len(s)-1]
			if ext == "jpg" || ext == "jpeg" || ext == "png" {
				data, err := os.ReadFile(fp.Join(staticDir, "img", img.Name()))
				imgF, err := os.Create(fp.Join(publicDir, "img", img.Name()))
				_, err = imgF.Write(data)
				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				}
				defer imgF.Close()

			}
		}
	}

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
