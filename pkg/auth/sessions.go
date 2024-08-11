package auth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"math/big"
	"time"
)

func (a *Auth) EncodeSession(id int, secret string) string {
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

	id := int(binary.BigEndian.Uint32(data[:4]))
	secret := string(data[4:])

	return id, secret, nil
}

// TODO: This is revolting, but the caller will need to persist it themselves
// and then generate their own token
func (a *Auth) NewSession(
	user *User,
	userAgent string,
	ipAddress string,
) (*Session, string, error) {
	// Populate known data
	var session Session
	session.UserId = user.Id
	session.UserAgent = userAgent
	session.IpAddress = ipAddress

	// Expiry and issued at
	currentTime := time.Now()
	session.CreatedAt = currentTime
	session.ExpiresAt = currentTime.Add(a.tokenDuration)

	// Generate secret
	secret, err := generateSessionSecret()
	if err != nil {
		return nil, "", err
	}

	// Hash secret
	hashedSecret, err := a.HashPassword(secret)
	if err != nil {
		return nil, "", err
	}
	session.HashedSecret = hashedSecret

	return &session, secret, err
}

func generateSessionSecret() (string, error) {
	characters := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz~!@#$%^&*()-_=+[{]}|;:,<.>/?"
	b := make([]byte, 16)
	for i := 0; i < 16; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		if err != nil {
			return "", err
		}
		b[i] = characters[num.Int64()]
	}

	return string(b), nil
}
