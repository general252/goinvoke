package example

var _default = NewEngine()

func GetDftEngine() *Engine {
	return _default
}

type Engine struct {
	serviceA ServiceA

	serviceB ServiceB
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) RegisterServiceA(s ServiceA) {
	e.serviceA = s
}

func (e *Engine) GetServiceA() (ServiceA, bool) {
	return e.serviceA, e.serviceA != nil
}

func (e *Engine) RegisterServiceB(s ServiceB) {
	e.serviceB = s
}

func (e *Engine) GetServiceB() (ServiceB, bool) {
	return e.serviceB, e.serviceB != nil
}

type (
	ServiceAHelloFunc func(a string, f func(int)) (arg1 int, arg2 string, arg3 error)

	ServiceAHello2Func func(param *Param) (arg1 string, arg2 error)

	ServiceAHello3Func func(args ...any) (arg1 string, arg2 error)

	ServiceAHello4Func func(args2 any) (arg1 string, arg2 error)

	ServiceAHello5Func func(args2 []int, f func(int, byte)) (arg1 string, arg2 error)

	ServiceAHello6Func func()
)

var _ ServiceA = (*UnimplementedServiceA)(nil)

type UnimplementedServiceA struct {
	ServiceAHelloFunc  ServiceAHelloFunc
	ServiceAHello2Func ServiceAHello2Func
	ServiceAHello3Func ServiceAHello3Func
	ServiceAHello4Func ServiceAHello4Func
	ServiceAHello5Func ServiceAHello5Func
	ServiceAHello6Func ServiceAHello6Func
}

func (u *UnimplementedServiceA) Hello(a string, f func(int)) (arg1 int, arg2 string, arg3 error) {
	if u.ServiceAHelloFunc != nil {
		arg1, arg2, arg3 = u.ServiceAHelloFunc(a, f)
	}

	return arg1, arg2, arg3
}

func (u *UnimplementedServiceA) Hello2(param *Param) (arg1 string, arg2 error) {
	if u.ServiceAHello2Func != nil {
		arg1, arg2 = u.ServiceAHello2Func(param)
	}

	return arg1, arg2
}

func (u *UnimplementedServiceA) Hello3(args ...any) (arg1 string, arg2 error) {
	if u.ServiceAHello3Func != nil {
		arg1, arg2 = u.ServiceAHello3Func(args)
	}

	return arg1, arg2
}

func (u *UnimplementedServiceA) Hello4(args2 any) (arg1 string, arg2 error) {
	if u.ServiceAHello4Func != nil {
		arg1, arg2 = u.ServiceAHello4Func(args2)
	}

	return arg1, arg2
}

func (u *UnimplementedServiceA) Hello5(args2 []int, f func(int, byte)) (arg1 string, arg2 error) {
	if u.ServiceAHello5Func != nil {
		arg1, arg2 = u.ServiceAHello5Func(args2, f)
	}

	return arg1, arg2
}

func (u *UnimplementedServiceA) Hello6() {
	if u.ServiceAHello6Func != nil {
		u.ServiceAHello6Func()
	}

	return
}

type (
	ServiceBVersionFunc func(a string, c int) (arg1 int, arg2 string, arg3 error)

	ServiceBInfoFunc func(param *Param) (arg1 string, arg2 error)
)

var _ ServiceB = (*UnimplementedServiceB)(nil)

type UnimplementedServiceB struct {
	ServiceBVersionFunc ServiceBVersionFunc
	ServiceBInfoFunc    ServiceBInfoFunc
}

func (u *UnimplementedServiceB) Version(a string, c int) (arg1 int, arg2 string, arg3 error) {
	if u.ServiceBVersionFunc != nil {
		arg1, arg2, arg3 = u.ServiceBVersionFunc(a, c)
	}

	return arg1, arg2, arg3
}

func (u *UnimplementedServiceB) Info(param *Param) (arg1 string, arg2 error) {
	if u.ServiceBInfoFunc != nil {
		arg1, arg2 = u.ServiceBInfoFunc(param)
	}

	return arg1, arg2
}
