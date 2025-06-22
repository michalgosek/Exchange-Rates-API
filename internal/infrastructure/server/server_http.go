package server

import (
	"context"
	"exchange-rates-api/internal/app"
	"exchange-rates-api/internal/ports"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Config struct {
	Addr string `mapstructure:"addr"`
	Port int    `mapstructure:"port"`
}

func (c *Config) SocketAddr() string { return fmt.Sprintf("%s:%d", c.Addr, c.Port) }

func RunHTTPWithGracefulShutdown(cfg *Config, app *app.Application) <-chan struct{} {
	srv := http.Server{
		Addr:           cfg.SocketAddr(),
		Handler:        ports.NewHTTP(app),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	done := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}

		log.Print("HTTP graceful shutdown")
		close(done)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	return done
}
