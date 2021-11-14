package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Config *Configuration

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

type DatabaseConfiguration struct {
	Driver       string
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         int
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

type ServerConfiguration struct {
	Port int
	Mode string
}

// Setup initialize configuration
func Setup(log *zap.SugaredLogger, configPath string) *Configuration {
	var configuration *Configuration

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to Unmarshal into struct, %v", err)
	}

	Config = configuration

	return Config
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
