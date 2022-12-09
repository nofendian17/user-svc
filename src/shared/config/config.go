package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

type DefaultConfig struct {
	Apps     *AppsConfig     `yaml:"apps"`
	Database *DatabaseConfig `yaml:"database"`
	Cache    *CacheConfig    `yaml:"cache"`
	Jwt      *JWTConfig      `yaml:"jwt"`
}

type AppsConfig struct {
	Name     string `yaml:"name"`
	Version  string `yaml:"version"`
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	GRPCPort int    `yaml:"grpcPort"`
	Debug    bool   `yaml:"debug"`
}

type DatabaseConfig struct {
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Name        string `yaml:"name"`
	Schema      string `yaml:"schema"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	MaxIdleConn int    `yaml:"maxIdleConn"`
	MaxOpenConn int    `yaml:"maxOpenConn"`
}

type CacheConfig struct {
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	Db         int    `yaml:"db"`
	PoolSize   int    `yaml:"poolSize"`
	DefaultTTL int    `yaml:"defaultTTL"`
}

type JWTConfig struct {
	AccessSecretKey         string `yaml:"accessSecretKey"`
	AccessExpirationMinute  int64  `yaml:"accessExpirationMinute"`
	RefreshExpirationMinute int64  `yaml:"refreshExpirationMinute"`
	RefreshSecretKey        string `yaml:"refreshSecretKey"`
	Issuer                  string `yaml:"issuer"`
}

func InitConfig() *DefaultConfig {
	return loadFile()
}

func loadFile() *DefaultConfig {
	file, err := os.Open("./resources/config.yaml")
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	appConfig := DefaultConfig{}
	err = yaml.Unmarshal(b, &appConfig)
	if err != nil {
		panic(err)
	}
	return &appConfig
}

func (c *AppsConfig) AppPort() string {
	return fmt.Sprintf(":%v", c.Port)
}

func (c *AppsConfig) GrpcPort() string {
	return fmt.Sprintf(":%v", c.GRPCPort)
}

func (c *DefaultConfig) IsDebugMode() bool {
	return c.Apps.Debug
}
