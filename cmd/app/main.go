package main

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/qdm12/cli-template/internal/config/source/env"
	"github.com/qdm12/cli-template/internal/config/source/flags"
	"github.com/qdm12/cli-template/internal/config/source/mux"
	"github.com/qdm12/cli-template/internal/models"
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
	fmt.Fprintln(stdout, versionMessage)

	flags := flags.New(args)
	env := env.New()
	mux := mux.New(flags, env)

	settings, err := mux.Read()
	if err != nil {
		return fmt.Errorf("reading settings: %w", err)
	}
	settings.SetDefaults()
	err = settings.Validate()
	if err != nil {
		return fmt.Errorf("validating settings: %w", err)
	}

	fmt.Fprintln(stdout, settings)

	fmt.Fprint(stdout, "📁 Creating directory...")
	const dirPerms fs.FileMode = 0700
	err = os.MkdirAll(settings.Path, dirPerms)
	if err != nil {
		fmt.Fprintln(stdout, " ❌")
		return err
	}
	fmt.Fprintln(stdout, " ✔️")

	return nil
}
