package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/fs"
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
	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	buildInfo := models.BuildInfo{
		Version:   version,
		Commit:    commit,
		BuildDate: buildDate,
	}

	errorCh := make(chan error)
	go func() {
		errorCh <- _main(ctx, buildInfo, os.Args, os.Stdout, os.Stdin)
	}()

	select {
	case err := <-errorCh:
		close(errorCh)
		if err == nil { // expected exit
			os.Exit(0)
		}
		fmt.Println("Fatal error:", err)
		os.Exit(1)
	case <-ctx.Done():
		stop()
		fmt.Println("\nShutting down: OS signal received")
	}

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

func _main(_ context.Context, buildInfo models.BuildInfo,
	args []string, stdout io.Writer, _ io.Reader) error {
	versionMessage := fmt.Sprintf("🤖 Version %s (commit %s built on %s)",
		buildInfo.Version, buildInfo.Commit, buildInfo.BuildDate)
	fmt.Fprint(stdout, versionMessage)

	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)
	pathPtr := flagSet.String("path", ".", "path")
	if err := flagSet.Parse(args[1:]); err != nil {
		return err
	}
	path := *pathPtr

	fmt.Fprint(stdout, "📁 Creating directory...")
	const dirPerms fs.FileMode = 0700
	err := os.MkdirAll(path, dirPerms)
	if err != nil {
		fmt.Fprint(stdout, "❌")
		return err
	}
	fmt.Fprint(stdout, "✔️")

	return nil
}
