package hashing

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

type HashOptions struct {
	Value   string
	Salt    []byte
	Time    uint32
	Memory  uint32
	Threads uint8
	KeyLen  uint32
}

var DefaultHashOptions HashOptions = HashOptions{
	Time:    1,
	Memory:  65536,
	Threads: 4,
	KeyLen:  32,
}

const ArgonFormatStr string = "$argon2id$v=19$m=65536,t=1,p=4$"

func Hash(options HashOptions) (string, error) {
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	options.Salt = salt

	hash := argon2.IDKey(
		[]byte(options.Value),
		options.Salt,
		options.Time,
		options.Memory,
		options.Threads,
		options.KeyLen,
	)

	b64salt := base64.RawStdEncoding.EncodeToString(salt)
	b64hash := base64.RawStdEncoding.EncodeToString(hash)

	return ArgonFormatStr + b64salt + "$" + b64hash, nil
}

func VerifyPassword(password, encodedHash string) (bool, error) {
	fmt.Println(encodedHash)
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false, ErrInvalidHashFormat
	}

	// Get options
	var memory, iterations, parallelism int
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism)
	if err != nil {
		return false, err
	}

	// Decode salt and original hash
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	originalHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	// Check current password hash
	newHash := argon2.IDKey(
		[]byte(password),
		salt,
		uint32(iterations),
		uint32(memory),
		uint8(parallelism),
		uint32(len(originalHash)),
	)

	// Similar passwords
	return subtle.ConstantTimeCompare(originalHash, newHash) == 1, nil
}
