package server

import (
	"os"
	"strings"
)

type EnvConfig struct {
	Port        string
	UseHttp2    bool
	CorsOrigins []string
}

func LoadConfig() (*EnvConfig, error) {
	useHttp2Str := os.Getenv("USE_HTTP2")
	useHttp2 := useHttp2Str == "true"

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	origins := strings.Split(os.Getenv("CORS_ORIGINS"), ",")
	return &EnvConfig{
		Port:        port,
		UseHttp2:    useHttp2,
		CorsOrigins: origins,
	}, nil
}
