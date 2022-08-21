package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
	"transaction_service/internal/config"
	"transaction_service/internal/middleware"
	"transaction_service/internal/routers"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
)

type envConfig struct {
	DBHost          string `envconfig:"DB_HOST" default:"127.0.0.1"`
	DBName          string `envconfig:"DB_NAME" default:"pos_majoo"`
	DBUsername      string `envconfig:"DB_USER" default:"postgres"`
	DBPassword      string `envconfig:"DB_PWD" default:"root"`
	DBPort          int    `envconfig:"DB_PORT" default:"5432"`
	JWTSignatureKey string `envconfig:"SIGNATURE_KEY" default:"IoluSDS49v"`
}

func main() {
	app := echo.New()
	// init environtment config
	env := envConfig{}
	envconfig.MustProcess("", &env)
	// init db
	db, err := config.ConnectDB(env.DBHost, env.DBName, env.DBPort, env.DBUsername, env.DBPassword)
	if err == nil {
		ctx := context.WithValue(context.Background(),
			"JWT_config",
			map[string]string{
				"Signature_Key": env.JWTSignatureKey,
			})

		JWTMiddleware := middleware.NewJWTMiddleware(env.JWTSignatureKey)
		app.Use(JWTMiddleware.Middleware())
		// Routers
		routers.InitAuthRouters(ctx, db, app)
		transactionGroup := app.Group("/transactions")
		{
			routers.InitTransactionRouters(ctx, db, transactionGroup)
		}

		// Start server
		go func() {
			if err := app.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
				app.Logger.Fatal("shutting down the server")
			}
		}()

		// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
		// Use a buffered channel to avoid missing signals as recommended for signal.Notify
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := app.Shutdown(ctx); err != nil {
			app.Logger.Fatal(err)
		}
	} else {
		fmt.Println(err.Error())
	}
}
