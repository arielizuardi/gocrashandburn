package http

import (
	"net/http"

	"github.com/arielizuardi/gocrashandburn/crash/usecase"
	"github.com/go-errors/errors"
	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"error"`
}

type ErrorWrapper struct {
	Err         string        `json:"error"`
	StackTraces []*StackFrame `json:"stack_traces"`
}

type StackFrame struct {
	File       string `json:"file"`
	Name       string `json:"name"`
	LineNumber int    `json:"line"`
	Package    string `json:"pkg"`
}

// CrashHTTPHandler ...
type CrashHTTPHandler struct {
	CrashUsecase *usecase.CrashUsecase
}

// HandleCrash ...
func (pl *CrashHTTPHandler) HandleCrash(c echo.Context) error {
	err := pl.CrashUsecase.IgniteErrors()
	if err != nil {
		stackFrames := []*StackFrame{}
		for _, frame := range err.(*errors.Error).StackFrames() {
			stackFrames = append(stackFrames, &StackFrame{
				File:       frame.File,
				Name:       frame.Name,
				LineNumber: frame.LineNumber,
				Package:    frame.Package,
			})
		}

		return c.JSON(http.StatusInternalServerError, &ErrorWrapper{Err: err.Error(), StackTraces: stackFrames})

	}

	return c.String(http.StatusNoContent, ``)
}

func Init(e *echo.Echo, u *usecase.CrashUsecase) {
	handler := CrashHTTPHandler{CrashUsecase: u}
	e.GET(`/crash`, handler.HandleCrash)
}
