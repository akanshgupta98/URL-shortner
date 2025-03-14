package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	defer func() {
		os.Unsetenv(PORT_ENV)
		os.Unsetenv(IP_ENV)
		os.Unsetenv(DB_IP_ENV)
		os.Unsetenv(DB_PORT_ENV)
	}()
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
		{name: "ALL DEFAULT VALUES", wantPort: PORT_DEFAULT_VAL, wantIP: IP_DEFAULT_VAL, wantDB_IP: DB_IP_DEFAULT_VAL, wantDB_Port: DB_PORT_DEFAULT_VAL},
		{name: "ALL CUSTOM VALUES", Port: "1234", wantPort: "1234", IP: "192.168.7.1", wantIP: "192.168.7.1", DB_IP: "192.168.7.2", wantDB_IP: "192.168.7.2", DB_Port: "7654", wantDB_Port: "7654"},
		{name: "MIXED VALUES", wantPort: PORT_DEFAULT_VAL, IP: "192.18.1.1", wantIP: "192.18.1.1", wantDB_IP: DB_IP_DEFAULT_VAL, wantDB_Port: DB_PORT_DEFAULT_VAL},
	}

	for _, tc := range tests {
		// Reset all at the start of test.
		os.Unsetenv(PORT_ENV)
		os.Unsetenv(IP_ENV)
		os.Unsetenv(DB_IP_ENV)
		os.Unsetenv(DB_PORT_ENV)
		// Set if value given
		if tc.IP != "" {
			os.Setenv(IP_ENV, tc.IP)
		}
		// Set if value given
		if tc.Port != "" {
			os.Setenv(PORT_ENV, tc.Port)
		}
		// Set if value given
		if tc.DB_Port != "" {
			os.Setenv(DB_PORT_ENV, tc.DB_Port)
		}
		// Set if value given
		if tc.DB_IP != "" {
			os.Setenv(DB_IP_ENV, tc.DB_IP)
		}

		// Call the Initialize func.
		result := Initialize()
		assert.Equal(t, tc.wantIP, result.IP, "IP does not match the expected")
		assert.Equal(t, tc.wantPort, result.Port, "Port does not match the expected")
		assert.Equal(t, tc.wantDB_IP, result.DB_IP, "DB IP does not match the expected")
		assert.Equal(t, tc.wantDB_Port, result.DB_Port, "DB port does not match the expected")
	}
}
