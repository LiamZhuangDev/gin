package configuration_test

import (
	"testing"

	"github.com/LiamZhuangDev/gin/configuration"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestReadConfig(t *testing.T) {
	configuration.SetupViper()

	port := viper.GetString("server.port")
	dbHost := viper.GetString("database.host")

	require.Equal(t, "8080", port)
	require.Equal(t, "localhost", dbHost)
}

func TestReadConfigViaStruct(t *testing.T) {
	configuration.SetupViper()

	config, err := configuration.LoadConfig()
	require.NoError(t, err)

	port := config.Server.Port
	dbhost := config.Databse.Host

	require.Equal(t, "8080", port)
	require.Equal(t, "localhost", dbhost)
}
