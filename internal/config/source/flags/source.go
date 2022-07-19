package flags

import (
	"github.com/qdm12/cli-template/internal/config/source"
)

var _ source.Source = (*Source)(nil)

// Source is a command line flags settings source.
type Source struct {
	args []string
}

// New returns a new command line flags source.
func New(args []string) *Source {
	return &Source{
		args: args,
	}
}

func (s *Source) String() string {
	return "command line flags"
}
