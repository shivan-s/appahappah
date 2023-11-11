package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	md "github.com/gomarkdown/markdown"
)

type Layout struct {
	Title  string
	Header template.HTML
	Main   template.HTML
	Footer template.HTML
}

func main() {
	headerContent, _ := os.ReadFile("content/header.md")
	header := md.ToHTML(headerContent, nil, nil)
	bodyContent, _ := os.ReadFile("content/index.md")
	body := md.ToHTML(bodyContent, nil, nil)
	footerContent, _ := os.ReadFile("content/footer.md")
	footer := md.ToHTML(footerContent, nil, nil)

	os.Mkdir("public", 0755)
	f, _ := os.Create("public/index.html")
	t, _ := template.ParseFiles("static/layout.html")
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

	os.Mkdir("public/css", 0755)
	css, _ := os.ReadFile("static/css/style.css")
	os.WriteFile("public/css/style.css", css, 0644)

	os.Mkdir("public/img", 0755)
	imgs, _ := os.ReadDir("static/img")
	for _, img := range imgs {
		if img.IsDir() == false {
			s := strings.Split(img.Name(), ".")
			ext := s[len(s)-1]
			if ext == "jpg" || ext == "jpeg" || ext == "png" {
				data, err := os.ReadFile(fmt.Sprintf("static/img/%v", img.Name()))
				err = os.WriteFile(fmt.Sprintf("public/img/%v", img.Name()), data, 0644)
				if err != nil {
					fmt.Println(err)
				}

			}
		}
	}
}
