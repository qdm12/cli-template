package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/qdm12/REPONAME/internal/models"
	"github.com/qdm12/REPONAME/internal/params"
	"github.com/qdm12/REPONAME/internal/setup"
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

	exitCode := 0
	select {
	case signal := <-signalsCh:
		fmt.Println("\nShutting down: signal", signal)
		exitCode = 1
		cancel()
		timer := time.NewTimer(time.Second)
		select {
		case <-errorCh:
			if !timer.Stop() {
				<-timer.C
			}
		case <-timer.C:
			fmt.Println("Shutdown timed out")
		}
	case err := <-errorCh:
		if err != nil {
			fmt.Println("Fatal error:", err)
			exitCode = 1
		}
		cancel()
	}
	os.Exit(exitCode)
}

func _main(ctx context.Context, args []string, buildInfo models.BuildInfo) error {
	fmt.Printf("🤖 Version %s (commit %s built on %s)\n",
		buildInfo.Version, buildInfo.Commit, buildInfo.BuildDate)

	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)
	repoPath := flagSet.String("path", ".", "path")
	if err := flagSet.Parse(args[1:]); err != nil {
		return err
	}

	fmt.Print("📁 Creating directory...")
	err = os.MkdirAll(path, 0700)
	if err != nil {
		fmt.Println("❌")
		return err
	}
	fmt.Println("✔️")

	return nil
}
