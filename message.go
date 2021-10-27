package wxwork

// SendTextMessage 发送文本消息
func (w *WxApp) SendTextMessage(recipient *Recipient, content string, isSafe bool) error {
	return w.SendMessage(recipient, "text", map[string]interface{}{"content": content}, isSafe)
}

// SendImageMessage 发送图片消息
func (w *WxApp) SendImageMessage(recipient *Recipient, mediaId string, isSafe bool) error {
	return w.SendMessage(recipient, "image", map[string]interface{}{"media_id": mediaId}, isSafe)
}

// SendMessage 发送消息基类
func (w *WxApp) SendMessage(recipient *Recipient,
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
