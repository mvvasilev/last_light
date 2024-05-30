package player

import "mvvasilev/last_light/game/input"

func PlayerTurn(inputSystem *input.InputSystem) (complete, requeue bool) {
	requeue = true
	complete = false

	nextAction := inputSystem.NextAction()

	if nextAction == input.InputAction_None {
		return
	}

	return
}
