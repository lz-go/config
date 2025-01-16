package config

import "github.com/spf13/viper"

// Reset is intended for testing, will reset all to default settings.
// In the public interface for the viper package so applications
// can use it in their testing as well.
// It is a wrapper around viper.Reset.
var Reset = viper.Reset

// Set sets the value for the key in the override register.
// Set is case-insensitive for a key.
// Will be used instead of values obtained via
// flags, config file, ENV, default, or key/value store.
// It is a wrapper around viper.Set.
var Set = viper.Set
