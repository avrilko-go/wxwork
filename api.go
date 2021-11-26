package wxwork

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// execMessageSend 消息推送发送
func (w *wxApp) execMessageSend(req *reqMessage) (respMessageSend, error) {
	var resp respMessageSend
	err := w.executeJSONPost("/cgi-bin/message/send", req, &resp, true)
	if err != nil {
		return respMessageSend{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return respMessageSend{}, bizErr
	}
	return resp, nil
}

// executeJSONPost json post快捷方法
func (w *wxApp) executeJSONPost(path string, req IBody, respObj interface{}, needAccessToken bool) error {
	urlStr := w.createUrl(path, needAccessToken)
	body, err := req.intoBody()
	fmt.Println(string(body))

	if err != nil {
		return err
	}

	resp, err := w.client.Post(urlStr, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(respObj)
}

// createUrl 创建请求url
func (w *wxApp) createUrl(path string, needAccessToken bool) string {
	url := fmt.Sprintf("%s%s", DefaultAPIHost, path)
	if needAccessToken {
		url = fmt.Sprintf("%s?access_token=%s", url, w.getToken())
	}
	return url
}
