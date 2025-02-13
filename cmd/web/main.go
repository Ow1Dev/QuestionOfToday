package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/Ow1Dev/QustionOfToday/internal/config"
	"github.com/Ow1Dev/QustionOfToday/internal/database"
	"github.com/Ow1Dev/QustionOfToday/internal/repository"
	"github.com/Ow1Dev/QustionOfToday/internal/server"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

func main() {
	if err := run(os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(stdout io.Writer, args []string) error {
	_ = stdout
	_ = args

	debug := flag.Bool("Debug", false, "Turn on debug")
	flag.Parse()

	godotenv.Load()

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	config := config.Config{
		Host: "0.0.0.0",
		Port: "3000",
	}

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		With().
		Timestamp().
		Logger()

	if *debug {
		logger.Debug().Msg("Debug is turned on")
	}

	db, err := database.Connect(ctx)
	if err != nil {
		return err
	}

	repo := repository.New(db)

	srv := server.NewServer(
		&logger,
		&config,
		repo,
		*debug,
	)

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(config.Host, config.Port),
		Handler: srv,
	}

	go func() {
		logger.Info().Msgf("listening on %s", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil
}
