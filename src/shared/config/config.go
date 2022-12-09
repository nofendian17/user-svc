package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
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
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resources")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Couldn't load configuration, cannot start. Terminating. Error: " + err.Error())
	}
	log.Println("Config loaded successfully...")
	log.Println("Getting environment variables...")
	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}

	appConfig := DefaultConfig{}
	err = viper.Unmarshal(&appConfig)
	if err != nil {
		panic(err)
	}
	return &appConfig
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(res) == 0 {
		panic("Mandatory env variable not found:" + env)
	}
	return res
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
