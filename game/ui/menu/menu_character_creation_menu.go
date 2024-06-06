package menu

import (
	"fmt"
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/model"
	"mvvasilev/last_light/game/systems"
	"mvvasilev/last_light/game/ui"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type statSelection struct {
	stat            model.Stat
	label           *ui.UILabel
	plusButton      *ui.UILabel
	statNumberLabel *ui.UILabel
	minusButton     *ui.UILabel
}

type StatState struct {
	Stat  model.Stat
	Value int
}

type CharacterCreationMenuState struct {
	AvailablePoints    int
	CurrentHighlight   int
	Stats              []*StatState
	RandomizeCharacter func()
	StartGame          func()
}

type CharacterCreationMenu struct {
	window *ui.UIWindow

	availablePointsLabel *ui.UILabel

	state *CharacterCreationMenuState

	stats           []*statSelection
	randomizeButton *ui.UISimpleButton
	startGameButton *ui.UISimpleButton

	style tcell.Style
}

func CreateCharacterCreationMenu(state *CharacterCreationMenuState, style tcell.Style) *CharacterCreationMenu {
	ccm := &CharacterCreationMenu{
		state:  state,
		window: ui.CreateWindow(0, 0, engine.TERMINAL_SIZE_WIDTH, engine.TERMINAL_SIZE_HEIGHT, "Create Character", style),
		style:  style,
	}

	ccm.UpdateState(state)

	return ccm
}

func (ccm *CharacterCreationMenu) UpdateState(state *CharacterCreationMenuState) {
	ccm.state = state

	width, height := ccm.Size().WH()

	availablePointsText := fmt.Sprintf("Available Points: %v", state.AvailablePoints)

	ccm.availablePointsLabel = ui.CreateSingleLineUILabel(
		width/2-len(availablePointsText)/2,
		1,
		availablePointsText,
		ccm.style,
	)

	statPlaceholderText := fmt.Sprintf("%-12s [ < ] 00 [ > ]", "Placeholder")
	statX := width/2 - len(statPlaceholderText)/2

	statsArr := make([]*statSelection, 0, 4)

	for i, s := range state.Stats {

		labelStyle := ccm.style

		if i == state.CurrentHighlight {
			labelStyle = ccm.style.Attributes(tcell.AttrBold).Background(tcell.ColorWhite).Foreground(tcell.ColorBlack)
		}

		statsArr = append(
			statsArr,
			&statSelection{
				stat: s.Stat,
				label: ui.CreateSingleLineUILabel(
					statX,
					3+i,
					model.StatLongName(s.Stat),
					labelStyle,
				),
				minusButton: ui.CreateSingleLineUILabel(
					statX+12+3, // Account for highlighting brackets
					3+i,
					"<",
					ccm.style,
				),
				statNumberLabel: ui.CreateSingleLineUILabel(
					statX+12+3+1+3,
					3+i,
					fmt.Sprintf("%02v", s.Value),
					ccm.style,
				),
				plusButton: ui.CreateSingleLineUILabel(
					statX+12+3+1+3+2+3, // Account for highlighting brackets
					3+i,
					">",
					ccm.style,
				),
			},
		)
	}

	ccm.stats = statsArr

	randomizeLabel := "Randomize"

	ccm.randomizeButton = ui.CreateSimpleButton(
		width/2-len(randomizeLabel)/2,
		3+len(statsArr)+1,
		randomizeLabel,
		ccm.style,
		ccm.style.Attributes(tcell.AttrBold),
		state.RandomizeCharacter,
	)

	startGameLabel := "Start Game"

	ccm.startGameButton = ui.CreateSimpleButton(
		width/2-len(startGameLabel)/2,
		height-2,
		startGameLabel,
		ccm.style,
		ccm.style.Attributes(tcell.AttrBold),
		state.StartGame,
	)

	if state.CurrentHighlight == len(state.Stats) {
		ccm.randomizeButton.Highlight()
	}

	if state.CurrentHighlight == len(state.Stats)+1 {
		ccm.startGameButton.Highlight()
	}
}

func (ccm *CharacterCreationMenu) SelectHighlight() {
	if ccm.state.CurrentHighlight == len(ccm.state.Stats) {
		ccm.randomizeButton.Select()
	}

	if ccm.state.CurrentHighlight == len(ccm.state.Stats)+1 {
		ccm.startGameButton.Select()
	}
}

func (ccm *CharacterCreationMenu) MoveTo(x int, y int) {
}

func (ccm *CharacterCreationMenu) Position() engine.Position {
	return engine.PositionAt(0, 0)
}

func (ccm *CharacterCreationMenu) Size() engine.Size {
	return engine.SizeOf(engine.TERMINAL_SIZE_WIDTH, engine.TERMINAL_SIZE_HEIGHT)
}

func (ccm *CharacterCreationMenu) Input(inputAction systems.InputAction) {

}

func (ccm *CharacterCreationMenu) UniqueId() uuid.UUID {
	return uuid.New()
}

func (ccm *CharacterCreationMenu) Draw(v views.View) {
	ccm.window.Draw(v)

	ccm.availablePointsLabel.Draw(v)

	for _, val := range ccm.stats {
		val.label.Draw(v)
		val.minusButton.Draw(v)
		val.statNumberLabel.Draw(v)
		val.plusButton.Draw(v)
	}

	ccm.randomizeButton.Draw(v)

	ccm.startGameButton.Draw(v)
}
