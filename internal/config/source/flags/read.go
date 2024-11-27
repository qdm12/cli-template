package flags

import (
	"flag"
	"fmt"

	"github.com/qdm12/cli-template/internal/config/settings"
)

// Read reads the settings from the cli flags.
func (source *Source) Read() (settings settings.Settings, err error) {
	flagSet, flagSettings, rawStrings := configureFlagSet(source.args[0])
	if err := flagSet.Parse(source.args[1:]); err != nil {
		return flagSettings, err
	}

	err = postProcessRawStrings(rawStrings, &flagSettings)
	if err != nil {
		return flagSettings, err
	}

	flagSet.Visit(func(f *flag.Flag) {
		visitFlag(f.Name, &settings, flagSettings)
	})

	return settings, nil
}

type rawStrings struct{}

func configureFlagSet(flagSetName string) (flagSet *flag.FlagSet,
	flagSettings settings.Settings, rawStrings rawStrings, //nolint:unparam
) {
	flagSet = flag.NewFlagSet(flagSetName, flag.ExitOnError)

	// set pointers to non-nil values and
	// use default values for flag documentation
	flagSettings.SetDefaults()

	// note the default values here are only for information purposes,
	// the actual default are set in the settings package.
	flagSet.StringVar(&flagSettings.Path, "path", flagSettings.Path, "Path.")

	return flagSet, flagSettings, rawStrings
}

func postProcessRawStrings(rawStrings rawStrings, settings *settings.Settings) (err error) { //nolint:revive
	return nil
}

func visitFlag(flagName string, destination *settings.Settings,
	source settings.Settings,
) {
	switch flagName { //nolint:gocritic
	case "path":
		destination.Path = source.Path
		return
	}

	panic(fmt.Sprintf("flag not added to switch case: %s", flagName))
}
