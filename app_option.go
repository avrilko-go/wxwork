package wxwork

import "net/http"

const DefaultAPIHost = "https://qyapi.weixin.qq.com"

type AppOptionFunc func(w *WxApp)

// WithAgentID 配置应用id
func WithAgentID(agentID int64) AppOptionFunc {
	return func(w *WxApp) {
		w.agentID = agentID
	}
}

// WithCorpID 配置企业id
func WithCorpID(corpID string) AppOptionFunc {
	return func(w *WxApp) {
		w.corpID = corpID
	}
}

// WithCorpSecret 配置app密钥
func WithCorpSecret(corpSecret string) AppOptionFunc {
	return func(w *WxApp) {
		w.corpSecret = corpSecret
	}
}

// WithITokenFunc 配置自定义token获取方法
func WithITokenFunc(tokenFunc IToken) AppOptionFunc {
	return func(w *WxApp) {
		w.tokenFunc = tokenFunc
	}
}

func WithHttpClient(client *http.Client) AppOptionFunc {
	return func(w *WxApp) {
		w.client = client
	}
}
