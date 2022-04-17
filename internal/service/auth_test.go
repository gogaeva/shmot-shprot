package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePasswordHash(t *testing.T) {
	const password = "qwerty123"
	const hash = "68687364616875686b75daaad6e5604e8e17bd9f108d91e26afe6281dac8fda0091040a7a6d7bd9b43b5"
	assert.Equal(t, hash, generatePasswordHash(password))

}
