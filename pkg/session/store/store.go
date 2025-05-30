package store

import (
	"errors"
	"log/slog"

	"github.com/phrkdll/monomatch/pkg/session"
)

type SessionStore struct {
	sessions map[session.SessionId]*session.Session
}

var (
	instance                *SessionStore
	ErrSessionAlreadyExists = errors.New("session already exists")
	ErrSessionDoesNotExist  = errors.New("session does not exists")
)

func Instance() *SessionStore {
	if instance == nil {
		slog.Info("creating session store instance")
		instance = &SessionStore{sessions: make(map[session.SessionId]*session.Session)}
	}

	return instance
}

func (ss *SessionStore) Add(s *session.Session) error {
	if _, ok := ss.sessions[s.ID]; ok {
		return ErrSessionAlreadyExists
	}

	ss.sessions[s.ID] = s
	slog.Info("session added", "count", len(ss.sessions))

	return nil
}

func (ss *SessionStore) Exists(id session.SessionId) bool {
	if _, ok := ss.sessions[id]; ok {
		return true
	}

	return false
}

func (ss *SessionStore) Get(id session.SessionId) (*session.Session, error) {
	if session, ok := ss.sessions[id]; ok {
		return session, nil
	}

	return nil, ErrSessionDoesNotExist
}
