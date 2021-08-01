package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed frontend/public/*
var frontendRoot embed.FS

func main() {
	contentStatic, _ := fs.Sub(frontendRoot, "frontend/public")

	http.Handle("/", http.FileServer(http.FS(contentStatic)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
