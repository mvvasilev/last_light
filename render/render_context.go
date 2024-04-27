package render

import (
	"errors"
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

const (
	TERMINAL_SIZE_WIDTH      int         = 80
	TERMINAL_SIZE_HEIGHT     int         = 24
	DEFAULT_STYLE_BACKGROUND tcell.Color = tcell.ColorReset
	DEFAULT_STYLE_FOREGROUND tcell.Color = tcell.ColorReset
)

type Drawable interface {
	UniqueId() uuid.UUID
	Draw(v views.View)
}

func Multidraw(drawables ...Drawable) []Drawable {
	arr := make([]Drawable, 0)

	if drawables == nil {
		return arr
	}

	for _, d := range drawables {
		if d == nil {
			continue
		}

		arr = append(arr, d)
	}

	return arr
}

type RenderContext struct {
	screen tcell.Screen
	view   *views.ViewPort

	events    chan tcell.Event
	quit      chan struct{}
	drawables chan Drawable
}

func CreateRenderContext() (*RenderContext, error) {
	screen, sErr := tcell.NewScreen()

	if sErr != nil {
		log.Fatalf("%~v", sErr)
	}

	stopScreen := func() {
		screen.Fini()
	}

	if err := screen.Init(); err != nil {
		stopScreen()
		log.Fatal(err)
		return nil, err
	}

	width, height := screen.Size()

	if width < TERMINAL_SIZE_WIDTH || height < TERMINAL_SIZE_HEIGHT {
		stopScreen()
		log.Fatal("Unable to start; Terminal must be at least 80x24")
		return nil, errors.New("Terminal is undersized; must be at least 80x24")
	}

	view := views.NewViewPort(
		screen,
		(width/2)-(TERMINAL_SIZE_WIDTH/2),
		(height/2)-(TERMINAL_SIZE_HEIGHT/2),
		TERMINAL_SIZE_WIDTH,
		TERMINAL_SIZE_HEIGHT,
	)

	events := make(chan tcell.Event)
	quit := make(chan struct{})

	go screen.ChannelEvents(events, quit)

	context := new(RenderContext)

	context.screen = screen
	context.events = events
	context.quit = quit
	context.view = view
	context.drawables = make(chan Drawable)

	return context, nil
}

func (c *RenderContext) Stop() {
	c.screen.Fini()
}

func (c *RenderContext) CollectInputEvents() []*tcell.EventKey {
	events := make([]tcell.Event, len(c.events))

	select {
	case e := <-c.events:
		events = append(events, e)
	default:
	}

	inputEvents := make([]*tcell.EventKey, 0, len(events))

	for _, e := range events {
		switch ev := e.(type) {
		case *tcell.EventKey:
			inputEvents = append(inputEvents, ev)
		case *tcell.EventResize:
			c.onResize(ev)
		}
	}

	return inputEvents
}

func (c *RenderContext) DrawableQueue() chan Drawable {
	return c.drawables
}

func (c *RenderContext) onResize(ev *tcell.EventResize) {
	width, height := ev.Size()

	c.screen.Clear()

	c.view.Resize(
		(width/2)-(TERMINAL_SIZE_WIDTH/2),
		(height/2)-(TERMINAL_SIZE_HEIGHT/2),
		TERMINAL_SIZE_WIDTH,
		TERMINAL_SIZE_HEIGHT,
	)

	c.screen.Sync()
}

func (c *RenderContext) Draw(deltaTime int64, drawables []Drawable) {
	fps := 1_000_000 / deltaTime

	c.view.Clear()

	fpsText := CreateText(0, 0, 16, 1, fmt.Sprintf("%v FPS", fps), tcell.StyleDefault)

	for _, d := range drawables {
		d.Draw(c.view)
	}

	fpsText.Draw(c.view)

	c.screen.Show()
}
