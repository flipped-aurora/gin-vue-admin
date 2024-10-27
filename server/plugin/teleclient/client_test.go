package teleclient

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"testing"
)

func TestSendmsg(t *testing.T) {
	bot, err := Telecli()
	if err != nil {
		return
	}

	//"id": 6880320852,  mikewang ID
	msg := tgbotapi.NewMessage(ChatID,
		`24小时全自动TRX兑换机器人
			1️⃣进U即兑,全自动返TRX,1U起兑
			2️⃣如您的trx不足转账，可联系客服预支
			3️⃣不要使用交易所转账，
			trx兑换客服：@Tronu2bot
			门店客服: @66666
			门店地址：6666汇宝货币`)
	send, err := bot.Send(msg)
	if err != nil {
		return
	}
	fmt.Println(send)
}

func TestKeyboardfn(t *testing.T) {
	//Keyboardfn(Telecli())
	//CommandHandling(Telecli())
	//InlineKeyboard(Telecli())
	//
	//err := NewTelemsg("测试消息", "我是地址xxxx", "我是消息6666").SendmsgHtml(Telecli())
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//	text := `*bold \*text*
	//_italic \*text_
	//__underline__
	//~strikethrough~
	//||spoiler||
	//*bold _italic bold ~italic bold strikethrough ||italic bold strikethrough spoiler||~ __underline italic bold___ bold*
	//[inline URL](http://www.example.com/)
	//[inline mention of a user](tg://user?id=123456789)
	//![👍](tg://emoji?id=5368324170671202286)
	//>Block quotation started
	//>Block quotation continued
	//>The last line of the block quotation`

	htmltext := `
				<b>bold</b>, <strong>bold</strong>
				<i>italic</i>, <em>italic</em>
				<u>underline</u>, <ins>underline</ins>
				<s>strikethrough</s>, <strike>strikethrough</strike>, <del>strikethrough</del>
				<span class="tg-spoiler">spoiler</span>, <tg-spoiler>spoiler</tg-spoiler>
				<b>bold <i>italic bold <s>italic bold strikethrough <span class="tg-spoiler">italic bold strikethrough spoiler</span></s> <u>underline italic bold</u></i> bold</b>
				<a href="http://www.example.com/">inline URL</a>
				<a href="tg://user?id=123456789">inline mention of a user</a>
				<tg-emoji emoji-id="5368324170671202286">👍</tg-emoji>
				<code>inline fixed-width code</code>
				<pre>pre-formatted fixed-width code block</pre>
				<pre><code class="language-python">pre-formatted fixed-width code block written in the Python programming language</code></pre>
				<blockquote>Block quotation started\nBlock quotation continued\nThe last line of the block quotation</blockquote>
				`
	bot, _ := Telecli()
	BotMark(bot, "HTML", -4190696656, htmltext)
}
