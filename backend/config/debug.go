package config

import (
	"os"
)

func IsDebug() bool {
	if os.Getenv("Mode") == "debug" {
		return true
	} else {
		return false
	}
}
