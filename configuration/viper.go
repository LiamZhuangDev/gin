// Viper is a popular Go library for application configuration management.
// Viper can read from multiple configuration sources and merges them together into one set of configuration keys and values.
// Viper uses the following precedence for merging, it handles priority/overrides automatically (1 > 2 > 3 ... > 6).
// 1. explicit call to Set
// 2. flags
// 3. environment variables
// 4. config files
// 5. external key/value stores
// 6. defaults
package configuration

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

func SetupViper() {
	// Name of the config file without an extension (Viper will intuit the type
	// from an extension on the actual file)
	viper.SetConfigName("config")

	// Add search paths to find the file
	viper.AddConfigPath("/etc/appname/")
	viper.AddConfigPath("$HOME/.appname")
	viper.AddConfigPath(".")

	// Find and read the config file
	err := viper.ReadInConfig()

	// Handle errors
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

// mapstructure is a Go lib that converts generic maps (like map[string]interface{}) into typed Go structs.
// map / JSON / YAML / env -> Go struct
// Viper uses mapstructure under the hood to translate config keys into struct fields.
type ServerConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
	Mode string `mapstructure:"mode"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBname   string `mapstructure:"dbname"`
}

type JWTConfig struct {
	Secret string        `mapstructure:"secret"`
	Expire time.Duration `mapstructure:"expire"`
}

type Config struct {
	Server  ServerConfig   `mapstructure:"server"`
	Databse DatabaseConfig `mapstructure:"database"`
	JWT     JWTConfig      `mapstructure:"jwt"`
}

func LoadConfig() (*Config, error) {
	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
