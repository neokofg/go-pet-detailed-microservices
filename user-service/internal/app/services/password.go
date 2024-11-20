package services

import (
	"errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

const (
	minPasswordLength = 8
	maxPasswordLength = 72
	defaultCost       = 12
)

var (
	ErrPasswordTooShort = errors.New("password is too short")
	ErrPasswordTooLong  = errors.New("password is too long")
)

type PasswordService struct {
	logger *zap.Logger
}

func NewPasswordService(logger *zap.Logger) *PasswordService {
	return &PasswordService{
		logger: logger,
	}
}

func (s *PasswordService) HashPassword(password string) (string, error) {
	if err := s.validate(password); err != nil {
		return "", err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), defaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (s *PasswordService) Compare(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func (s *PasswordService) validate(password string) error {
	length := len(password)

	if length < minPasswordLength {
		return ErrPasswordTooShort
	}

	if length > maxPasswordLength {
		return ErrPasswordTooLong
	}

	return nil
}
