package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fernet/fernet-go"
)

const (
	secret         = "cw_0x689RpI-jtRR7oE8h_eQsKImvJapLeSbXpwF4e4="
	tokenLongevity = 7 * 24 * time.Hour
)

func main() {
	key, err := fernet.DecodeKey(secret)
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
	key, err = fernet.DecodeKey(secret)
	if err != nil {
		panic(err)
	}

	msg := decodeIndex(key, token)
	log.Println(string(msg))

	// msg == gAAAAABhZG_zULYCUxwlP4-9RMWVsrTQY7qRzs2dvB-eumuNknrbO9y0TF949FP4PVgIB0O4WXzkWjfngmu2_APX66yHPdNuMQ==
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
