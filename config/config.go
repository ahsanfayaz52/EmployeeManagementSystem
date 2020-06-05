package config

import (
	"github.com/spf13/viper"
)

// keys for database configuration.
const (
	DbName      = "db.name"
	DbHost      = "db.ip"
	DbPort      = "db.port"
	DbUser      = "db.user"
	DbPass      = "db.pass"
	ServerHost  = "server.host"
	ServerPort  = "server.port"
)

func init() {
	// defaults.
	viper.SetDefault(DbName, "employee_system")
	viper.SetDefault(DbHost, "localhost")
	viper.SetDefault(DbPort, "3306")
	viper.SetDefault(DbUser, "root")

	viper.SetDefault(ServerHost, "127.0.0.1")
	viper.SetDefault(ServerPort, "8080")
}
