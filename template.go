package beginInvoke

import (
	"log"
	"text/template"
)

func newTemplate() (*template.Template, error) {
	tmpl, err := template.New("test").Parse(templateText)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return tmpl, nil
}

var templateText = `

package {{.PackageName}}

var _default = NewEngine()

func GetDftEngine() *Engine {
	return _default
}


type Engine struct {
{{ range $service := .Services }}
    {{$service.LowName}} {{$service.Name}}
{{end}}
}

func NewEngine() *Engine {
	return &Engine{}
}

{{ range $service := .Services }}
    func (e *Engine) Register{{$service.Name}}(s {{$service.Name}}) {
        e.{{$service.LowName}} = s
    }

    func (e *Engine) Get{{$service.Name}}() ({{$service.Name}}, bool) {
    	return e.{{$service.LowName}}, e.{{$service.LowName}} != nil
    }
{{end}}


{{ range $service :=.Services }}

type (
	{{ range $method := $service.Methods }} {{ $service.Name }}{{ $method.Name}}Func func( {{ $method.InString }} ) ( {{ $method.OutString }} )
{{end}}
)


var _ {{ $service.Name }} = (*Unimplemented{{ $service.Name }})(nil)

type Unimplemented{{ $service.Name }} struct {
    {{ range $method := .Methods }}
        {{ $service.Name }}{{ $method.Name }}Func   {{ $service.Name }}{{$method.Name}}Func   {{end}}
}

{{ range $method :=.Methods }}
func (u *Unimplemented{{ $service.Name }}) {{$method.Name}}({{ $method.InString }} ) ({{ $method.OutString }}) {
	if u.{{ $service.Name }}{{$method.Name}}Func != nil {
		{{ if eq $method.OutArgs "" }} {{else}} {{ $method.OutArgs }} = {{end}} u.{{ $service.Name }}{{$method.Name}}Func( {{ $method.InArgs }} )
	}

	return {{ $method.OutArgs }}
}
{{end}}

{{end}}


`
