package config

import "os"

type ServerConfig struct {
	Port string
	IP   string
}

func Initialize() (cfg ServerConfig) {

	if port := os.Getenv(PORT_ENV); port == "" {
		cfg.Port = PORT_DEFAULT_VAL
	}
	if ip := os.Getenv(IP_ENV); ip == "" {
		cfg.IP = IP_DEFAULT_VAL
	}
	return

}
