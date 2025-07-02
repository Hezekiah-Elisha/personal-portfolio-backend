package util

import "github.com/spf13/viper"

type Config struct {
	DatabaseDriver   string `env:"DB_DRIVER" envDefault:"mysql"`
	DatabaseUser     string `env:"DB_USER" envDefault:"hezekiahdev"`
	DatabasePassword string `env:"DB_PASSWORD" envDefault:"qwerty"`
	DatabaseHost     string `env:"DB_HOST" envDefault:"localhost"`
	DatabasePort     string `env:"DB_PORT" envDefault:"3306"`
	DatabaseName     string `env:"DB_NAME" envDefault:"portfolio_v1"`
	DatabaseURL      string `env:"DATABASE_URL" envDefault:"mysql://hezekiahdev:qwerty@localhost/portfolio_v1"`
	Port             string `env:"PORT" envDefault:"8080"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
