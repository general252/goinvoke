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
	ServiceAHelloFunc func(a string, f func(arg1 int)) (arg1 int, arg2 string, arg3 error)

	ServiceAHello2Func func(param *Param) (arg1 string, arg2 error)
)

var _ ServiceA = (*UnimplementedServiceA)(nil)

type UnimplementedServiceA struct {
	ServiceAHelloFunc  ServiceAHelloFunc
	ServiceAHello2Func ServiceAHello2Func
}

func (u *UnimplementedServiceA) Hello(a string, f func(arg1 int)) (arg1 int, arg2 string, arg3 error) {
	if u.ServiceAHelloFunc != nil {
		return u.ServiceAHelloFunc(a, f)
	}

	return arg1, arg2, arg3
}

func (u *UnimplementedServiceA) Hello2(param *Param) (arg1 string, arg2 error) {
	if u.ServiceAHello2Func != nil {
		return u.ServiceAHello2Func(param)
	}

	return arg1, arg2
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
		return u.ServiceBVersionFunc(a, c)
	}

	return arg1, arg2, arg3
}

func (u *UnimplementedServiceB) Info(param *Param) (arg1 string, arg2 error) {
	if u.ServiceBInfoFunc != nil {
		return u.ServiceBInfoFunc(param)
	}

	return arg1, arg2
}
