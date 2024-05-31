package state

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/input"
	"mvvasilev/last_light/game/rpg"
	"mvvasilev/last_light/game/turns"
	"mvvasilev/last_light/game/ui/menu"

	"github.com/gdamore/tcell/v2"
)

const (
	MinStatValue = 1
	MaxStatValue = 20
)

type CharacterCreationState struct {
	turnSystem  *turns.TurnSystem
	inputSystem *input.InputSystem

	startGame bool

	menuState *menu.CharacterCreationMenuState
	ccMenu    *menu.CharacterCreationMenu
}

func CreateCharacterCreationState(turnSystem *turns.TurnSystem, inputSystem *input.InputSystem) *CharacterCreationState {

	menuState := &menu.CharacterCreationMenuState{
		AvailablePoints:  21,
		CurrentHighlight: 0,
		Stats: []*menu.StatState{
			{
				Stat:  rpg.Stat_Attributes_Strength,
				Value: 1,
			},
			{
				Stat:  rpg.Stat_Attributes_Dexterity,
				Value: 1,
			},
			{
				Stat:  rpg.Stat_Attributes_Intelligence,
				Value: 1,
			},
			{
				Stat:  rpg.Stat_Attributes_Constitution,
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
		ccs.menuState.AvailablePoints = 21

		for _, s := range ccs.menuState.Stats {
			if ccs.menuState.AvailablePoints == 0 {
				break
			}

			limit := ccs.menuState.AvailablePoints

			if limit > 20 {
				limit = 20
			}

			s.Value = engine.RandInt(1, limit+1)
			ccs.menuState.AvailablePoints -= s.Value - 1
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

func (ccs *CharacterCreationState) InputContext() input.Context {
	return input.InputContext_Menu
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
		stats := map[rpg.Stat]int{}

		for _, s := range ccs.menuState.Stats {
			stats[s.Stat] = s.Value
		}

		return CreatePlayingState(ccs.turnSystem, ccs.inputSystem, stats)
	}

	action := ccs.inputSystem.NextAction()

	switch action {
	case input.InputAction_Menu_HighlightRight:
		ccs.IncreaseStatValue()
	case input.InputAction_Menu_HighlightLeft:
		ccs.DecreaseStatValue()
	case input.InputAction_Menu_HighlightDown:
		ccs.menuState.CurrentHighlight++
	case input.InputAction_Menu_HighlightUp:
		ccs.menuState.CurrentHighlight--
	case input.InputAction_Menu_Select:
		ccs.ccMenu.SelectHighlight()
	}

	ccs.ccMenu.UpdateState(ccs.menuState)

	return ccs
}

func (ccs *CharacterCreationState) CollectDrawables() []engine.Drawable {
	return []engine.Drawable{ccs.ccMenu}
}
