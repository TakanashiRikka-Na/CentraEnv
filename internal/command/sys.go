package command

import (
	"github.com/TakanashiRikka-Na/CentraEnv/internal/config"
	"golang.org/x/sys/windows/registry"
	"strings"
)

func SetToSysPath(env string) error {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\Session Manager\Environment`, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()
	paths := append(config.ReadPathEnv(), env)
	join := strings.Join(paths, ",")
	err = key.SetStringValue("Path", join)
	if err != nil {
		return err
	}
	return nil
}
