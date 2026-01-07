package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App           AppConfig      `mapstructure:"app"`
	Web           WebConfig      `mapstructure:"web"`
	Database      DatabaseConfig `mapstructure:"database"`
	Kafka         KafkaConfig    `mapstructure:"kafka"`
	CaptchaSecret string         `mapstructure:"captchaSecret"`
	JWTSecret     string         `mapstructure:"jwt_secret"`
	Midtrans      MidtransConfig `mapstructure:"midtrans"`
}
type MidtransConfig struct {
	ServerKey   string `mapstructure:"server_key"`
	ClientKey   string `mapstructure:"client_key"`
	Environment string `mapstructure:"environment"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
}

type WebConfig struct {
	BaseUrl     string `mapstructure:"baseUrl"`
	BaseUrls    string `mapstructure:"baseUrls"`
	Prefork     bool   `mapstructure:"prefork"`
	Port        string `mapstructure:"port"`
	AppPassword string `mapstructure:"app_password"`
}

type DatabaseConfig struct {
	Username string     `mapstructure:"username"`
	Password string     `mapstructure:"password"`
	Host     string     `mapstructure:"host"`
	Port     int        `mapstructure:"port"`
	Name     string     `mapstructure:"name"`
	Pool     PoolConfig `mapstructure:"pool"`
}

type PoolConfig struct {
	Idle     int `mapstructure:"idle"`
	Max      int `mapstructure:"max"`
	Lifetime int `mapstructure:"lifetime"`
}

type KafkaConfig struct {
	Bootstrap KafkaBootstrap `mapstructure:"bootstrap"`
	Group     KafkaGroup     `mapstructure:"group"`
	Auto      KafkaAuto      `mapstructure:"auto"`
}

type KafkaBootstrap struct {
	Servers string `mapstructure:"servers"`
}

type KafkaGroup struct {
	ID string `mapstructure:"id"`
}

type KafkaAuto struct {
	Offset KafkaOffset `mapstructure:"offset"`
}

type KafkaOffset struct {
	Reset string `mapstructure:"reset"`
}

func (db *DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db.Username, db.Password, db.Host, db.Port, db.Name)
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile("config.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	// Set defaults
	viper.SetDefault("jwt_secret", "secret")
	viper.SetDefault("web.port", "0.0.0.0:3000") // Default port

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Warning: Error reading config file, using defaults/env: %s\n", err)
	}

	viper.AutomaticEnv()
	// Replace dot in config path with underscore/env replacement logic if needed
	// Viper automaticEnv usually allows APP_NAME for app.name with SetEnvKeyReplacer
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %v", err)
	}

	return &config, nil
}
