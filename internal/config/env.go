package config

import (
	"os"
	"regexp"
	"strings"
)

func ReadEnv() []string {
	return os.Environ()
}

func ReadPathEnv() (path []string) {
	environ := os.Environ()
	re := regexp.MustCompile(`\bPath\b`)
	for _, s := range environ {
		if re.MatchString(s) {
			path = strings.Split(s[5:], ",")
		}
	}
	return
}
