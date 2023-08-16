package command

import (
	"github.com/TakanashiRikka-Na/CentraEnv/internal/config"
	"golang.org/x/sys/windows/registry"
	"strings"
)

func SetToUserPath(env string) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, "Environment", registry.QUERY_VALUE)
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
