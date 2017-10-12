package usecase

import (
	"github.com/arielizuardi/gocrashandburn/crash/repository"
)

type CrashUsecase struct {
	Repo *repository.CrashRepository
}

func (c *CrashUsecase) IgniteErrors() error {
	return c.Repo.IgniteErrors()
}
