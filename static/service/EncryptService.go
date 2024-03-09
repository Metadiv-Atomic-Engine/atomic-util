package service

import (
	"github.com/Metadiv-Atomic-Engine/aes"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

var EncryptService = new(encryptService)

type encryptService struct{}

func (e *encryptService) Encrypt(content []byte) []byte {
	return aes.EncryptBytesToBytes(content, atomic.Engine.EnvString(atomic.DB_ENCRYPT_KEY))
}

func (e *encryptService) Decrypt(content []byte) []byte {
	return aes.DecryptBytesToBytes(content, atomic.Engine.EnvString(atomic.DB_ENCRYPT_KEY))
}
