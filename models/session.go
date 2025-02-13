package models

import (
	"database/sql"
	"fmt"
	"github.com/sviut/photo-locker/rand"
)

type Session struct {
	ID     int
	UserId int
	// only set when creating a new session
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userId int) (*Session, error) {
	token, err := rand.SessionToken()
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
