package svc

import (
	"os"

	"github.com/st-osi/krow/core/app"
	"github.com/st-osi/krow/core/utils"
)

func LoadConfig() (*app.Config, error) {
	config := &app.Config{}
	configPath := utils.Pwd() + "/" + "config.yaml"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return config, err
	}
	err := utils.LoadWithYaml(configPath, config)
	if err != nil {
		return getDefaultConfig(), err
	}

	return config, nil
}

func getDefaultConfig() *app.Config {
	return &app.Config{
		Env:     "",
		EnvPath: "",
	}
}
