package teleclient

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

// const Apitoken = "5550390206:AAFfI965rdF6vVoGsmPtnR6ZMi845D_gKyU"  7283274026
const Apitoken = "7693615380:AAFVJLI33EhY-CJYFP7puZYsSzlZfQwdoto"

const ChatID = -1002375968132 //机器人推送个人账户
const GroupDao = -4190696656  //机器人推送群消息
const ButtonText = "24小时全自动TRX兑换机器人\n1️⃣进U即兑,全自动返TRX,1U起兑\n2️⃣如您的trx不足转账，可联系客服预支\n3️⃣不要使用交易所转账，\ntrx兑换客服：@Tronu2bot\n门店客服: @66666\n门店地址：6666汇宝货币"

func Telecli() (*tgbotapi.BotAPI, error) {
	// 填入你的 Bot API Token

	bot, err := tgbotapi.NewBotAPI(Apitoken)
	if err != nil {
		log.Panic(err)

	}

	// 设置 Debug 模式打印日志
	//bot.Debug = true

	return bot, nil
}
