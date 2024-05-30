package engine

import (
	"time"
)

type GameEvent struct {
	time     time.Time
	contents string
}

func (ge *GameEvent) Time() time.Time {
	return ge.time
}

func (ge *GameEvent) Contents() string {
	return ge.contents
}

type GameEventLog struct {
	logs []*GameEvent

	maxSize int
}

func CreateGameEventLog(maxSize int) *GameEventLog {
	return &GameEventLog{
		logs:    make([]*GameEvent, 0, 10),
		maxSize: maxSize,
	}
}

func (log *GameEventLog) Log(contents string) {
	log.logs = append(log.logs, &GameEvent{
		time:     time.Now(),
		contents: contents,
	})

	if len(log.logs) > log.maxSize {
		log.logs = log.logs[1:len(log.logs)]
	}
}

func (log *GameEventLog) Tail(n int) []*GameEvent {
	if n > len(log.logs) {
		return log.logs
	}

	return log.logs[len(log.logs)-n : len(log.logs)]
}
