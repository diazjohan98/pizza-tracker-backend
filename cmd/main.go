package main

import (
	"log"
	"log/slog"
	"os"
	"pizza-tracker-go/internal/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := loadConfig()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	dbModel, err := models.InitDB(cfg.DBPath)
	if err != nil {
		slog.Error("Failed to initialized database", "Error", err)
		os.Exit(1)
	}

	slog.Info("Database initialized successfully")

	RegisterCustomValidators()

	h := NewHandler(dbModel)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	if err := loadTemplates(router); err != nil {
		slog.Error("Failed to load templates", "error", err)
		os.Exit(1)
	}

	sessionStore := setupSessionStore(dbModel.DB, []byte(cfg.SessionSecretKey))

	setupRouter(router, h, sessionStore)

	slog.Info("Server starting", "url", "http://localhost:"+cfg.Port)

	router.Run(":" + cfg.Port)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
