package file

import (
	"github.com/asticode/go-toolbox/dir"
	"io/ioutil"
	"os"
	"path"
)

func GetContentBytes(sPath string) []byte {
	// Read file.
	sFileContent, oErr := ioutil.ReadFile(sPath)

	// Error.
	if oErr != nil {
		panic(oErr)
	}

	// Return file content.
	return sFileContent
}

func GetContent(sPath string) string {
	return string(GetContentBytes(sPath))
}

func Exists(sPath string) bool {
	if _, oErr := os.Stat(sPath); oErr == nil {
		return true
	} else {
		return false
	}
}

func DirPath(Path string) string {
	return path.Dir(Path)
}

func Create(sPath string) (file *os.File, err error) {
	// Make sure directory exists
	oErr := dir.Create(DirPath(sPath))
	if oErr != nil {
		return nil, oErr
	}

	// Create file
	oFile, oErr := os.OpenFile(sPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	oFile.Close()
	return oFile, oErr
}

func Append(sPath string, sMessage string) error {
	// Open file
	oFile, oErr := os.OpenFile(sPath, os.O_APPEND|os.O_WRONLY, 0600)
	if oErr != nil {
		oFile.Close()
		return oErr
	}

	// Write in file
	_, oErr = oFile.WriteString(sMessage)
	if oErr != nil {
		oFile.Close()
		return oErr
	}
	oFile.Close()

	return nil
}
