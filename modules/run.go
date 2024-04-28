package modules

import (
	"fmt"

	"github.com/MukundSinghRajput/gobot/helper"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func RunCode(b *gotgbot.Bot, ctx *ext.Context) error {
	text := ctx.EffectiveMessage.Text
	if len(text) < 5 {
		_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintln("So what shall i run bro ?"), &gotgbot.SendMessageOpts{ParseMode: "html"})
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		return nil
	}
	output, err := helper.Runner(text[4:])
	if err != nil {
		_, errore := ctx.EffectiveMessage.Reply(b, fmt.Sprintln(err), &gotgbot.SendMessageOpts{ParseMode: "html"})
		if errore != nil {
			return fmt.Errorf(errore.Error())
		}
		return fmt.Errorf(err.Error())
	}
	var compileErrr string
	if output.CompileErrors != "" {
		compileErrr = output.CompileErrors
	} else {
		compileErrr = "None"
	}
	outputStr := output.Output
	message := fmt.Sprintf("Code:\n\n```\n%v\n```\nCompile Errors:\n\n```\n%v\n```\nOutput:\n\n```\n%v\n```\n", text[4:], compileErrr, outputStr)
	_, err = ctx.EffectiveMessage.Reply(b, message, &gotgbot.SendMessageOpts{ParseMode: "markdown"})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}
