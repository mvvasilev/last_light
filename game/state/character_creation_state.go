package state

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/model"
	"mvvasilev/last_light/game/systems"
	"mvvasilev/last_light/game/ui/menu"

	"github.com/gdamore/tcell/v2"
)

const (
	MinStatValue = 1
	MaxStatValue = 20
)

type CharacterCreationState struct {
	turnSystem  *systems.TurnSystem
	inputSystem *systems.InputSystem

	startGame bool

	menuState *menu.CharacterCreationMenuState
	ccMenu    *menu.CharacterCreationMenu
}

func CreateCharacterCreationState(turnSystem *systems.TurnSystem, inputSystem *systems.InputSystem) *CharacterCreationState {

	menuState := &menu.CharacterCreationMenuState{
		AvailablePoints:  21,
		CurrentHighlight: 0,
		Stats: []*menu.StatState{
			{
				Stat:  model.Stat_Attributes_Strength,
				Value: 1,
			},
			{
				Stat:  model.Stat_Attributes_Dexterity,
				Value: 1,
			},
			{
				Stat:  model.Stat_Attributes_Intelligence,
				Value: 1,
			},
			{
				Stat:  model.Stat_Attributes_Constitution,
				Value: 1,
			},
		},
	}

	ccs := &CharacterCreationState{
		turnSystem:  turnSystem,
		inputSystem: inputSystem,
		menuState:   menuState,
		ccMenu:      menu.CreateCharacterCreationMenu(menuState, tcell.StyleDefault),
	}

	ccs.menuState.RandomizeCharacter = func() {
		stats := model.RandomStats(21, 1, 20, []model.Stat{
			model.Stat_Attributes_Strength,
			model.Stat_Attributes_Constitution,
			model.Stat_Attributes_Intelligence,
			model.Stat_Attributes_Dexterity,
		})

		ccs.menuState.AvailablePoints = 0
		ccs.menuState.Stats = []*menu.StatState{}

		for k, v := range stats {
			ccs.menuState.Stats = append(ccs.menuState.Stats, &menu.StatState{
				Stat:  k,
				Value: v,
			})
		}

		ccs.ccMenu.UpdateState(ccs.menuState)
	}

	ccs.menuState.StartGame = func() {
		if ccs.menuState.AvailablePoints > 0 {
			return
		}

		ccs.startGame = true
	}

	return ccs
}

func (ccs *CharacterCreationState) InputContext() systems.InputContext {
	return systems.InputContext_Menu
}

func (ccs *CharacterCreationState) IncreaseStatValue() {
	// If there are no points to allocate, stop
	if ccs.menuState.AvailablePoints == 0 {
		return
	}

	// If the current highlight is beyond the array range, stop
	if ccs.menuState.CurrentHighlight < 0 || ccs.menuState.CurrentHighlight >= len(ccs.menuState.Stats) {
		return
	}

	// If the allowed max state value has already been reached
	if ccs.menuState.Stats[ccs.menuState.CurrentHighlight].Value+1 > MaxStatValue {
		return
	}

	ccs.menuState.Stats[ccs.menuState.CurrentHighlight].Value++
	ccs.menuState.AvailablePoints--
}

func (ccs *CharacterCreationState) DecreaseStatValue() {
	// If the current highlight is beyond the array range, stop
	if ccs.menuState.CurrentHighlight < 0 || ccs.menuState.CurrentHighlight >= len(ccs.menuState.Stats) {
		return
	}

	// If the allowed min state value has already been reached
	if ccs.menuState.Stats[ccs.menuState.CurrentHighlight].Value-1 < MinStatValue {
		return
	}

	ccs.menuState.Stats[ccs.menuState.CurrentHighlight].Value--
	ccs.menuState.AvailablePoints++
}

func (ccs *CharacterCreationState) OnTick(dt int64) GameState {
	if ccs.startGame {
		stats := map[model.Stat]int{}

		for _, s := range ccs.menuState.Stats {
			stats[s.Stat] = s.Value
		}

		return CreatePlayingState(ccs.turnSystem, ccs.inputSystem, stats)
	}

	action := ccs.inputSystem.NextAction()

	switch action {
	case systems.InputAction_Menu_HighlightRight:
		ccs.IncreaseStatValue()
	case systems.InputAction_Menu_HighlightLeft:
		ccs.DecreaseStatValue()
	case systems.InputAction_Menu_HighlightDown:
		if ccs.menuState.CurrentHighlight > len(ccs.menuState.Stats) {
			break
		}

		ccs.menuState.CurrentHighlight++
	case systems.InputAction_Menu_HighlightUp:
		if ccs.menuState.CurrentHighlight == 0 {
			break
		}

		ccs.menuState.CurrentHighlight--
	case systems.InputAction_Menu_Select:
		ccs.ccMenu.SelectHighlight()
	}

	ccs.ccMenu.UpdateState(ccs.menuState)

	return ccs
}

func (ccs *CharacterCreationState) CollectDrawables() []engine.Drawable {
	return []engine.Drawable{ccs.ccMenu}
}
