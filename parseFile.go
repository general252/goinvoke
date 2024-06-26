package beginInvoke

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// ParseServiceFile 解析文件中的服务
func ParseServiceFile(filename string) (*ServiceFile, error) {

	file, err := parser.ParseFile(token.NewFileSet(), filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var result = ServiceFile{
		PackageName: file.Name.Name,
		Services:    nil,
	}

	for _, decl := range file.Decls {
		// 检查是否为通用声明
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			// 检查是否为类型声明
			if genDecl.Tok != token.TYPE {
				continue
			}

			// 遍历类型声明的规范
			for _, spec := range genDecl.Specs {
				// 检查是否为接口类型声明
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				// 检查是否为接口类型
				if interfaceType, ok := typeSpec.Type.(*ast.InterfaceType); ok {
					// 打印接口名称
					// fmt.Println("接口名称:", typeSpec.Name.Name)
					if !strings.HasPrefix(typeSpec.Name.Name, "Service") {
						continue
					}

					var objectService = ObjectService{
						Name:    typeSpec.Name.Name,
						Methods: nil,
					}

					// 遍历接口方法
					for _, method := range interfaceType.Methods.List {
						var arg = 1
						// 检查是否为函数类型
						if funcType, ok := method.Type.(*ast.FuncType); ok {
							// 打印方法名称
							objectMethod := convertMethod(method.Names[0].Name, funcType, &arg)

							objectService.Methods = append(objectService.Methods, objectMethod)
						}
					}

					result.Services = append(result.Services, &objectService)
				}
			}
		}
	}

	result.build()
	return &result, nil
}

type ServiceFile struct {
	PackageName string
	Services    []*ObjectService
}

func (s *ServiceFile) build() {
	for _, service := range s.Services {
		service.build()
	}
}

type ObjectService struct {
	Name    string
	LowName string
	Methods []*ObjectMethod
}

func (s *ObjectService) build() {
	s.LowName = strings.ToLower(s.Name[:1]) + s.Name[1:]
	for _, method := range s.Methods {
		method.build()
	}
}

type ObjectMethod struct {
	Name     string
	InTypes  []ObjectInOut
	OutTypes []ObjectInOut

	InString  string
	OutString string
	InArgs    string
	OutArgs   string
}

func (o *ObjectMethod) String() string {
	return fmt.Sprintf("func(%v) (%v)", o.InString, o.OutString)
}

func (o *ObjectMethod) build() *ObjectMethod {
	{
		var parts []string
		for _, v := range o.InTypes {
			parts = append(parts, fmt.Sprintf("%v %v", v.Name, v.Type))
		}

		o.InString = strings.Join(parts, ", ")
	}

	{
		var parts []string
		for _, v := range o.OutTypes {
			parts = append(parts, fmt.Sprintf("%v %v", v.Name, v.Type))
		}

		o.OutString = strings.Join(parts, ", ")
	}

	{
		var parts []string
		for _, v := range o.InTypes {
			parts = append(parts, fmt.Sprintf("%v", v.Name))
		}

		o.InArgs = strings.Join(parts, ", ")
	}
	{
		var parts []string
		for _, v := range o.OutTypes {
			parts = append(parts, fmt.Sprintf("%v", v.Name))
		}

		o.OutArgs = strings.Join(parts, ", ")
	}

	return o
}

type ObjectInOut struct {
	Name string
	Type string
}

func convertMethod(methodName string, funcType *ast.FuncType, arg *int) *ObjectMethod {
	// fmt.Println("方法名称:", method.Names[0].Name)
	objectMethod := ObjectMethod{
		Name:     methodName,
		InTypes:  nil,
		OutTypes: nil,
	}

	// 遍历函数参数
	if funcType.Params != nil {
		for _, field := range funcType.Params.List {
			// 打印参数名称和类型
			var typeName = fieldTypeString(field.Type)
			if len(field.Names) > 0 {
				for _, name := range field.Names {
					// fmt.Printf("参数名称: %s, 类型: %s\n", name.Name, typeName)
					objectMethod.InTypes = append(objectMethod.InTypes, ObjectInOut{
						Name: name.Name,
						Type: typeName,
					})
				}
			} else {
				objectMethod.InTypes = append(objectMethod.InTypes, ObjectInOut{
					Name: fmt.Sprintf("arg%v", *arg),
					Type: typeName,
				})

				*arg++
			}
		}
	}

	if funcType.Results != nil {
		for _, field := range funcType.Results.List {
			// 打印参数名称和类型
			var typeName = fieldTypeString(field.Type)
			if len(field.Names) > 0 {
				for _, name := range field.Names {
					// fmt.Printf("参数名称: %s, 类型: %s\n", name.Name, typeName)
					objectMethod.OutTypes = append(objectMethod.OutTypes, ObjectInOut{
						Name: name.Name,
						Type: typeName,
					})
				}
			} else {
				objectMethod.OutTypes = append(objectMethod.OutTypes, ObjectInOut{
					Name: fmt.Sprintf("arg%v", *arg),
					Type: typeName,
				})

				*arg++
			}
		}
	}

	return &objectMethod
}

// 获取类型名称
func fieldTypeString(fieldType ast.Expr) string {
	switch t := fieldType.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + fieldTypeString(t.X)
	case *ast.SelectorExpr:
		return fieldTypeString(t.X) + "." + t.Sel.Name
	case *ast.ArrayType:
		return "[]" + fieldTypeString(t.Elt)
	case *ast.MapType:
		return "map[" + fieldTypeString(t.Key) + "]" + fieldTypeString(t.Value)
	case *ast.ChanType:
		dir := ""
		switch t.Dir {
		case ast.SEND:
			dir = "chan<-"
		case ast.RECV:
			dir = "<-chan"
		default:
			dir = "chan"
		}
		return dir + " " + fieldTypeString(t.Value)
	case *ast.Ellipsis:
		return "..." + fieldTypeString(t.Elt)
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.FuncType:
		params := fieldListString(t.Params)
		results := fieldListString(t.Results)
		return "func" + params + results
	case *ast.StructType:
		return "struct{}"
	case *ast.BasicLit:
		return t.Value
	case *ast.ParenExpr:
		return "(" + fieldTypeString(t.X) + ")"
	case *ast.FuncLit:
		return "func" + fieldTypeString(t.Type)
	case *ast.SliceExpr:
		return fieldTypeString(t.X) + "[" + fieldTypeString(t.Low) + ":" + fieldTypeString(t.High) + "]"
	case *ast.IndexExpr:
		return fieldTypeString(t.X) + "[" + fieldTypeString(t.Index) + "]"
	case *ast.TypeAssertExpr:
		return fieldTypeString(t.X) + "." + fieldTypeString(t.Type)
	case *ast.CallExpr:
		funcName := fieldTypeString(t.Fun)
		args := make([]string, len(t.Args))
		for i, arg := range t.Args {
			args[i] = fieldTypeString(arg)
		}
		return funcName + "(" + joinStrings(args, ", ") + ")"
	default:
		return "<unknown>" // reflect.TypeOf(fieldType).String()
	}
}

func fieldListString(fields *ast.FieldList) string {
	if fields == nil {
		return ""
	}

	result := "("
	for i, field := range fields.List {
		if i > 0 {
			result += ", "
		}
		result += fieldTypeString(field.Type)
	}
	result += ")"
	return result
}

func joinStrings(strs []string, separator string) string {
	if len(strs) == 0 {
		return ""
	}
	return strs[0] + joinStrings(strs[1:], separator)
}
