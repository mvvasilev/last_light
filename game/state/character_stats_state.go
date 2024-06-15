package state

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/model"
	"mvvasilev/last_light/game/systems"
	"mvvasilev/last_light/game/ui"
)

type CharacterStatsState struct {
	window *ui.UIWindow

	player *model.Player
}

func (css *CharacterStatsState) InputContext() systems.InputContext {
	return systems.InputContext_Menu
}

func (css *CharacterStatsState) OnTick(dt int64) GameState {
	return css
}

func (css *CharacterStatsState) CollectDrawables() []engine.Drawable {
	return []engine.Drawable{}
}
