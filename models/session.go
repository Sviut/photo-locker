package models

import "database/sql"

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
	return nil, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	return nil, nil
}
