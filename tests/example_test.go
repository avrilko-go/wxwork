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
	return wxwork.NewTokenInfo("xxxxxxxx", time.Now().Add(time.Hour*10))
}

func TestSendMessage(t *testing.T) {
	app := wxwork.New(corpID, corpSecret, agentID, wxwork.WithITokenFunc(&A{}))

	for {
		err := app.SendTextMessage(&wxwork.Recipient{
			UserIDs:  []string{"8443"},
			PartyIDs: nil,
			TagIDs:   nil,
		}, "hahha", false)

		fmt.Println(err)
		time.Sleep(time.Second * 1)

		err = app.SendImageMessage(&wxwork.Recipient{
			UserIDs:  []string{"8443"},
			PartyIDs: nil,
			TagIDs:   nil,
		}, "3OrtsJ5Q_bTmtHvPSzaKRAc7svk9aEyLRNiZKbSogCtmUi4gsiP3IG3CjQDctst7h", false)

		fmt.Println(err)
		time.Sleep(time.Second * 1)
	}

}
