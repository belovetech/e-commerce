package config

import "github.com/spf13/viper"

type Config struct {
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
}

func LoadConfig() (Config, error) {
	var config Config
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	return config, err
}
