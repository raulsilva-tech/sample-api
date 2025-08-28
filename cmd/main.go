package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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

	cfg, err := configs.LoadConfig(getRootPath() + "/.env")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(cfg.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?parseTime=true", cfg.DBUser, cfg.DBUserPassword, cfg.DBHost, cfg.DBPort, cfg.DBDatabaseName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := gin.Default()
	createRoutes(db, r, cfg)

	log.Println("Web server running on ", cfg.WebServerPort)

	log.Fatal(r.Run(":" + fmt.Sprint(cfg.WebServerPort)))
}

func createRoutes(db *sql.DB, r *gin.Engine, cfg *configs.Config) {

	ah := handlers.NewAuthMiddlewareHandler(cfg.JWTSecret)

	uR := repository.NewUserRepository(db)
	evR := repository.NewEventRepository(db)
	etR, _ := repository.NewEventTypeRepository(context.Background(), db)

	etH := handlers.NewEventTypeHandler(etR, evR)
	evH := handlers.NewEventHandler(evR)
	uH := handlers.NewUserHandler(uR, evR, etR, cfg.JWTSecret, cfg.JWTExpiresIn)

	etGroup := r.Group("/event_types")
	etGroup.Use(ah.Authenticate())
	etGroup.POST("/", etH.Insert)
	etGroup.DELETE("/:id", etH.Delete)
	etGroup.PUT("/:id", etH.Update)
	etGroup.GET("/:id", etH.GetOne)
	etGroup.GET("/", etH.GetAll)

	uGroup := r.Group("/users")
	uGroup.Use(ah.Authenticate())
	uGroup.POST("/", uH.Insert)
	uGroup.DELETE("/:id", uH.Delete)
	uGroup.PUT("/:id", uH.Update)
	uGroup.GET("/:id", uH.GetOne)
	uGroup.GET("/", uH.GetAll)
	r.POST("/login", uH.Login)

	evGroup := r.Group("/events")
	evGroup.Use(ah.Authenticate())
	evGroup.POST("/", evH.Insert)
	// evGroup.DELETE("/:id", evH.Delete)
	// evGroup.PUT("/:id", evH.Update)
	// evGroup.GET("/:id", evH.GetOne)
	// evGroup.GET("/",evH.GetAll)

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
