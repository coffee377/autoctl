package configs

import "time"

type Config struct {
	Redis RedisConfig `mapstructure:"redis"`
	Token TokenConfig `mapstructure:"token"`
}

type RedisConfig struct {
	Address  string `mapstructure:"address"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type TokenConfig struct {
	Account      string        `mapstructure:"account"`
	Expire       time.Duration `mapstructure:"expire"`
	SessionFixed bool          `mapstructure:"session-fixed"`
}
