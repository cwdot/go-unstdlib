package environment

import (
	"os"
	"strconv"
	"time"
)

type envMap map[string]Value

func (e envMap) Set(key, value string, firstWins bool) {
	if _, ok := e[key]; !ok || !firstWins {
		e[key] = Value(value)
	}
}

func (e envMap) SetEnviron(key string, firstWins bool) {
	e.Set(key, os.Getenv(key), firstWins)
}

type Value string

func (v Value) S() string {
	return string(v)
}

func (v Value) Int(defaultValue int) (int, error) {
	value, err := strconv.Atoi(string(v))
	if err == nil {
		return value, nil
	}
	return defaultValue, err
}

func (v Value) MustInt(defaultValue int) int {
	value, err := v.Int(defaultValue)
	if err != nil {
		return defaultValue
	}
	return value
}

func (v Value) MustIntP(defaultValue int) *int {
	value := v.MustInt(defaultValue)
	return &value
}

func (v Value) MustBool(defaultValue bool) bool {
	vs := string(v)
	if vs == "" {
		return defaultValue
	}
	return vs == "1" || vs == "true"
}

func (v Value) MustBoolP(defaultValue bool) *bool {
	value := v.MustBool(defaultValue)
	return &value
}

func (v Value) Duration(defaultDuration time.Duration) (time.Duration, error) {
	value, err := strconv.Atoi(string(v))
	if err == nil {
		return time.Duration(value), nil
	}
	return defaultDuration, err
}

func (v Value) MustDuration(defaultDuration time.Duration) time.Duration {
	value, err := v.Duration(defaultDuration)
	if err != nil {
		return defaultDuration
	}
	return value
}

func (v Value) MustDurationP(defaultDuration time.Duration) *time.Duration {
	duration := v.MustDuration(defaultDuration)
	return &duration
}
