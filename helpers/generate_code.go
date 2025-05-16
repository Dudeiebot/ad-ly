package helpers

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
)

func GenerateShortCode(userID uint64) (string, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, uint64(userID))
	if err != nil {
		return "", err
	}
	userIDBytes := buf.Bytes()

	randomBytes := make([]byte, 8)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	combined := append(randomBytes, userIDBytes...)

	hasher := sha256.New()
	hasher.Write(combined)
	hashSum := hasher.Sum(nil)

	encoded := base64.URLEncoding.EncodeToString(hashSum)

	return encoded[:8], nil
}
