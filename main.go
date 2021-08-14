package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/bandozia/lolover/src/global"
	"github.com/bandozia/lolover/src/handler"
	"github.com/bandozia/lolover/src/middleware"
)

//go:embed frontend/public/*
var frontendRoot embed.FS

func main() {
	if e := global.LoadConfigs(); e != nil {
		log.Fatal("Error loading configurations", e.Error())
	}

	startServer(frontendRoot)
}

func startServer(staticContent embed.FS) {
	content, _ := fs.Sub(frontendRoot, "frontend/public")

	dirHandler := handler.Handler{
		HandleFunc: handler.GetDir,
	}

	http.HandleFunc("/api/dir", dirHandler.Handle)

	handler.AddMiddleware(middleware.DevCors)
	http.Handle("/", http.FileServer(http.FS(content)))

	fmt.Printf("RUNNING: http://%s:%d\n", global.Addr, global.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", global.Addr, global.Port), nil))
}
