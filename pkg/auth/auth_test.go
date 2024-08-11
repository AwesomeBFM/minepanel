package auth_test

import (
	"github.com/awesomebfm/minepanel/pkg/auth"
	"testing"
	"time"
)

func TestAuth_PasswordHashing(t *testing.T) {
	ath := auth.NewAuth(
		&auth.Params{
			Memory:      64 * 1024,
			Iterations:  3,
			Parallelism: 2,
			SaltLength:  16,
			KeyLength:   32,
		},
		90*24*time.Hour,
	)

	hashed, err := ath.HashPassword("deFvd*f1W8Y9")
	if err != nil {
		t.Errorf("Did not expect error while hashing password but got: %v", err)
	}

	result, err := ath.HashMatches("deFvd*f1W8Y9", hashed)
	if err != nil {
		t.Errorf("Did not expect error while comparing hash but got: %v", err)
	}
	if !result {
		t.Errorf("Expected result %v but got %v", true, result)
	}
}
