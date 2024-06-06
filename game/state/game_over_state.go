package state

import (
	"fmt"
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/systems"
	"mvvasilev/last_light/game/ui"

	"github.com/gdamore/tcell/v2"
)

type GameOverState struct {
	inputSystem *systems.InputSystem

	gameOverTitle *engine.Raw
	deathText     *ui.UILabel

	backToMainMenuBtn *ui.UISimpleButton

	returnToMainMenu bool
}

func CreateGameOverState(inputSystem *systems.InputSystem) *GameOverState {
	gos := &GameOverState{
		inputSystem: inputSystem,
		gameOverTitle: engine.CreateRawDrawable(
			14, 1, tcell.StyleDefault.Attributes(tcell.AttrBold).Foreground(tcell.ColorYellow),
			"_____                        _____                 ",
			"|  __ \\                      |  _  |               ",
			"| |  \\/ __ _ _ __ ___   ___  | | | |_   _____ _ __ ",
			"| | __ / _` | '_ ` _ \\ / _ \\ | | | \\ \\ / / _ \\ '__|",
			"| |_\\ \\ (_| | | | | | |  __/ \\ \\_/ /\\ V /  __/ |   ",
			" \\____/\\__,_|_| |_| |_|\\___|  \\___/  \\_/ \\___|_|   ",
		),
		deathText: ui.CreateUILabel(
			14, 8, 51, 5,
			fmt.Sprintf(
				"For all your efforts, your endeavour was ultimately cut short. "+
					"You have been left bleeding out on the dungeon floor. Your remains "+
					"will serve as a warning to future seekers of the last light.",
			),
			tcell.StyleDefault,
		),
	}

	gos.backToMainMenuBtn = ui.CreateSimpleButton(
		engine.TERMINAL_SIZE_WIDTH/2-len("Back to Main Menu")/2,
		16,
		"Back to Main Menu",
		tcell.StyleDefault,
		tcell.StyleDefault.Attributes(tcell.AttrBold),
		func() {
			gos.returnToMainMenu = true
		},
	)

	gos.backToMainMenuBtn.Highlight()

	return gos
}

func (gos *GameOverState) InputContext() systems.InputContext {
	return systems.InputContext_Menu
}

func (gos *GameOverState) OnTick(dt int64) GameState {
	if gos.inputSystem.NextAction() == systems.InputAction_Menu_Select {
		gos.backToMainMenuBtn.Select()
	}

	if gos.returnToMainMenu {
		return CreateMainMenuState(systems.CreateTurnSystem(), gos.inputSystem)
	}

	return gos
}

func (gos *GameOverState) CollectDrawables() []engine.Drawable {
	return []engine.Drawable{
		gos.gameOverTitle,
		gos.deathText,
		gos.backToMainMenuBtn,
	}
}
