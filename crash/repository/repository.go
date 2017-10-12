package repository

import (
	"github.com/go-errors/errors"
)

type CrashRepository struct {
}

func (c *CrashRepository) IgniteErrors() error {
	err := errors.New(`Something happened in repository`)
	if err != nil {
		//logrus.Error(stacktrace.Propagate(err, "Sh*t happens in repo"))
		return errors.Wrap(err, 1)
	}

	return nil
}
