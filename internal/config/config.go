package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Host           string
	Port           string
	AllowedOrigins []string
	CookieName     string
	CookieSecret   []byte
}

func New() Config {
	host := getEnvDefault("HOST", "0.0.0.0")
	port := getEnvDefault("PORT", "8080")
	return Config{
		Host: host,
		Port: port,
		AllowedOrigins: strings.Split(
			getEnvDefault("ALLOWED_ORIGINS", fmt.Sprintf("http://%s:%s,https://%s:%s", host, port, host, port)), ",",
		),
	}
}

func getEnvDefault(key, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return def
}
