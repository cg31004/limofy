package config

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func newAppConfig() IAppConfig {
	obj := &AppConfigSetup{
		v:              viper.New(),
		lastChangeTime: time.Now(),
	}

	obj.Load()

	return obj
}

type IAppConfig interface {
	GetAppLogConfig() LogConfig
	GetServerConfig() ServerConfig
	GetGinConfig() GinConfig
	GetMySQLConfig() MySQLConfig
	GetLocalCacheConfig() LocalCacheConfig
	GetSaltConfig() SaltConfig
}

type AppConfigSetup struct {
	v              *viper.Viper
	lastChangeTime time.Time

	AppConfig AppConfig `mapstructure:"app_config"`
}

func (c *AppConfigSetup) Load() {
	c.loadYaml()
}

func (c *AppConfigSetup) GetLastChangeTime() time.Time {
	return c.lastChangeTime
}

func (c *AppConfigSetup) loadYaml() {
	path, err := filepath.Abs("conf.d")
	if err != nil {
		panic(err)
	}

	c.v.SetConfigName("app.yaml")
	c.v.SetConfigType("yaml")
	fmt.Println(path)
	c.v.AddConfigPath(path)

	if err := c.v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := c.v.Unmarshal(c); err != nil {
		panic(err)
	}

	c.v.OnConfigChange(func(in fsnotify.Event) {
		if err := c.v.Unmarshal(c); err != nil {
			panic(err)
		}
		c.lastChangeTime = time.Now()
	})

	c.v.WatchConfig()
}

func (c *AppConfigSetup) GetAppLogConfig() LogConfig {
	return c.AppConfig.LogConfig
}
func (c *AppConfigSetup) GetServerConfig() ServerConfig {
	return c.AppConfig.ServerConfig
}
func (c *AppConfigSetup) GetGinConfig() GinConfig {
	return c.AppConfig.GinConfig
}
func (c *AppConfigSetup) GetMySQLConfig() MySQLConfig {
	return c.AppConfig.MySQLConfig
}
func (c *AppConfigSetup) GetLocalCacheConfig() LocalCacheConfig {
	return c.AppConfig.LocalCacheConfig
}
func (c *AppConfigSetup) GetSaltConfig() SaltConfig {
	return c.AppConfig.SaltConfig
}
