package repository

import (
	"errors"
	"fmt"
	"testing"
	"url_shortner/internal/config"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func MockConfig() (cfg config.Config) {

	cfg.DBCfg.IP = "0.0.0.0"
	cfg.DBCfg.Port = "5432"
	cfg.DBCfg.SSLMode = "disable"
	cfg.DBCfg.DB = "urls"
	cfg.DBCfg.Pwd = "password"
	cfg.DBCfg.User = "postgres"
	return
}

func generateRandomURLTest() string {
	id, _ := uuid.NewRandom()

	shortURL := id.String()[:7]
	return shortURL

}

func TestInitialize(t *testing.T) {

	tests := []struct {
		name        string
		expectedErr error
		mock        bool
	}{
		{name: "DB init success", expectedErr: nil, mock: true},                // 1
		{name: "DB init failure", expectedErr: ErrNotInitialized, mock: false}, // 2
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cfg := MockConfig()
			if tc.name == "DB init failure" {
				cfg.DBCfg.Port = "1234"
			}
			t.Cleanup(func() {
				resetDB()
			})

			actualErr := Initialize(cfg)
			if tc.expectedErr != nil {
				if assert.Error(t, actualErr) {
					assert.True(t, errors.Is(actualErr, tc.expectedErr), fmt.Sprintf("Test case: %d failed", i+1))
				}
			} else {
				assert.Equal(t, tc.expectedErr, actualErr, fmt.Sprintf("Test case: %d failed", i+1))
			}
		})

	}

}
func TestStore(t *testing.T) {
	url := generateRandomURLTest()
	tests := []struct {
		name        string
		InKey       string
		InVal       string
		DBInit      bool
		expectedErr error
	}{

		{name: "Positive: VALID CASE", InKey: url, InVal: "www.google.com", DBInit: true, expectedErr: nil},                                           // 1
		{name: "Negative: DB un-initialized", InKey: generateRandomURLTest(), InVal: "www.google.com", DBInit: false, expectedErr: ErrNotInitialized}, // 2
		{name: "Negative: Already Exist", InKey: url, InVal: "www.google.com", DBInit: true, expectedErr: ErrStore},                                   // 3
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.DBInit {
				cfg := MockConfig()
				err := Initialize(cfg)
				if err != nil {
					t.Fatal(err)
				}
			}
			t.Cleanup(func() {
				resetDB()
			})
			err := Store(tc.InKey, tc.InVal)

			if tc.expectedErr != nil {
				if assert.Error(t, err) {
					assert.True(t, errors.Is(err, tc.expectedErr), fmt.Sprintf("Test case: %d failed", i+1))
				}
			} else {
				assert.Equal(t, tc.expectedErr, err, fmt.Sprintf("Test case: %d failed", i+1))
			}
		})

	}
}
