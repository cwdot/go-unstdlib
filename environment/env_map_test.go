package environment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestEnvMap(t *testing.T) {
	env := envMap{}
	t.Run("SetOverwrite", func(t *testing.T) {
		env.Set("foo1", "bar", false)
		require.Equal(t, "bar", env["foo1"].S())

		env.Set("foo1", "baz", false)
		require.Equal(t, "baz", env["foo1"].S())
	})
	t.Run("SetFirstWins", func(t *testing.T) {

		env.Set("foo2", "qux", true)
		require.Equal(t, "qux", env["foo2"].S())

		env.Set("foo2", "quux", true)
		require.Equal(t, "qux", env["foo2"].S())
	})
}

func TestEnvMap_Int(t *testing.T) {
	env := envMap{}
	t.Run("Valid", func(t *testing.T) {
		env.Set("valid", "1", false)
		value, ok := env["valid"]
		require.True(t, ok)
		require.Equal(t, "1", value.S())
		require.Equal(t, 1, value.MustInt(42))
		output, err := value.Int(42)
		require.NoError(t, err)
		require.Equal(t, 1, output)
	})
	t.Run("Invalid", func(t *testing.T) {
		env.Set("invalid", "x", false)
		value, ok := env["invalid"]
		require.True(t, ok)
		require.Equal(t, "x", value.S())

		output := value.MustInt(42)
		require.Equal(t, 42, output)

		output, err := value.Int(42)
		require.Error(t, err)
		require.Equal(t, 42, output)
	})
	t.Run("Empty", func(t *testing.T) {
		env.Set("empty", "", false)
		value, ok := env["empty"]
		require.True(t, ok)
		require.Equal(t, "", value.S())

		output := value.MustInt(42)
		require.Equal(t, 42, output)
	})
	t.Run("P", func(t *testing.T) {
		env.Set("intp", "1", false)
		value, ok := env["intp"]
		require.True(t, ok)
		require.Equal(t, "1", value.S())
		output := value.MustIntP(42)
		require.Equal(t, 1, *output)
	})
}

func TestEnvMap_Bool(t *testing.T) {
	env := envMap{}
	t.Run("ValidTrue", func(t *testing.T) {
		env.Set("valid-true", "true", false)
		value, ok := env["valid-true"]
		require.True(t, ok)
		require.Equal(t, "true", value.S())
		require.Equal(t, true, value.MustBool(false))
	})
	t.Run("Valid1", func(t *testing.T) {
		env.Set("valid-1", "1", false)
		value, ok := env["valid-1"]
		require.True(t, ok)
		require.Equal(t, "1", value.S())
		require.Equal(t, true, value.MustBool(false))
	})
	t.Run("Invalid", func(t *testing.T) {
		env.Set("invalid", "x", false)
		value, ok := env["invalid"]
		require.True(t, ok)
		require.Equal(t, "x", value.S())

		output := value.MustBool(true)
		require.Equal(t, false, output) // value is not empty; it does not equal true
	})
	t.Run("Empty", func(t *testing.T) {
		env.Set("empty", "", false)
		value, ok := env["empty"]
		require.True(t, ok)
		require.Equal(t, "", value.S())

		output := value.MustBool(true)
		require.Equal(t, true, output)
	})
	t.Run("P", func(t *testing.T) {
		env.Set("boolp", "1", false)
		value, ok := env["boolp"]
		require.True(t, ok)
		require.Equal(t, "1", value.S())
		require.True(t, *value.MustBoolP(true))
	})
}

func TestEnvMap_Duration(t *testing.T) {
	const defaultDuration = 4000000 // 4 seconds
	env := envMap{}
	t.Run("Valid", func(t *testing.T) {
		env.Set("valid", "1000000", false)
		value, ok := env["valid"]
		require.True(t, ok)
		require.Equal(t, "1000000", value.S())
		require.Equal(t, int64(1), value.MustDuration(defaultDuration).Milliseconds())

		output, err := value.Duration(defaultDuration)
		require.NoError(t, err)
		require.Equal(t, int64(1), output.Milliseconds())
	})
	t.Run("Empty", func(t *testing.T) {
		env.Set("empty", "", false)
		value, ok := env["empty"]
		require.True(t, ok)
		require.Equal(t, "", value.S())

		output := value.MustDuration(defaultDuration)
		require.Equal(t, time.Duration(defaultDuration), output)
	})
	t.Run("Invalid", func(t *testing.T) {
		env.Set("invalid", "x", false)
		value, ok := env["invalid"]
		require.True(t, ok)
		require.Equal(t, "x", value.S())

		output := value.MustDuration(defaultDuration)
		require.Equal(t, time.Duration(defaultDuration), output)

		output, err := value.Duration(defaultDuration)
		require.Error(t, err)
		require.Equal(t, int64(4), output.Milliseconds())
	})
	t.Run("P", func(t *testing.T) {
		env.Set("valid", "1000000", false)
		value, ok := env["valid"]
		require.True(t, ok)
		require.Equal(t, "1000000", value.S())
		actual := value.MustDurationP(4000000)
		require.Equal(t, int64(1), actual.Milliseconds())
		require.Equal(t, int64(1), (*actual).Milliseconds())
	})
}
