package cmd

import (
	"fmt"
	"os"
	"time"

	"context"
	"net"
	"net/http"
	"sync"

	"github.com/Ow1Dev/QustionOfToday/internal/config"
	"github.com/Ow1Dev/QustionOfToday/internal/database"
	"github.com/Ow1Dev/QustionOfToday/internal/repository"
	"github.com/Ow1Dev/QustionOfToday/internal/server"
	"github.com/rs/zerolog"
	"github.com/urfave/cli/v2"
)

var CmdWeb = &cli.Command{
	Name:   "web",
	Action: runWeb,
}

func runWeb(ctx *cli.Context) error {
	debug := false

	config := config.Config{
		Host: "0.0.0.0",
		Port: "3000",
	}

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		With().
		Timestamp().
		Logger()

	if debug {
		logger.Debug().Msg("Debug is turned on")
	}

	db, err := database.Connect(ctx.Context)
	if err != nil {
		return err
	}

	repo := repository.New(db)

	srv := server.NewServer(
		&logger,
		&config,
		repo,
		debug,
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
