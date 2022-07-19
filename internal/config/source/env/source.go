package env

import "github.com/qdm12/cli-template/internal/config/source"

var _ source.Source = (*Source)(nil)

// Source is an environment variables settings source.
type Source struct{}

// New returns a new environment variables source.
func New() *Source {
	return &Source{}
}

func (s *Source) String() string {
	return "environment variables"
}
