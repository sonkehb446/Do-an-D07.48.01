package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/configor"
)

const (
	EnvProduction = "production"
	EnvTesting    = "testing"
	EnvLocal      = "local"
)

type server struct {
	Port int
}

type database struct {
	Name               string
	Host               string
	Port               int
	User               string
	Password           string
	MaxIdleConnections int
	MaxOpenConnections int
}

type gmail struct {
	User     string
	Password string
}

type s3 struct {
	Bucket          string
	Region          string
	AccessKey       string
	SecretAccessKey string
	Token           string
}

type envValue struct {
	Server   server
	Database database
	Gmail    gmail
	S3       s3
}

type env struct {
	Production envValue
	Testing    envValue
	Local      envValue
}

type config struct {
	Env    env
	Secret string
}

var configValue = &config{}

func Loads(filePath string) *config {
	if err := configor.Load(configValue, filePath); err != nil {
		fmt.Println("Configuration Load Failed")
		panic(err)
	}

	return configValue
}

func GetEnvValue() *envValue {
	if configValue == nil {
		panic("Must run Loads() first")
	}
	env := GetEnv()
	switch env {
	case EnvProduction:
		return &configValue.Env.Production
	case EnvTesting:
		return &configValue.Env.Testing
	case EnvLocal:
		return &configValue.Env.Local
	default:
		return nil
	}
}

func SetEnv(env string) {
	if env != EnvProduction && env != EnvTesting && env != EnvLocal {
		panic("Invalid env")
	}
	if err := os.Setenv("env", env); err != nil {
		panic("Set environment error")
	}
}

func GetEnv() string {
	env := os.Getenv("env")
	if env == "" {
		panic("Environment not set")
	}
	return env
}
