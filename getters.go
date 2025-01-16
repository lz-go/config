package config

import (
	"time"

	"github.com/spf13/viper"
)

// AllSettings merges all settings and returns them as a map[string]any.
// It is a wrapper around viper.AllSettings.
var AllSettings = viper.AllSettings

// GetString retrieves a string value from Viper using the given key.
// It is a wrapper around viper.GetString.
var GetString = viper.GetString

// GetBool retrieves a boolean value from Viper using the given key.
// It is a wrapper around viper.GetBool.
var GetBool = viper.GetBool

// GetInt retrieves an integer value from Viper using the given key.
// It is a wrapper around viper.GetInt.
var GetInt = viper.GetInt

// GetInt32 retrieves a 32-bit integer value from Viper using the given key.
// It is a wrapper around viper.GetInt32.
var GetInt32 = viper.GetInt32

// GetInt64 retrieves a 64-bit integer value from Viper using the given key.
// It is a wrapper around viper.GetInt64.
var GetInt64 = viper.GetInt64

// GetUint retrieves an unsigned integer value from Viper using the given key.
// It is a wrapper around viper.GetUint.
var GetUint = viper.GetUint

// GetUint16 retrieves a 16-bit unsigned integer value from Viper using the given key.
// It is a wrapper around viper.GetUint16.
var GetUint16 = viper.GetUint16

// GetUint32 retrieves a 32-bit unsigned integer value from Viper using the given key.
// It is a wrapper around viper.GetUint32.
var GetUint32 = viper.GetUint32

// GetUint64 retrieves a 64-bit unsigned integer value from Viper using the given key.
// It is a wrapper around viper.GetUint64.
var GetUint64 = viper.GetUint64

// GetFloat64 retrieves a float64 value from Viper using the given key.
// It is a wrapper around viper.GetFloat64.
var GetFloat64 = viper.GetFloat64

// GetDuration retrieves a time.Duration value from Viper using the given key.
// It is a wrapper around viper.GetDuration.
var GetDuration = viper.GetDuration

// GetStringOr retrieves a string value from Viper using the given key.
// If the key is not set, it returns the provided defaultValue.
func GetStringOr(key string, defaultValue string) string {
	if viper.IsSet(key) {
		return viper.GetString(key)
	}

	return defaultValue
}

// GetBoolOr retrieves a boolean value from Viper using the given key.
// If the key is not set, it returns the provided defaultValue.
func GetBoolOr(key string, defaultValue bool) bool {
	if viper.IsSet(key) {
		return viper.GetBool(key)
	}

	return defaultValue
}

// GetIntOr retrieves an integer value from Viper using the given key.
// If the key is not set, it returns the provided defaultValue.
func GetIntOr(key string, defaultValue int) int {
	if viper.IsSet(key) {
		return viper.GetInt(key)
	}

	return defaultValue
}

// GetInt32Or retrieves a 32-bit integer value from Viper using the given key.
// If the key is not set, it returns the provided defaultValue.
func GetInt32Or(key string, defaultValue int32) int32 {
	if viper.IsSet(key) {
		return viper.GetInt32(key)
	}

	return defaultValue
}

// GetInt64Or retrieves a 64-bit integer value from Viper using the given key.
// If the key is not set, it returns the provided defaultValue.
func GetInt64Or(key string, defaultValue int64) int64 {
	if viper.IsSet(key) {
		return viper.GetInt64(key)
	}

	return defaultValue
}

// GetUintOr retrieves an unsigned integer value from Viper using the given key.
// If the key is not set, it returns the provided defaultValue.
func GetUintOr(key string, defaultValue uint) uint {
	if viper.IsSet(key) {
		return viper.GetUint(key)
	}

	return defaultValue
}

// GetUint16Or retrieves a 16-bit unsigned integer value from Viper using the given key.
// If the key is not set, it returns the provided defaultValue.
func GetUint16Or(key string, defaultValue uint16) uint16 {
	if viper.IsSet(key) {
		return viper.GetUint16(key)
	}

	return defaultValue
}

// GetUint32Or retrieves a 32-bit unsigned integer value from Viper using the given key.
// If the key is not set, it returns the provided defaultValue.
func GetUint32Or(key string, defaultValue uint32) uint32 {
	if viper.IsSet(key) {
		return viper.GetUint32(key)
	}

	return defaultValue
}

// GetUint64Or retrieves a 64-bit unsigned integer value from Viper using the given key.
// If the key is not set, it returns the provided defaultValue.
func GetUint64Or(key string, defaultValue uint64) uint64 {
	if viper.IsSet(key) {
		return viper.GetUint64(key)
	}

	return defaultValue
}

// GetFloat64Or retrieves a float64 value from Viper using the given key.
// If the key is not set, it returns the provided defaultValue.
func GetFloat64Or(key string, defaultValue float64) float64 {
	if viper.IsSet(key) {
		return viper.GetFloat64(key)
	}

	return defaultValue
}

// GetDurationOr retrieves a time.Duration value from Viper using the given key.
// If the key is not set, it returns the provided defaultValue.
func GetDurationOr(key string, defaultValue time.Duration) time.Duration {
	if viper.IsSet(key) {
		return viper.GetDuration(key)
	}

	return defaultValue
}
