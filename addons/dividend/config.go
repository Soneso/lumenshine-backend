package main

import (
	"github.com/Soneso/lumenshine-backend/helpers"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//Config for the app
type Config struct {
	AppName string

	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int64
	DBName     string

	CoreDBUser     string
	CoreDBPassword string
	CoreDBHost     string
	CoreDBPort     int64
	CoreDBName     string

	Port            int
	SQLMigrationDir string
	ApplicationDir  string
	IsDevSystem     bool
}

//Cnf holds the application configuration
var Cnf *Config

//ReadConfig reads configuration from command
func ReadConfig(cmd *cobra.Command) error {
	Cnf = new(Config)

	if err := helpers.ReadConfig(cmd, "db-config", "db-config-local", Cnf); err != nil {
		return err
	}

	if viper.GetString("ApplicationDirectory") != "" {
		Cnf.ApplicationDir = viper.GetString("ApplicationDirectory")
	} else {
		Cnf.ApplicationDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	}

	if Cnf.IsDevSystem {
		Cnf.SQLMigrationDir = filepath.Join(Cnf.ApplicationDir, "tmp", "migrations")
	} else {
		Cnf.SQLMigrationDir = filepath.Join(Cnf.ApplicationDir, "db", "migrations")
	}
	helpers.CreateDirIfNotExists(Cnf.SQLMigrationDir)

	return nil
}
