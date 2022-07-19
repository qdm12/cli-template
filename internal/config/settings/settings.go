package settings

import (
	"fmt"
	"os"

	"github.com/qdm12/gosettings/defaults"
	"github.com/qdm12/gosettings/merge"
	"github.com/qdm12/gotree"
)

type Settings struct {
	Path string
}

// MergeWith sets only zero-ed fields in the receiving settings
// with fields from the other settings given.
func (s *Settings) MergeWith(other Settings) {
	s.Path = merge.String(s.Path, other.Path)
}

// OverrideWith sets fields in the receiving settings
// from non-zero fields from the other settings given.
func (s *Settings) OverrideWith(other Settings) {
	s.Path = merge.String(s.Path, other.Path)
}

// SetDefaults sets the defaults to all the zero-ed fields
// in the receiving settings.
func (s *Settings) SetDefaults() {
	s.Path = defaults.String(s.Path, "./path")
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
