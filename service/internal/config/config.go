package config

type EnvType = string

const (
	EenTypeLocal EnvType = "local"
	EenTypeProd  EnvType = "prod"
)

func NewAppConfig() IAppConfig {
	return newAppConfig()
}

func NewServiceConfig() IServiceConfig {
	return newServiceConfig()
}
