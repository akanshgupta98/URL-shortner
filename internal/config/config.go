package config

import (
	"os"
)

type Config struct {
	ServerCfg ServerConfig
	DBCfg     DatabaseConfig
	CacheCfg  CacheConfig
}

type DatabaseConfig struct {
	IP      string
	Port    string
	User    string
	Pwd     string
	DB      string
	SSLMode string
}

type ServerConfig struct {
	IP   string
	Port string
}
type CacheConfig struct {
	IP   string
	Port string
}

func Initialize() (cfg Config) {

	// Init server config
	cfg.ServerCfg.Port = getEnv(ENV_SERVER_PORT, DEFAULT_VAL_SERVER_PORT)
	cfg.ServerCfg.IP = getEnv(ENV_SERVER_IP, DEFAULT_VAL_SERVER_IP)

	// Init DB config
	cfg.DBCfg.IP = getEnv(ENV_DB_IP, DEFAULT_VAL_DB_IP)
	cfg.DBCfg.Port = getEnv(ENV_DB_PORT, DEFAULT_VAL_DB_PORT)
	cfg.DBCfg.User = getEnv(ENV_DB_USER, DEFAULT_VAL_DB_USER)
	cfg.DBCfg.Pwd = getEnv(ENV_DB_PWD, DEFAULT_VAL_DB_PWD)
	cfg.DBCfg.DB = getEnv(ENV_DB_NAME, DEFAULT_VAL_DB_NAME)
	cfg.DBCfg.SSLMode = getEnv(ENV_DB_SSL_MODE, DEFAULT_VAL_DB_SSL_MODE)

	// Init Cache Config

	return

}
func getEnv(env string, defaultVal string) (val string) {
	val = os.Getenv(env)
	if val == "" {
		return defaultVal
	}
	return val

}
