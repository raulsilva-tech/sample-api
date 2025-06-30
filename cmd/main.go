package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/raulsilva-tech/SampleAPI/configs"
)

func main() {

	cfg, err := configs.LoadConfig(getRootPath() + "/config.env")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	log.Println("Web server running on ", cfg.WebServerPort)
	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(cfg.WebServerPort), r))
}

func getRootPath() string {
	// Check if running in "go run" mode (temp binary in /tmp or AppData, etc.)
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	// If temp folder, fallback to working dir
	if strings.Contains(exPath, os.TempDir()) {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		return wd
	}

	return exPath
}
