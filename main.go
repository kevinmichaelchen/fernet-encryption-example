package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fernet/fernet-go"
)

const (
	// Fernet key must be 32 random bytes base64-encoded.
	// You can generate one on your Terminal with:
	// dd if=/dev/urandom bs=32 count=1 2>/dev/null | openssl base64
	encodedKey     = "eVmrm+HzKImS3ezoXRQ2QfmIyaJ9SJlvKn5s3L8upKQ="
	tokenLongevity = 7 * 24 * time.Hour
)

func main() {
	key, err := fernet.DecodeKey(encodedKey)
	if err != nil {
		panic(err)
	}

	// ENCODE
	token, err := encodeIndex(key, "index=100")
	if err != nil {
		panic(err)
	}
	log.Println(string(token))

	// DECODE
	key, err = fernet.DecodeKey(encodedKey)
	if err != nil {
		panic(err)
	}

	msg := decodeIndex(key, token)
	log.Println(string(msg))

	// msg == gAAAAABhZHNntqnpp73b-VRCrx1pvOT08HPyF15zrC08Vk0WuHkDP92wb4cjlvFeiYy2rYSdniMvafTYNqrFeJKOEv9C7T4pUw==
}

func encodeIndex(key *fernet.Key, in string) ([]byte, error) {
	// TODO make sure it's UTF-8 encoded before encrypting
	token, err := fernet.EncryptAndSign([]byte(in), key)
	if err != nil {
		return nil, fmt.Errorf("failed to encode: %w", err)
	}

	return token, nil
}

func decodeIndex(key *fernet.Key, in []byte) []byte {
	return fernet.VerifyAndDecrypt(in, tokenLongevity, []*fernet.Key{key})
}
