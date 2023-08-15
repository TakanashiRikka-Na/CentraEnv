package command

import (
	"fmt"
	"os"
)

func SetToEnv(s []string) error {
	shell := os.Getenv("SHELL")
	file, err := os.OpenFile(fmt.Sprintf("~/.%src", shell), os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, val := range s {
		_, err = file.WriteString(val + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
