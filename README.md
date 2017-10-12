# gocrashandburn
Test the errors format using library https://github.com/go-errors/errors

Expected error result will be :

```json
{
    "error": "Something happened in repository",
    "stack_traces": [
        {
            "file": "/Users/arie/Documents/Projects/Go/src/github.com/arielizuardi/gocrashandburn/crash/repository/repository.go",
            "name": "(*CrashRepository).IgniteErrors",
            "line": 11,
            "pkg": "github.com/arielizuardi/gocrashandburn/crash/repository"
        },
        {
            "file": "/Users/arie/Documents/Projects/Go/src/github.com/arielizuardi/gocrashandburn/crash/usecase/usecase.go",
            "name": "(*CrashUsecase).IgniteErrors",
            "line": 12,
            "pkg": "github.com/arielizuardi/gocrashandburn/crash/usecase"
        },
        {
            "file": "/Users/arie/Documents/Projects/Go/src/github.com/arielizuardi/gocrashandburn/crash/http/http.go",
            "name": "(*CrashHTTPHandler).HandleCrash",
            "line": 34,
            "pkg": "github.com/arielizuardi/gocrashandburn/crash/http"
        },
        {
            "file": "/Users/arie/Documents/Projects/Go/src/github.com/arielizuardi/gocrashandburn/crash/http/http.go",
            "name": "(*CrashHTTPHandler).HandleCrash-fm",
            "line": 97,
            "pkg": "github.com/arielizuardi/gocrashandburn/crash/http"
        },
        {
            "file": "/Users/arie/Documents/Projects/Go/src/github.com/arielizuardi/gocrashandburn/vendor/github.com/labstack/echo/echo.go",
            "name": "(*Echo).Add.func1",
            "line": 473,
            "pkg": "github.com/arielizuardi/gocrashandburn/vendor/github.com/labstack/echo"
        },
        {
            "file": "/Users/arie/Documents/Projects/Go/src/github.com/arielizuardi/gocrashandburn/vendor/github.com/labstack/echo/echo.go",
            "name": "(*Echo).ServeHTTP.func1",
            "line": 570,
            "pkg": "github.com/arielizuardi/gocrashandburn/vendor/github.com/labstack/echo"
        },
        {
            "file": "/Users/arie/Documents/Projects/Go/src/github.com/arielizuardi/gocrashandburn/vendor/github.com/labstack/echo/echo.go",
            "name": "(*Echo).ServeHTTP",
            "line": 579,
            "pkg": "github.com/arielizuardi/gocrashandburn/vendor/github.com/labstack/echo"
        },
        {
            "file": "/usr/local/Cellar/go/1.9/libexec/src/net/http/server.go",
            "name": "serverHandler.ServeHTTP",
            "line": 2619,
            "pkg": "net/http"
        },
        {
            "file": "/usr/local/Cellar/go/1.9/libexec/src/net/http/server.go",
            "name": "(*conn).serve",
            "line": 1801,
            "pkg": "net/http"
        },
        {
            "file": "/usr/local/Cellar/go/1.9/libexec/src/runtime/asm_amd64.s",
            "name": "goexit",
            "line": 2337,
            "pkg": "runtime"
        }
    ]
}
```
