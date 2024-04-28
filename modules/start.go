package modules

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Start(b *gotgbot.Bot, ctx *ext.Context) error {
	message := `
<b>👋 𝙷𝙴𝚈</b> %s

<b>𝙸 𝙰𝙼 𝙰 𝙱𝙾𝚃 𝚃𝙾 𝚁𝚄𝙽 𝚈𝙾𝚄𝚁 𝙶𝙾 𝙲𝙾𝙳𝙴𝚂 🌨</b>

<pre>/run 

package main 

import "fmt"

func main() {
        fmt.Println("Hello Mukund")
}</pre>

`
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf(message, ctx.EffectiveUser.FirstName), &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}
