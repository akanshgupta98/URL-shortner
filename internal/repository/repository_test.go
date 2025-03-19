package repository

import (
	"errors"
	"fmt"
	"testing"
	"url_shortner/internal/config"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func generateRandomURLTest() string {
	id, _ := uuid.NewRandom()

	shortURL := id.String()[:7]
	return shortURL

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
		{name: "Negative: DB un-initialized", InKey: generateRandomURLTest(), InVal: "www.google.com", DBInit: false, expectedErr: ErrNotInitialized}, // 1
		{name: "Positive: VALID CASE", InKey: url, InVal: "www.google.com", DBInit: true, expectedErr: nil},                                           // 2
		{name: "Negative: Already Exist", InKey: url, InVal: "www.google.com", DBInit: true, expectedErr: ErrStore},                                   // 3
	}

	for i, tc := range tests {
		if tc.DBInit {
			cfg := config.Initialize()
			Initialize(cfg)
		}
		err := Store(tc.InKey, tc.InVal)

		if tc.expectedErr != nil {
			if assert.Error(t, err) {
				assert.True(t, errors.Is(err, tc.expectedErr), fmt.Sprintf("Test case: %d failed", i+1))
			}
		} else {
			assert.Equal(t, tc.expectedErr, err, fmt.Sprintf("Test case: %d failed", i+1))
		}

	}
}
