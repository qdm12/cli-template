package env

import (
	"os"

	"github.com/qdm12/cli-template/internal/config/settings"
)

// Read reads the settings from the environment variables.
func (s *Source) Read() (settings settings.Settings, err error) {
	settings.Path = os.Getenv("SOME_PATH")

	return settings, nil
}
