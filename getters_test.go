package config

import (
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGetters(t *testing.T) {
	t.Run("AllSettings", func(t *testing.T) {
		viper.Reset()
		viper.Set("app.env", "local")

		expected := map[string]interface{}{
			"app": map[string]interface{}{
				"env": "local",
			},
		}

		actual := AllSettings()

		assert.Equal(t, expected, actual)
	})

	t.Run("GetString", func(t *testing.T) {
		viper.Reset()
		viper.Set("app.env", "local")

		assert.Equal(t, GetString("app.env"), "local")
		assert.Equal(t, GetString("app.environment"), "")
	})

	t.Run("GetStringOr", func(t *testing.T) {
		viper.Reset()
		viper.Set("app.env", "local")

		assert.Equal(t, GetStringOr("app.env", "dev"), "local")
		assert.Equal(t, GetStringOr("app.environment", "dev"), "dev")
	})

	t.Run("GetBool", func(t *testing.T) {
		viper.Reset()
		viper.Set("app.debug", true)

		assert.Equal(t, GetBool("app.debug"), true)
		assert.Equal(t, GetBool("app.is_debug"), false)
	})

	t.Run("GetBoolOr", func(t *testing.T) {
		viper.Reset()
		viper.Set("app.debug", true)

		assert.Equal(t, GetBoolOr("app.debug", false), true)
		assert.Equal(t, GetBoolOr("app.is_debug", true), true)
	})

	t.Run("GetInt", func(t *testing.T) {
		viper.Reset()
		viper.Set("http.port", 8080)

		assert.Equal(t, GetInt("http.port"), 8080)
		assert.Equal(t, GetInt("https.port"), 0)
	})

	t.Run("GetIntOr", func(t *testing.T) {
		viper.Reset()
		viper.Set("http.port", 8080)

		assert.Equal(t, GetIntOr("http.port", 3000), 8080)
		assert.Equal(t, GetIntOr("https.port", 8443), 8443)
	})

	t.Run("GetInt32", func(t *testing.T) {
		viper.Reset()
		viper.Set("http.port", 8080)

		assert.Equal(t, GetInt32("http.port"), int32(8080))
		assert.Equal(t, GetInt32("https.port"), int32(0))
	})

	t.Run("GetInt32Or", func(t *testing.T) {
		viper.Reset()
		viper.Set("http.port", 8080)

		assert.Equal(t, GetInt32Or("http.port", 3000), int32(8080))
		assert.Equal(t, GetInt32Or("https.port", 8443), int32(8443))
	})

	t.Run("GetInt64", func(t *testing.T) {
		viper.Reset()
		viper.Set("http.port", 8080)

		assert.Equal(t, GetInt64("http.port"), int64(8080))
		assert.Equal(t, GetInt64("https.port"), int64(0))
	})

	t.Run("GetInt64Or", func(t *testing.T) {
		viper.Reset()
		viper.Set("http.port", 8080)

		assert.Equal(t, GetInt64Or("http.port", 3000), int64(8080))
		assert.Equal(t, GetInt64Or("https.port", 8443), int64(8443))
	})

	t.Run("GetUint", func(t *testing.T) {
		viper.Reset()
		viper.Set("http.port", 8080)

		assert.Equal(t, GetUint("http.port"), uint(8080))
		assert.Equal(t, GetUint("https.port"), uint(0))
	})

	t.Run("GetUintOr", func(t *testing.T) {
		viper.Reset()
		viper.Set("http.port", 8080)

		assert.Equal(t, GetUintOr("http.port", 3000), uint(8080))
		assert.Equal(t, GetUintOr("https.port", 8443), uint(8443))
	})

	t.Run("GetUint16", func(t *testing.T) {
		viper.Reset()
		viper.Set("http.port", 8080)

		assert.Equal(t, GetUint16("http.port"), uint16(8080))
		assert.Equal(t, GetUint16("https.port"), uint16(0))
	})

	t.Run("GetUint16Or", func(t *testing.T) {
		viper.Reset()
		viper.Set("http.port", 8080)

		assert.Equal(t, GetUint16Or("http.port", 3000), uint16(8080))
		assert.Equal(t, GetUint16Or("https.port", 8443), uint16(8443))
	})

	t.Run("GetUint32", func(t *testing.T) {
		viper.Reset()
		viper.Set("http.port", 8080)

		assert.Equal(t, GetUint32("http.port"), uint32(8080))
		assert.Equal(t, GetUint32("https.port"), uint32(0))
	})

	t.Run("GetUint32Or", func(t *testing.T) {
		viper.Reset()
		viper.Set("http.port", 8080)

		assert.Equal(t, GetUint32Or("http.port", 3000), uint32(8080))
		assert.Equal(t, GetUint32Or("https.port", 8443), uint32(8443))
	})

	t.Run("GetUint64", func(t *testing.T) {
		viper.Reset()
		viper.Set("http.port", 8080)

		assert.Equal(t, GetUint64("http.port"), uint64(8080))
		assert.Equal(t, GetUint64("https.port"), uint64(0))
	})

	t.Run("GetUint64Or", func(t *testing.T) {
		viper.Reset()
		viper.Set("http.port", 8080)

		assert.Equal(t, GetUint64Or("http.port", 3000), uint64(8080))
		assert.Equal(t, GetUint64Or("https.port", 8443), uint64(8443))
	})

	t.Run("GetFloat64", func(t *testing.T) {
		viper.Reset()
		viper.Set("stable_diffusion.cfg", 4.5)

		assert.Equal(t, GetFloat64("stable_diffusion.cfg"), 4.5)
		assert.Equal(t, GetFloat64("stable_diffusion.prompt_strength"), float64(0))
	})

	t.Run("GetFloat64Or", func(t *testing.T) {
		viper.Reset()
		viper.Set("stable_diffusion.cfg", 4.5)

		assert.Equal(t, GetFloat64Or("stable_diffusion.cfg", 7), 4.5)
		assert.Equal(t, GetFloat64Or("stable_diffusion.prompt_strength", 0.85), 0.85)
	})

	t.Run("GetDuration", func(t *testing.T) {
		viper.Reset()
		viper.Set("services.thrift.request_timeout_s", 60)

		assert.Equal(t, GetDuration("services.thrift.request_timeout_s"), time.Duration(60))
		assert.Equal(t, GetDuration("services.thrift.request_timeout"), time.Duration(0))
	})

	t.Run("GetUint64Or", func(t *testing.T) {
		viper.Reset()
		viper.Set("services.thrift.request_timeout_s", 60)

		assert.Equal(t, GetDurationOr("services.thrift.request_timeout_s", 30), time.Duration(60))
		assert.Equal(t, GetDurationOr("services.thrift.request_timeout", 60), time.Duration(60))
	})
}
