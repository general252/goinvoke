package example

import (
	"runtime"
)

type Param struct {
	Msg  string
	Name string
}

type ServiceA interface {
	Hello(a string, f func(int)) (int, string, error)
	Hello2(param *Param) (string, error)
}

type ServiceB interface {
	Version(a string, c int) (int, string, error)
	Info(param *Param) (string, error)
}

func GetFilename() string {
	_, file, _, _ := runtime.Caller(0)
	return file
}
