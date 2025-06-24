package generators

import (
	"crypto/rand"
	"errors"
	"strings"
	"testing"
)

// errReader always returns an error.
type errReader struct{}

func (errReader) Read(p []byte) (int, error) {
	return 0, errors.New("rand fail")
}

func TestGenerateRandomPasswordLength(t *testing.T) {
	const length = 32
	pwd, err := GenerateRandomPassword(length)
	if err != nil {
		t.Fatalf("GenerateRandomPassword returned error: %v", err)
	}
	if len(pwd) != length {
		t.Fatalf("expected password length %d, got %d", length, len(pwd))
	}
}

func TestGenerateRandomPasswordCharacters(t *testing.T) {
	const length = 64
	pwd, err := GenerateRandomPassword(length)
	if err != nil {
		t.Fatalf("GenerateRandomPassword returned error: %v", err)
	}
	for i, r := range pwd {
		if !strings.ContainsRune(PasswordChars, r) {
			t.Fatalf("character %q at position %d not allowed", r, i)
		}
	}
}

func TestGenerateRandomPasswordStrength(t *testing.T) {
	pwd, err := GenerateRandomPassword(64)
	if err != nil {
		t.Fatalf("GenerateRandomPassword returned error: %v", err)
	}

	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, r := range pwd {
		switch {
		case 'A' <= r && r <= 'Z':
			hasUpper = true
		case 'a' <= r && r <= 'z':
			hasLower = true
		case '0' <= r && r <= '9':
			hasDigit = true
		default:
			if strings.ContainsRune("-=+!@#$%^&*.", r) {
				hasSpecial = true
			}
		}
	}

	if !(hasUpper && hasLower && hasDigit && hasSpecial) {
		t.Fatalf("generated password %q lacks required character variety", pwd)
	}
}

func TestGenerateRandomPasswordError(t *testing.T) {
	old := rand.Reader
	rand.Reader = errReader{}
	defer func() { rand.Reader = old }()

	pwd, err := GenerateRandomPassword(8)
	if err == nil {
		t.Fatal("expected error")
	}
	if pwd != "" {
		t.Fatalf("expected empty password on error, got %q", pwd)
	}
}
