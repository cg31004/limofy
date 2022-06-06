package config

type ServiceConfig struct {
	MySQLServiceConfig MySQLServiceConfig `mapstructure:"mysql_service"`
	RedisServiceConfig RedisServiceConfig `mapstructure:"redis_service"`
}

type MySQLServiceConfig struct {
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"user_name"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"limofy"`
}
type RedisServiceConfig struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	TTLSec   int    `mapstructure:"ttl_sec"`
}
