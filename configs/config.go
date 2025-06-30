package configs

import (
	"os"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type Config struct {
	WebServerPort  int    `mapstructure:"WEBSERVER_PORT"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         int    `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBUserPassword string `mapstructure:"DB_USER_PASSWORD"`
	DBDatabaseName string `mapstructure:"DB_NAME"`
	DBDriver       string `mapstructure:"DB_DRIVER"`
	JWTSecret      string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn   int    `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth      *jwtauth.JWTAuth
}

func LoadConfig(path string) (*Config, error) {

	if _, err := os.Stat(path); err != nil {

		err = os.WriteFile(path, []byte(`WEBSERVER_PORT=8000
DB_PORT=5432
DB_USER=
DB_USER_PASSWORD=
DB_NAME=
DB_DRIVER=postgres
JWT_SECRET=secret
JWT_EXPIRESIN=3000
`), 0755)

		if err != nil {
			return nil, err
		}
	}

	viper.SetConfigFile(path)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return &cfg, nil
}
