package astiopendata

import (
	"flag"
	"time"
)

// Flags
var (
	RetryMax   = flag.Int("ratp-client-retry-max", 0, "the RATP client max retry")
	RetrySleep = flag.Duration("ratp-client-retry-sleep", 0, "the RATP client max sleep")
)

// Configuration represents the RATP configuration
type Configuration struct {
	RetryMax   int           `toml:"retry_max"`
	RetrySleep time.Duration `toml:"retry_sleep"`
}

// FlagConfig generates a Configuration based on flags
func FlagConfig() Configuration {
	return Configuration{
		RetryMax:   *RetryMax,
		RetrySleep: *RetrySleep,
	}
}
