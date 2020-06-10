package config

import (
	"github.com/spf13/viper"
)

// keys for database configuration.
const (
	DbName     = "db.name"
	DbHost     = "db.ip"
	DbPort     = "db.port"
	DbUser     = "db.user"
	DbPass     = "db.pass"
	ServerHost = "server.host"
	ServerPort = "server.port"
	Filepath   = "filepath"
)

func init() {
	// defaults.
	_ = viper.BindEnv(DbPort, "DB_PORT")
	_ = viper.BindEnv(DbHost, "DB_HOST")
	_ = viper.BindEnv(Filepath, "FILE_PATH")

	viper.SetDefault(DbName, "employee_system")
	viper.SetDefault(DbHost, "localhost")
	viper.SetDefault(DbPort, "27017")
	viper.SetDefault(DbUser, "root")
	viper.SetDefault(Filepath, "../../data/records.csv")

	viper.SetDefault(ServerHost, "127.0.0.1")
	viper.SetDefault(ServerPort, "8080")
}
