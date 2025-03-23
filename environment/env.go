package environment

import (
	"os"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type ReadOpts struct {
	Paths                  []string
	IgnoreMissing          bool
	IncludeAllEnvironment  bool
	IncludeEnvironmentKeys []string
	FirstKeyWins           bool
}

func WithPaths(paths []string) func(*ReadOpts) {
	return func(o *ReadOpts) {
		o.Paths = paths
	}
}

func WithIgnoreMissing() func(*ReadOpts) {
	return func(o *ReadOpts) {
		o.IgnoreMissing = true
	}
}

func WithIncludeAllEnvironment() func(*ReadOpts) {
	return func(o *ReadOpts) {
		o.IncludeAllEnvironment = true
	}
}

func WithIncludeEnvironmentKeys(keys []string) func(*ReadOpts) {
	return func(o *ReadOpts) {
		o.IncludeEnvironmentKeys = keys
	}
}

func Read(opts ...func(*ReadOpts)) (map[string]Value, error) {
	options := &ReadOpts{}
	for _, o := range opts {
		o(options)
	}

	env := make(envMap)

	if options.IncludeAllEnvironment {
		for _, line := range os.Environ() {
			tokens := strings.Split(line, "=")
			key := tokens[0]
			value := tokens[1]
			env.Set(key, value, options.FirstKeyWins)
		}
	}
	if options.IncludeEnvironmentKeys != nil {
		for _, key := range options.IncludeEnvironmentKeys {
			value := os.Getenv(key)
			env.Set(key, value, options.FirstKeyWins)
		}
	}

	for _, p := range options.Paths {
		penv, err := load(p)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) && options.IgnoreMissing {
				continue
			}
			return nil, errors.Wrapf(err, "cannot load %s", p)
		}
		for k, v := range penv {
			env.Set(k, v, options.FirstKeyWins)
		}
	}
	return env, nil
}

func load(credentialsPath string) (map[string]string, error) {
	env, err := godotenv.Read(credentialsPath)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot find %s", credentialsPath)
	}
	return env, nil
}

func Validate(env map[string]Value, keys []string) error {
	var errs *multierror.Error
	for _, key := range keys {
		if _, ok := env[key]; !ok {
			errs = multierror.Append(errs, errors.Errorf("missing %s", key))
		}
	}
	return errs.ErrorOrNil()
}
