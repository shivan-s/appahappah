package main

import (
	"fmt"
	"os"

	md "github.com/gomarkdown/markdown"
)

func main() {
	fileIn := "content/index.md"
	content, _ := os.ReadFile(fileIn)
	html := md.ToHTML(content, nil, nil)
	fmt.Printf(string(html))

	fileOut := "public/index.html"
	os.Mkdir("public", 0755)
	os.WriteFile(fileOut, html, 0644)
}
