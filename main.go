package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vsrecorder/decktype-api/internal/beta"
	"github.com/vsrecorder/decktype-api/internal/handlers"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Use(cors.New(cors.Config{
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Access-Control-Request-Method",
			"Authorization",
			"Content-Type",
		},
		AllowMethods: []string{
			"GET",
			"OPTIONS",
		},
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://local.vsrecorder.mobi",
			"https://decktype.vsrecorder.mobi",
		},
		AllowCredentials: false,
		MaxAge:           24 * time.Hour,
	}))

	r.GET(
		"/decktypes/:id",
		handlers.GetMc,
	)

	r.GET(
		"/decktypes/:id/mc",
		handlers.GetMc,
	)
	r.GET(
		"/decktypes/:id/m2a",
		handlers.GetM2a,
	)

	r.GET(
		"/decktypes/:id/m2",
		handlers.GetM2,
	)

	r.GET(
		"/decktypes/:id/m1",
		handlers.GetM1,
	)

	r.GET(
		"/api/v1beta/decktypes/:id",
		beta.GetM2a,
	)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:    ":8930",
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 3 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
