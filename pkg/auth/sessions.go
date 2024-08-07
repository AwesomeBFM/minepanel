package auth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"time"

	"golang.org/x/crypto/argon2"
)

func (a *Auth) EncodeSession(id int, secret []byte) string {
	idBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(idBytes, uint32(id))

	data := append(idBytes, secret...)

	return base64.StdEncoding.EncodeToString(data)
}

func DecodeSession(token string) (int, string, error) {
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return 0, "", err
	}

	id := int(binary.BigEndian.Uint32(data[:8]))
	secret := string(data[8:])

	return id, secret, nil
}

// TODO: This is revolting, but the caller will need to persist it themselves
// and then generate their own token
func (a *Auth) NewSession(
	user *User,
	userAgent string,
	ipAddress string,
) (session *Session, secret []byte, err error) {
	// Populate known data
	session.UserId = user.Id
	session.UserAgent = userAgent
	session.IpAddress = ipAddress

	// Expiry and issued at
	currentTime := time.Now()
	session.CreatedAt = currentTime
	session.ExpiresAt = currentTime.Add(a.tokenDuration)

	// Generate secret
	secret, err = generateSessionSecret()
	if err != nil {
		return nil, nil, err
	}

	// Hash secret
	hashedSecret, err := a.hashSecret(secret)
	if err != nil {
		return nil, nil, err
	}
	session.HashedSecret = string(hashedSecret)

	return
}

func (a *Auth) hashSecret(secret []byte) ([]byte, error) {
	// Generate salt
	salt, err := a.generateRandomBytes(a.params.SaltLength)
	if err != nil {
		return nil, err
	}

	// Generate hash with Argon2id
	hash := argon2.IDKey(
		secret,
		salt,
		a.params.Iterations,
		a.params.Memory,
		a.params.Parallelism,
		a.params.KeyLength,
	)
	return hash, nil
}

func generateSessionSecret() ([]byte, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	return b, err
}
