package systems

import "mvvasilev/last_light/engine"

type turn struct {
	cost   int
	action func() (complete, requeue bool)
}

type TurnSystem struct {
	turnQueue *engine.PriorityQueue[*turn]

	paused bool
}

func CreateTurnSystem() *TurnSystem {
	return &TurnSystem{
		turnQueue: engine.CreatePriorityQueue[*turn](),
		paused:    false,
	}
}

func (ts *TurnSystem) NextTurn() {
	if ts.paused {
		return
	}

	turnCost, turn := ts.turnQueue.Peek()

	if turn == nil {
		return
	}

	complete, requeue := turn.action()

	// If the action isn't complete, we'll re-do it again next time
	if !complete {
		return
	}

	ts.turnQueue.Dequeue()

	ts.turnQueue.AdjustPriorities(-turnCost)

	if requeue {
		ts.turnQueue.Enqueue(turn.cost, turn)
	}
}

func (ts *TurnSystem) Schedule(cost int, action func() (complete, requeue bool)) {
	ts.turnQueue.Enqueue(cost, &turn{
		cost:   cost,
		action: action,
	})
}

func (ts *TurnSystem) Clear() {
	ts.turnQueue.Clear()
}
