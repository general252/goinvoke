package beginInvoke

import (
	"bytes"
	"os"
	"text/template"

	"github.com/fsgo/go_fmt/gofmtapi"
)

// Generate 解析文件中的以Service开头的接口
// serviceEngine.go
// fmt.Sprintf("%vEngine.go", strings.TrimSuffix(filename, filepath.Ext(filename)))
func Generate(filename string, tmpl *template.Template) ([]byte, error) {
	result, err := ParseServiceFile(filename)
	if err != nil {
		return nil, err
	}

	if tmpl == nil {
		if tmp, err := newTemplate(); err != nil {
			return nil, err
		} else {
			tmpl = tmp
		}
	}

	var buff bytes.Buffer

	if err = tmpl.Execute(&buff, result); err != nil {
		return nil, err
	}

	// 创建一个临时文件, 用于格式化代码
	if fp, err := os.CreateTemp("", "cross-invoke-"); err != nil {
		return nil, err
	} else {
		tmpFilename := fp.Name()

		if _, err = fp.Write(buff.Bytes()); err != nil {
			_ = fp.Close()
			_ = os.Remove(tmpFilename)
			return nil, err
		}
		_ = fp.Close()

		if err = FormatFiles([]string{tmpFilename}); err != nil {
			return nil, err
		}

		data, err := os.ReadFile(tmpFilename)
		if err != nil {
			return nil, err
		}

		_ = os.Remove(tmpFilename)
		return data, nil
	}
}

// FormatFiles 格式化go文件
func FormatFiles(files []string) error {
	gf := gofmtapi.NewFormatter()
	opt := gofmtapi.NewOptions()

	gf.PrintResult = func(fileName string, change bool, err error) {}
	opt.MergeImports = false
	opt.Files = files

	if err := gf.Execute(opt); err != nil {
		return err
	}

	return nil
}
