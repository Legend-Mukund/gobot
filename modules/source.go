package modules

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Source(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintln("https://github.com/MukundSinghRajput/gobot"), &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}
