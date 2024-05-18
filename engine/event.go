package engine

import "github.com/gdamore/tcell/v2"

type Event interface {
}

type InputEvent struct {
	tcellEvent *tcell.EventKey
}

func (e *InputEvent) TcellEvent() *tcell.EventKey {
	return e.tcellEvent
}

type ResizeEvent struct {
	tcellEvent *tcell.EventResize
}

func (e *ResizeEvent) TcellEvent() *tcell.EventResize {
	return e.tcellEvent
}
