package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/arielizuardi/gocrashandburn/crash/http"
	"github.com/arielizuardi/gocrashandburn/crash/repository"
	"github.com/arielizuardi/gocrashandburn/crash/usecase"
	"github.com/labstack/echo"
)

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})
	e := echo.New()
	r := &repository.CrashRepository{}
	u := &usecase.CrashUsecase{r}

	http.Init(e, u)
	e.Start(`:7723`)
}
