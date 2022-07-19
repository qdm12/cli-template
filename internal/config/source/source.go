package source

import (
	"github.com/qdm12/cli-template/internal/config/settings"
)

// Source is a settings source.
type Source interface {
	// String returns the name of the source.
	String() string
	// Read reads the settings from the source.
	Read() (s settings.Settings, err error)
}
