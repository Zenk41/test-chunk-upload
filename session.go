package main

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"
)

type Session struct {
	Data      map[string]interface{}
	ExpiresAt time.Time
}

type SessionManager struct {
	store map[string]*Session
	mu    sync.RWMutex
}

func NewSessionManage() *SessionManager {
	return &SessionManager{
		store: make(map[string]*Session),
	}
}

func (sm *SessionManager) CreateSession(Data map[string]interface{}) (string, error) {
	sessionID, err := generateSessionID()
	if err != nil {
		return "", err
	}
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.store[sessionID] = &Session{
		Data:      Data,
		ExpiresAt: time.Now().Add(30 * time.Minute),
	}
	return sessionID, nil
}

func generateSessionID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}


func (sm *SessionManager) getSession(sessionID string) (*Session, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	session, exists := sm.store[sessionID]
	if exists && session.ExpiresAt.After(time.Now()) {
		return session, true
	}
	return nil, false
}

func (sm *SessionManager) DeleteSession(sessionID string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.store, sessionID)
}