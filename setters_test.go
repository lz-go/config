package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetters(t *testing.T) {
	t.Run("Set", func(t *testing.T) {
		Reset()

		Set("app.env", "local")
		assert.Equal(t, "local", GetString("app.env"))
	})

	t.Run("Reset", func(t *testing.T) {
		Reset()

		Set("app.env", "local")
		assert.Equal(t, "local", GetString("app.env"))

		Reset()
		assert.Equal(t, "", GetString("app.env"))
	})
}
