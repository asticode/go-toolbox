package environment

import (
	"os"
	"path/filepath"
)

func GetCurrentDir() string {
	sDirectory, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	} else {
		return sDirectory
	}
}
