package config

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func newServiceConfig() IServiceConfig {
	obj := &ServiceConfigSetup{
		v:              viper.New(),
		lastChangeTime: time.Now(),
	}

	obj.Load()

	return obj
}

type IServiceConfig interface {
	GetMySQLServiceConfig() MySQLServiceConfig
	GetRedisServiceConfig() RedisServiceConfig
}

type ServiceConfigSetup struct {
	v              *viper.Viper
	lastChangeTime time.Time

	ServiceConfig ServiceConfig `mapstructure:"service_config"`
}

func (c *ServiceConfigSetup) Load() {
	c.loadYaml()
}

func (c *ServiceConfigSetup) GetLastChangeTime() time.Time {
	return c.lastChangeTime
}

func (c *ServiceConfigSetup) loadYaml() {
	path, err := filepath.Abs("conf.d")
	if err != nil {
		panic(err)
	}

	c.v.SetConfigName("service.yaml")
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

func (c *ServiceConfigSetup) GetMySQLServiceConfig() MySQLServiceConfig {
	return c.ServiceConfig.MySQLServiceConfig
}
func (c *ServiceConfigSetup) GetRedisServiceConfig() RedisServiceConfig {
	return c.ServiceConfig.RedisServiceConfig
}
