package main

// Parsing multiple templates into a single instance
import (
	"embed"
	"fmt"
	"html/template"
	"os"
)

//go:embed files/*
var files embed.FS

func main() {
	// x, err := files.ReadFile("files/about/index.go.html")
	// if err == nil {
	// 	fmt.Printf("%v\n", string(x))
	// }
	fmt.Println("========================")
	// tmpl, _ := template.ParseFS(files, "files/about/index.go.html", "files/layout.go.html")
	tmpl, _ := template.ParseFS(files, "files/*.go.html")
	tmpl.ParseGlob("files/about/*.go.html")
	tmpl.ExecuteTemplate(os.Stdout, "about/index", "World")
	// tmpl.Execute(os.Stdout, "World")
}
