package util

import "github.com/spf13/viper"

// type Config struct {
// 	DBDriver      string `env:"DB_DRIVER,required"`
// 	DBSource      string `env:"DB_SOURCE,required"`
// 	ServerAddress string `env:"SERVER_ADDRESS,required"`
// }

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	var config Config
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)
	return config, err
}
