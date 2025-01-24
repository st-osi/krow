package app

import "time"

type Config struct {
	Env     string `yaml:"Env"`
	EnvPath string `yaml:"EnvPath"`
}

type App struct {
	Name        string
	ProxyUrl    string
	Version     string
	Config      *Config
	RequestTime time.Duration
}

const APP_NAME = "krow"

func NewApp(proxyUrl string, version string, config *Config) *App {
	return &App{
		Name:     APP_NAME,
		ProxyUrl: proxyUrl,
		Version:  version,
		Config:   config,
	}
}

func (a *App) Terminate() {
	a = nil
}
