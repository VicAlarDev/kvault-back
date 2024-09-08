package main

import (
	"context"
	"fmt"
	"github.com/VicAlarDev/kvault-back/internal/adapter/logger"
	"github.com/VicAlarDev/kvault-back/internal/adapter/storage/postgres"
	"github.com/VicAlarDev/kvault-back/internal/adapter/storage/postgres/repository"
	"log/slog"
	"os"

	_ "github.com/VicAlarDev/kvault-back/docs"
	paseto "github.com/VicAlarDev/kvault-back/internal/adapter/auth"
	"github.com/VicAlarDev/kvault-back/internal/adapter/config"
	"github.com/VicAlarDev/kvault-back/internal/adapter/handler/http"
	"github.com/VicAlarDev/kvault-back/internal/core/service"
)

// @title						KVault API
// @version					1.0
// @description				Backend API for KVault APP
//
// @contact.name				Victor Alarcon
// @contact.url				https://github.com/VicAlarDev/
//
// @license.name				MIT
//
// @BasePath					/v1
// @schemes					http https
//
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and the access token.
func main() {
	// Load environment variables
	Config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	// Set logger
	logger.Set(Config.App)

	slog.Info("Starting the application", "app", Config.App.Name, "env", Config.App.Env)

	// Init database
	ctx := context.Background()
	db, err := postgres.New(ctx, Config.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	slog.Info("Successfully connected to the database", "db", Config.DB.Connection)

	// Migrate database
	err = db.Migrate()
	if err != nil {
		slog.Error("Error migrating database", "error", err)
		os.Exit(1)
	}

	slog.Info("Successfully migrated the database")

	/* // Init cache service
	cache, err := redis.New(ctx, Config.Redis)
	if err != nil {
		slog.Error("Error initializing cache connection", "error", err)
		os.Exit(1)
	}
	defer cache.Close()

	slog.Info("Successfully connected to the cache server") */

	// Init token service
	token, err := paseto.New(Config.Token)
	if err != nil {
		slog.Error("Error initializing token service", "error", err)
		os.Exit(1)
	}

	// Dependency injection
	// User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := http.NewUserHandler(*userService)

	// Auth
	authService := service.NewAuthService(userRepo, token)
	authHandler := http.NewAuthHandler(authService)

	// Init router
	router, err := http.NewRouter(
		Config.HTTP,
		token,
		*userHandler,
		*authHandler,
	)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", Config.HTTP.URL, Config.HTTP.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
