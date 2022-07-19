package mux

import (
	"github.com/qdm12/cli-template/internal/config/source"
)

var _ source.Source = (*Source)(nil)

// Source is a settings source multiplexing
// the sources given together.
type Source struct {
	sources []source.Source
}

// New returns a new mux source reading from
// each of the sources in the order they were
// added with precedence.
func New(sources ...source.Source) *Source {
	return &Source{
		sources: sources,
	}
}

func (s *Source) String() string {
	return "mux"
}
