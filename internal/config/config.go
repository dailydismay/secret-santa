package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type VKConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	BaseURL      string
}

type AuthConfig struct {
	AccessTokenSecret string
	AccessTokenTTL    time.Duration
}

type Config struct {
	Port               int
	PGConnectionString string
	VKConfig
	AuthConfig
}

func New() (Config, error) {
	setDefaults()
	loadDotenvFile()

	config := Config{
		Port: viper.GetInt("port"),
		VKConfig: VKConfig{
			ClientID:     viper.GetString("vk.client.id"),
			ClientSecret: viper.GetString("vk.client.secret"),
			RedirectURL:  viper.GetString("vk.redirectURL"),
		},
		AuthConfig: AuthConfig{
			AccessTokenTTL:    time.Duration(viper.GetInt("auth.accessToken.ttl")),
			AccessTokenSecret: viper.GetString("auth.accessToken.secret"),
		},
		PGConnectionString: viper.GetString("postgres.connectionURI"),
	}
	return config, nil
}

func loadDotenvFile() {
	configType := os.Getenv("ENV")
	if configType == "" {
		configType = "local"
	}

	_ = godotenv.Load(fmt.Sprintf("config/.env.%s", configType))

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func setDefaults() {
	viper.SetDefault("port", "3000")
}
