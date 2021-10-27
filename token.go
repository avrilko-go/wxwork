package wxwork

import (
	"sync"
	"time"
)

type TokenInfo struct {
	content   string
	expiresAt time.Time
	lock      sync.Mutex
}

func NewTokenInfo(token string, expiresAt time.Time) *TokenInfo {
	return &TokenInfo{
		content:   token,
		expiresAt: expiresAt,
	}
}

type IToken interface {
	GetToken() *TokenInfo
}

func (w *wxApp) getToken() string {
	w.token.lock.Lock()
	defer w.token.lock.Unlock()
	if w.token.expiresAt.Before(time.Now()) {
		w.token = w.tokenFunc.GetToken()
	}

	return w.token.content
}
