package initializers

import (
	"github.com/MukundSinghRajput/gobot/modules"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func LoadModules(dis *ext.Dispatcher) {
	dis.AddHandler(handlers.NewCommand("start", modules.Start))
	dis.AddHandler(handlers.NewCommand("run", modules.RunCode))
	dis.AddHandler(handlers.NewCommand("source", modules.Source))
	dis.AddHandler(handlers.NewCommand("repo", modules.Source))
}
