package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func UnsetAllENV() {
	os.Unsetenv(ENV_SERVER_PORT)
	os.Unsetenv(ENV_SERVER_IP)
	os.Unsetenv(ENV_DB_IP)
	os.Unsetenv(ENV_DB_PORT)

}
func TestInitialize(t *testing.T) {
	defer UnsetAllENV()
	tests := []struct {
		name        string
		Port        string
		IP          string
		DB_IP       string
		DB_Port     string
		wantPort    string
		wantIP      string
		wantDB_IP   string
		wantDB_Port string
	}{
		{name: "ALL DEFAULT VALUES", wantPort: DEFAULT_VAL_SERVER_PORT, wantIP: DEFAULT_VAL_SERVER_IP, wantDB_IP: DEFAULT_VAL_DB_IP, wantDB_Port: DEFAULT_VAL_DB_PORT},
		{name: "ALL CUSTOM VALUES", Port: "1234", wantPort: "1234", IP: "192.168.7.1", wantIP: "192.168.7.1", DB_IP: "192.168.7.2", wantDB_IP: "192.168.7.2", DB_Port: "7654", wantDB_Port: "7654"},
		{name: "MIXED VALUES", wantPort: DEFAULT_VAL_SERVER_PORT, IP: "192.18.1.1", wantIP: "192.18.1.1", wantDB_IP: DEFAULT_VAL_DB_IP, wantDB_Port: DEFAULT_VAL_DB_PORT},
	}

	for _, tc := range tests {
		// Reset all at the start of test.
		UnsetAllENV()
		// Set if value given
		if tc.IP != "" {
			os.Setenv(ENV_SERVER_IP, tc.IP)
		}
		// Set if value given
		if tc.Port != "" {
			os.Setenv(ENV_SERVER_PORT, tc.Port)
		}
		// Set if value given
		if tc.DB_Port != "" {
			os.Setenv(ENV_DB_PORT, tc.DB_Port)
		}
		// Set if value given
		if tc.DB_IP != "" {
			os.Setenv(ENV_DB_IP, tc.DB_IP)
		}

		// Call the Initialize func.
		result := Initialize()
		assert.Equal(t, tc.wantIP, result.ServerCfg.IP, "IP does not match the expected")
		assert.Equal(t, tc.wantPort, result.ServerCfg.Port, "Port does not match the expected")
		assert.Equal(t, tc.wantDB_IP, result.DBCfg.IP, "DB IP does not match the expected")
		assert.Equal(t, tc.wantDB_Port, result.DBCfg.Port, "DB port does not match the expected")
	}
}
