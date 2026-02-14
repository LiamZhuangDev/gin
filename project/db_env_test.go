package project

import (
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

type Config struct {
	Database DBConfig `mapstructure:"database"`
}

func TestDatabaseConfigFromEnv(t *testing.T) {
	// set env vars
	t.Setenv("DATABASE_HOST", "localhost")
	t.Setenv("DATABASE_PORT", "5432")
	t.Setenv("DATABASE_USER", "admin")
	t.Setenv("DATABASE_PASSWORD", "secret")
	t.Setenv("DATABASE_NAME", "mydb")

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	require.Equal(t, "localhost", cfg.Database.Host)
	require.Equal(t, 5432, cfg.Database.Port)
	require.Equal(t, "admin", cfg.Database.User)
	require.Equal(t, "secret", cfg.Database.Password)
	require.Equal(t, "mydb", cfg.Database.Name)

	require.Equal(t, "localhost", viper.GetString("database.host")) // this requires SetEnvKeyReplacer
}

func LoadConfig() (*Config, error) {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv() // read from env vars

	// set default values or use viper.BindEnv to bind env vars to config keys
	// or even better use config file
	viper.SetDefault("database.host", "")
	viper.SetDefault("database.port", 0)
	viper.SetDefault("database.user", "")
	viper.SetDefault("database.password", "")
	viper.SetDefault("database.name", "")

	// viper.BindEnv("database.host", "DATABASE_HOST")
	// viper.BindEnv("database.port", "DATABASE_PORT")
	// viper.BindEnv("database.user", "DATABASE_USER")
	// viper.BindEnv("database.password", "DATABASE_PASSWORD")
	// viper.BindEnv("database.name", "DATABASE_NAME")

	var cfg Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
