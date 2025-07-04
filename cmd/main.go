package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/raulsilva-tech/SampleAPI/configs"
	"github.com/raulsilva-tech/SampleAPI/internal/repository"
	"github.com/raulsilva-tech/SampleAPI/internal/webserver/handlers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	cfg, err := configs.LoadConfig(getRootPath() + "/config.env")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(cfg.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBUserPassword, cfg.DBHost, cfg.DBPort, cfg.DBDatabaseName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := gin.Default()
	createRoutes(db, r)

	log.Println("Web server running on ", cfg.WebServerPort)
	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(cfg.WebServerPort), r))
}

func createRoutes(db *sql.DB, r *gin.Engine) {

	// uR := repository.NewUserRepository(db)
	evR := repository.NewEventRepository(db)
	etR := repository.NewEventTypeRepository(db)

	etH := handlers.NewEventTypeHandler(etR, evR)
	// evH:= handlers.NewEventHandler(evR)
	// uH := handlers.NewUserRepository(uR,evR)

	etGroup := r.Group("/event_types")
	etGroup.POST("/", etH.Insert)
	etGroup.DELETE("/", etH.Delete)

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
