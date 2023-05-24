package encryptor

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCipherText_BadIV(t *testing.T) {
	keyData := base64.StdEncoding.EncodeToString([]byte("1234567890123456"))
	badData := base64.StdEncoding.EncodeToString([]byte("1"))
	_, err := GetCipherText(keyData, badData, badData)
	assert.Error(t, err)
}
