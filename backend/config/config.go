package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type Config struct {
	App      App      `yaml:"app"`
	Otel     Otel     `yaml:"otel"`
	Database Database `yaml:"database"`
}

func (c *Config) Version() string {
	return Version
}

type App struct {
	Name string `yaml:"name"`
}

type Otel struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (otel Otel) String() string {
	return fmt.Sprintf("%s:%d", otel.Host, otel.Port)
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
