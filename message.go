package wxwork

import "encoding/json"

// SendTextMessage 发送文本消息
func (w *wxApp) SendTextMessage(recipient *Recipient, content string, isSafe bool) error {
	return w.sendMessage(recipient, "text", map[string]interface{}{"content": content}, isSafe)
}

// SendImageMessage 发送图片消息
func (w *wxApp) SendImageMessage(recipient *Recipient, mediaId string, isSafe bool) error {
	return w.sendMessage(recipient, "image", map[string]interface{}{"media_id": mediaId}, isSafe)
}

// SendCardMessage 发送卡片消息
func (w *wxApp) SendCardMessage(recipient *Recipient, content *CardContent, isSafe bool) error {
	cardParams := make(map[string]interface{})
	b, _ := json.Marshal(content)
	_ = json.Unmarshal(b, &cardParams)
	return w.sendMessage(recipient, "textcard", cardParams, isSafe)
}

// SendMpNewsMessage 发送图文消息
func (w *wxApp) SendMpNewsMessage(recipient *Recipient, content []*Article, isSafe bool) error {
	return w.sendMessage(recipient, "mpnews", map[string]interface{}{"articles": content}, isSafe)
}

// SendMarkdownMessage 发送markdown消息
func (w *wxApp) SendMarkdownMessage(recipient *Recipient, content string, isSafe bool) error {
	return w.sendMessage(recipient, "markdown", map[string]interface{}{"content": content}, isSafe)
}

// SendMessage 发送消息基类
func (w *wxApp) sendMessage(recipient *Recipient,
	msgType string,
	content map[string]interface{},
	isSafe bool) error {

	req := &reqMessage{
		ToUser:  recipient.UserIDs,
		ToParty: recipient.PartyIDs,
		ToTag:   recipient.TagIDs,
		AgentID: w.agentID,
		MsgType: msgType,
		Content: content,
		IsSafe:  isSafe,
	}

	_, err := w.execMessageSend(req)
	if err != nil {
		return err
	}
	return nil
}
