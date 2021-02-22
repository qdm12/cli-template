package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/qdm12/REPONAME/internal/models"
)

//nolint:gochecknoglobals
var (
	version   = "unknown"
	commit    = "unknown"
	buildDate = "an unknown date"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	buildInfo := models.BuildInfo{
		Version:   version,
		Commit:    commit,
		BuildDate: buildDate,
	}

	errorCh := make(chan error)
	go func() {
		errorCh <- _main(ctx, os.Args, buildInfo)
	}()

	signalsCh := make(chan os.Signal, 1)
	signal.Notify(signalsCh,
		syscall.SIGINT,
		syscall.SIGTERM,
		os.Interrupt,
	)

	select {
	case err := <-errorCh:
		close(errorCh)
		if err == nil { // expected exit such as healthcheck
			os.Exit(0)
		}
		fmt.Println("Fatal error:", err)
		os.Exit(1)
	case signal := <-signalsCh:
		fmt.Println("\nShutting down: signal", signal)
	}

	cancel()

	const shutdownGracePeriod = time.Second
	timer := time.NewTimer(shutdownGracePeriod)
	select {
	case <-errorCh:
		if !timer.Stop() {
			<-timer.C
		}
	case <-timer.C:
		fmt.Println("Shutdown timed out")
	}

	os.Exit(1)
}

func _main(ctx context.Context, args []string, buildInfo models.BuildInfo) error {
	fmt.Printf("🤖 Version %s (commit %s built on %s)\n",
		buildInfo.Version, buildInfo.Commit, buildInfo.BuildDate)

	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)
	pathPtr := flagSet.String("path", ".", "path")
	if err := flagSet.Parse(args[1:]); err != nil {
		return err
	}
	path := *pathPtr

	fmt.Print("📁 Creating directory...")
	err := os.MkdirAll(path, 0700)
	if err != nil {
		fmt.Println("❌")
		return err
	}
	fmt.Println("✔️")

	return nil
}
