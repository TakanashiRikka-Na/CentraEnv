package command

import (
	"os"
	"os/exec"
)

func SetPsEnv(s []string) error {
	cmd := exec.Command("pwsh", "$PROFILE")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(string(output), os.O_RDWR, 0666)
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
