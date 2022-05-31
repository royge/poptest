// Package poptest is a utility package that can be used when testing with
// gobuffalo pop connection.
//
// The primary goal of this package is to override the database name from the
// database.yml programatically.
package poptest

import (
	"os"
	"path/filepath"

	pop "github.com/gobuffalo/pop/v6"
)

// Deets return pop connection details.
//
// env is the database environment you want to handle with.
// The cb parameter can be used to modidy connections details before returning
// it.
func Deets(env string, cb func(*pop.ConnectionDetails) error) (d *pop.ConnectionDetails, err error) {
	findConfigPath := func() (string, error) {
		for _, p := range pop.LookupPaths() {
			path, _ := filepath.Abs(filepath.Join(p, pop.ConfigName))
			if _, err := os.Stat(path); err == nil {
				return path, err
			}
		}
		return "", pop.ErrConfigFileNotFound
	}

	path, err := findConfigPath()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	deets, err := pop.ParseConfig(f)
	if err != nil {
		return nil, err
	}

	d = deets[env]

	if cb != nil {
		if err := cb(d); err != nil {
			return nil, err
		}
	}

	return d, nil
}
