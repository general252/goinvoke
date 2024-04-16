package example

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	beginInvoke "github.com/general252/goinvoke"
)

func ExampleGenerate() {
	var (
		filename    = GetFilename()
		outFilename = fmt.Sprintf("%vEngine.go", strings.TrimSuffix(filename, filepath.Ext(filename)))
	)

	if data, err := beginInvoke.Generate(filename, nil); err == nil {
		err = os.WriteFile(outFilename, data, 0600)
		log.Println("WriteFile", err)
	} else {
		log.Println(err)
	}

	// output:
	//
}
