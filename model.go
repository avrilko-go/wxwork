package wxwork

import (
	"encoding/json"
	"strings"
)

type respCommon struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (r *respCommon) IsOK() bool {
	return r.ErrCode == 0
}

func (r *respCommon) TryIntoErr() error {
	if r.IsOK() {
		return nil
	}

	return &wxWorkClientError{
		Code: r.ErrCode,
		Msg:  r.ErrMsg,
	}
}

type respMessageSend struct {
	respCommon

	InvalidUsers   string `json:"invaliduser"`
	InvalidParties string `json:"invalidparty"`
	InvalidTags    string `json:"invalidtag"`
}

type reqMessage struct {
	ToUser  []string
	ToParty []string
	ToTag   []string
	AgentID int64
	MsgType string
	Content map[string]interface{}
	IsSafe  bool
}

// 将发送的结构体转换为json
func (r *reqMessage) intoBody() ([]byte, error) {
	safe := 0
	if r.IsSafe {
		safe = 1
	}

	objMap := map[string]interface{}{
		"msgtype": r.MsgType,
		"agentid": r.AgentID,
		"safe":    safe,
	}

	objMap[r.MsgType] = r.Content
	objMap["touser"] = strings.Join(r.ToUser, "|")
	objMap["toparty"] = strings.Join(r.ToParty, "|")
	objMap["totag"] = strings.Join(r.ToTag, "|")
	return json.Marshal(objMap)
}
