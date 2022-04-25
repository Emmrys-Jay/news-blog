package util

import "github.com/spf13/viper"

// Config stores all configurations of the application
// The values are read by viper from a config file or environmental variables
type Config struct {
	DBDriver     string `mapstructure:"DB_DRIVER"`
	DBSource     string `mapstructure:"DB_SOURCE"`
	ServerAdress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads config
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig() // start reading config values
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return

}
