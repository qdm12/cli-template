package config

import (
	"fmt"
	"os"

	"github.com/qdm12/gosettings"
	"github.com/qdm12/gosettings/reader"
	"github.com/qdm12/gotree"
)

type Settings struct {
	Path string
}

// OverrideWith sets fields in the receiving settings
// from non-zero fields from the other settings given.
func (s *Settings) OverrideWith(other Settings) {
	s.Path = gosettings.OverrideWithComparable(s.Path, other.Path)
}

// SetDefaults sets the defaults to all the zero-ed fields
// in the receiving settings.
func (s *Settings) SetDefaults() {
	s.Path = gosettings.DefaultComparable(s.Path, "./path")
}

// Validate validates all the settings are correct.
// Note `.SetDefaults()` must be called to ensure all
// the fields are not their zeroed value such as `nil`.
func (s *Settings) Validate() (err error) {
	_, err = os.Stat(s.Path)
	if err != nil {
		return fmt.Errorf("path: %w", err)
	}

	return nil
}

func (s *Settings) Read(reader *reader.Reader) (err error) {
	s.Path = reader.String("SOME_PATH")
	return nil
}

// toLinesNode returns a gotree.Node with the settings
// as a formatted tree node.
func (s *Settings) toLinesNode() *gotree.Node {
	node := gotree.New("Settings:")
	node.Appendf("Path: %s", s.Path)
	return node
}

func (s Settings) String() string {
	return s.toLinesNode().String()
}
