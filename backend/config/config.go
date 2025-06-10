package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type Config struct {
	Database Database `yaml:"database"`
}

type Database struct {
	URL      string `yaml:"url"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	SSL      string `yaml:"ssl"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (db Database) String() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", db.Username, db.Password, db.URL, db.Port, db.Name, db.SSL)
}

func Get() *Config {
	c := &Config{}
	viper.SetConfigType("yaml")

	if env, ok := os.LookupEnv("GO_ENV"); ok {
		getEnv(env)
	} else {
		getDev()
	}

	if err := viper.ReadRemoteConfig(); err != nil {
		panic(err)
	}

	viper.Unmarshal(c)

	return c
}

func getEnv(env string) {
	if env == "PROD" {
		viper.AddRemoteProvider("consul", "payd-consul:8500", "rms-api/prod")
	} else {
		getDev()
	}
}

func getDev() {
	viper.AddRemoteProvider("consul", "localhost:9090", "rms-api/dev")
}
