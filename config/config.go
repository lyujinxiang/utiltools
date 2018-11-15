package config

import (
	"os"
	"os/exec"
	"path/filepath"
)

func GetWorkPath() string {
	if file, err := exec.LookPath(os.Args[0]); err == nil {
		return filepath.Dir(file) + "/"
	}
	return "./"
}
