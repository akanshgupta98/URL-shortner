package config

import "os"

type ServerConfig struct {
	Port    string
	IP      string
	DB_IP   string
	DB_Port string
}

func Initialize() (cfg ServerConfig) {

	if cfg.Port = os.Getenv(PORT_ENV); cfg.Port == "" {
		cfg.Port = PORT_DEFAULT_VAL
	}
	if cfg.IP = os.Getenv(IP_ENV); cfg.IP == "" {
		cfg.IP = IP_DEFAULT_VAL
	}
	if cfg.DB_IP = os.Getenv(DB_IP_ENV); cfg.DB_IP == "" {
		cfg.DB_IP = DB_IP_DEFAULT_VAL
	}
	if cfg.DB_Port = os.Getenv(DB_PORT_ENV); cfg.DB_Port == "" {
		cfg.DB_Port = DB_PORT_DEFAULT_VAL
	}
	return

}
