package initializers

import (
	"github.com/MukundSinghRajput/gobot/modules"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func LoadModules(dis *ext.Dispatcher) {
	dis.AddHandler(handlers.NewCommand("start", modules.Start))
	dis.AddHandler(handlers.NewCommand("run", modules.RunCode))
}
