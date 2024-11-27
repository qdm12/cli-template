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

	"github.com/qdm12/cli-template/internal/config"
	"github.com/qdm12/cli-template/internal/models"
	"github.com/qdm12/gosettings/reader"
)

//nolint:gochecknoglobals
var (
	version = "unknown"
	commit  = "unknown"
	date    = "an unknown date"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	buildInfo := models.BuildInfo{
		Version: version,
		Commit:  commit,
		Date:    date,
	}

	errorCh := make(chan error)
	go func() {
		errorCh <- _main(ctx, buildInfo, os.Stdout, os.Stdin)
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
	stdout io.Writer, _ io.Reader,
) (err error) {
	versionMessage := fmt.Sprintf("🤖 Version %s (commit %s built on %s)",
		buildInfo.Version, buildInfo.Commit, buildInfo.Date)
	fmt.Fprintln(stdout, versionMessage)

	reader := reader.New(reader.Settings{})
	var settings config.Settings
	err = settings.Read(reader)
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
	const dirPerms fs.FileMode = 0o700
	err = os.MkdirAll(settings.Path, dirPerms)
	if err != nil {
		fmt.Fprintln(stdout, " ❌")
		return err
	}
	fmt.Fprintln(stdout, " ✔️")

	return nil
}
