package tests

import (
	"fmt"
	"github.com/avrilko-go/wxwork"
	"testing"
	"time"
)

const corpID = ""

const corpSecret = ""

const agentID = 1000003

type A struct {
}

func (a *A) GetToken() *wxwork.TokenInfo {
	return wxwork.NewTokenInfo("4mqw7_Uf_fsIQ1rOGM0jXJ1869W-pFWavUoUczn4nOudoFntud2pWEN9y_8-tG1b9psbgn4txS6JcntTqG1gJzHEjakZXVZmeI14geNWjQ5TR1zmlJU8UysxyvJqVlLrI-Slx2G6W3jtxb1j8-rsSMxj-_7kS6DEggmKUpEWUd9915fGMdV0y35VqvPepH7M1XR9NaqCPi7DSwP0o-T-WA", time.Now().Add(time.Hour*10))
}

func TestSendCardMessage(t *testing.T) {
	app := wxwork.New(corpID, corpSecret, agentID, wxwork.WithITokenFunc(&A{}))
	c := &wxwork.CardContent{
		Title:       "领奖通知",
		Url:         "http://www.baidu.com",
		BtnTxt:      "更多",
		Description: `<div class="gray">2016年9月26日</div> <div class="normal">恭喜你抽中iPhone 7一台，领奖码：xxxx</div><div class="highlight">请于2016年10月10日前联系行政同事领取</div>`,
	}

	err := app.SendCardMessage(&wxwork.Recipient{
		UserIDs:  []string{"8443"},
		PartyIDs: nil,
		TagIDs:   nil,
	}, c, false)

	fmt.Println(err)
}

func TestSendMpNewsMessage(t *testing.T) {
	app := wxwork.New(corpID, corpSecret, agentID, wxwork.WithITokenFunc(&A{}))
	c := []*wxwork.Article{{
		Title:            "我是标题",
		ThumbMediaId:     "3RaYQnz5f1vTdLz-sE4ka2fhf0TnCJxUzZiGicm3t8ZsA138e3yzxI4pFujjylZs4",
		Author:           "avrilko",
		ContentSourceUrl: "http://www.baiud.com",
		Content:          "哈哈哈哈",
		Digest:           "哈哈哈哈",
	}, {
		Title:            "我是标题",
		ThumbMediaId:     "3RaYQnz5f1vTdLz-sE4ka2fhf0TnCJxUzZiGicm3t8ZsA138e3yzxI4pFujjylZs4",
		Author:           "avrilko",
		ContentSourceUrl: "http://www.baiud.com",
		Content:          "哈哈哈哈",
		Digest:           "哈哈哈哈",
	}}

	err := app.SendMpNewsMessage(&wxwork.Recipient{
		UserIDs:  []string{"8443"},
		PartyIDs: nil,
		TagIDs:   nil,
	}, c, false)

	fmt.Println(err)
}

func TestSendMarkdownMessage(t *testing.T) {
	app := wxwork.New(corpID, corpSecret, agentID, wxwork.WithITokenFunc(&A{}))
	c := `您的会议室已经预定，稍后会同步到邮箱 
                                >**事项详情** 
                                >事　项：<font color=\"info\">开会</font> 
                                >组织者：@miglioguan 
                                >参与者：@miglioguan、@kunliu、@jamdeezhou、@kanexiong、@kisonwang 
                                > 
                                >会议室：<font color=\"info\">广州TIT 1楼 301</font> 
                                >日　期：<font color=\"warning\">2018年5月18日</font> 
                                >时　间：<font color=\"comment\">上午9:00-11:00</font> 
                                > 
                                >请准时参加会议。 
                                > 
                                >如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com)`

	err := app.SendMarkdownMessage(&wxwork.Recipient{
		UserIDs:  []string{"8443"},
		PartyIDs: nil,
		TagIDs:   nil,
	}, c, false)

	fmt.Println(err)
}
