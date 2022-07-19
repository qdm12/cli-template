package mux

import (
	"fmt"

	"github.com/qdm12/cli-template/internal/config/settings"
)

// Read reads the settings from each of the sources,
// in the order they were added with precedence.
func (s *Source) Read() (settings settings.Settings, err error) {
	for _, source := range s.sources {
		newSettings, err := source.Read()
		if err != nil {
			return settings, fmt.Errorf("%s source: %w", source, err)
		}
		settings.MergeWith(newSettings)
	}

	return settings, nil
}
