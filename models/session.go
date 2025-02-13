package models

import (
	"database/sql"
	"fmt"
	"github.com/sviut/photo-locker/rand"
)

const (
	MinBytesPerToken = 32
)

type Session struct {
	ID     int
	UserId int
	// only set when creating a new session
	Token     string
	TokenHash string
}

type SessionService struct {
	DB            *sql.DB
	BytesPerToken int
}

func (ss *SessionService) Create(userId int) (*Session, error) {
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}

	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("could not generate session token: %v", err)
	}
	session := Session{
		UserId: userId,
		Token:  token,
	}
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	return nil, nil
}
