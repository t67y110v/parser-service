package config

import "github.com/spf13/viper"

type Config struct {
	Port      string `mapstructure:"PORT"`
	PG_DBHost string `mapstructure:"PG_DB_HOST"`
	PG_DBUser string `mapstructure:"PG_DB_USER"`
	PG_DBPass string `mapstructure:"PG_DB_PASS"`
	PG_DBName string `mapstructure:"PG_DB_NAME"`
	PG_DBPort string `mapstructure:"PG_DB_PORT"`
	MG_DBUser string `mapstructure:"MG_DB_USER"`
	MG_DBHost string `mapstructure:"MG_DB_HOST"`
	MG_DBPort string `mapstructure:"MG_DB_PORT"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
