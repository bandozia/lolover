package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/bandozia/lolover/src/fileservice"
	"github.com/bandozia/lolover/src/global"
)

//go:embed frontend/public/*
var frontendRoot embed.FS

func main() {
	if e := global.LoadConfigs(); e != nil {
		log.Fatal("Error loading configurations", e.Error())
	}

	//startServer(frontendRoot)
	fileservice.RenderDir()
}

func startServer(staticContent embed.FS) {
	content, _ := fs.Sub(frontendRoot, "frontend/public")

	// configsHandler := handler.Handler{
	// 	HandleFunc: handler.GetConfigs,
	// }
	// handler.AddMiddleware(middleware.DevCors)
	// handler.AddMiddleware(middleware.TestMidle)

	// http.HandleFunc("/api", configsHandler.Handle)

	http.Handle("/", http.FileServer(http.FS(content)))

	fmt.Printf("RUNNING: http://%s:%d\n", global.Addr, global.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", global.Addr, global.Port), nil))
}
