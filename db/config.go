package db

import (
	"github.com/spf13/viper"
)

var (
	filePath = "C:/Users/Lenovo/go/src/booking_asm/config/"
	fileName = "config"
	fileType = "yaml"
)

type (
	Config struct {
		Postgres Postgres `mapstructure:"postgres"`
	}
	Postgres struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DbName   string `mapstructure:"dbname"`
	}
)

func AutoBindConfig() (Config, error) {
	vp := viper.New()
	vp.AddConfigPath(filePath)
	vp.SetConfigName(fileName)
	vp.SetConfigType(fileType)

	vp.AutomaticEnv()
	if err := vp.ReadInConfig(); err != nil {
		return Config{}, err
	}
	conf := Config{}
	if err := vp.Unmarshal(&conf); err != nil {
		return Config{}, err
	}
	return conf, nil
}
