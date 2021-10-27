package wxwork

import (
	"net/http"
	"time"
)

type wxApp struct {
	agentID    int64
	corpID     string
	corpSecret string
	token      *TokenInfo
	tokenFunc  IToken
	client     *http.Client
}

func New(corpID, corpSecret string, agentID int64, opt ...AppOptionFunc) *wxApp {
	w := &wxApp{agentID: agentID, corpID: corpID, corpSecret: corpSecret, client: &http.Client{}, token: NewTokenInfo("", time.Now().Add(-time.Hour))}
	for _, fn := range opt {
		fn(w)
	}
	return w
}
