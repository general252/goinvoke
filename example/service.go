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
	Hello3(args ...any) (string, error)
	Hello4(args2 interface{}) (string, error)
	Hello5(args2 []int, f func(int, byte)) (string, error)
	Hello6()
}

type ServiceB interface {
	Version(a string, c int) (int, string, error)
	Info(param *Param) (string, error)
}

func GetFilename() string {
	_, file, _, _ := runtime.Caller(0)
	return file
}
