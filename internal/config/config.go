package config

import "os"

type ServerConfig struct {
	Port string
	IP   string
}

func Initialize() (cfg ServerConfig) {

	if cfg.Port = os.Getenv(PORT_ENV); cfg.Port == "" {
		cfg.Port = PORT_DEFAULT_VAL
	}
	if cfg.IP = os.Getenv(IP_ENV); cfg.IP == "" {
		cfg.IP = IP_DEFAULT_VAL
	}
	return

}
