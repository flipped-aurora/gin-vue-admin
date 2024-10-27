package teleclient

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telemsg struct {
	Name    string `json:"name"`    //事件名称
	Address string `json:"address"` //用户地址
	Msg     string `json:"msg"`     //消息内容
}

// NewTelemsg 构造函数
func NewTelemsg(name string, address string, msg string) *Telemsg {
	return &Telemsg{
		Name:    name,
		Address: address,
		Msg:     msg,
	}
}

// SendmsgHtml 发送消息
func (t *Telemsg) SendmsgHtml(bot *tgbotapi.BotAPI, chatid int64) error {
	//<i>斜体
	msgText := fmt.Sprintf("<b>🔔  GroundDao平台 %v!\n\n</b> ├ 💹Address:<code>%v\n</code><tg-spoiler>├ 💯 Message: %v\n</tg-spoiler>", t.Name, t.Address, t.Msg)
	// 发送消息
	msg := tgbotapi.NewMessage(chatid, msgText)
	msg.ParseMode = "HTML"
	if _, err := bot.Send(msg); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func BotMark(bot *tgbotapi.BotAPI, parsemod string, chatid int64, msgtext string) {

	msg := tgbotapi.NewMessage(chatid, msgtext)
	msg.ParseMode = parsemod

	// 发送消息
	if _, err := bot.Send(msg); err != nil {
		fmt.Println(err)
	}
}
