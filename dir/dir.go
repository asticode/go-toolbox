package dir

import (
	"os"
)

func Create(sPath string) error {
	return os.MkdirAll(sPath, 0755)
}
