package render

import (
	"errors"
	"log"
	"time"

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

type renderContext struct {
	screen       tcell.Screen
	view         *views.ViewPort
	defaultStyle tcell.Style

	events chan tcell.Event
	quit   chan struct{}

	lastRenderTime time.Time

	renderHandler func(view views.View, deltaTime int64)
	inputHandler  func(ev *tcell.EventKey)
}

func CreateRenderContext() (*renderContext, error) {
	s, err := tcell.NewScreen()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	stopScreen := func() {
		s.Fini()
	}

	if err := s.Init(); err != nil {
		stopScreen()
		log.Fatal(err)
		return nil, err
	}

	width, height := s.Size()

	if width < TERMINAL_SIZE_WIDTH || height < TERMINAL_SIZE_HEIGHT {
		stopScreen()
		log.Fatal("Unable to start; Terminal must be at least 80x24")
		return nil, errors.New("Terminal is undersized; must be at least 80x24")
	}

	view := views.NewViewPort(
		s,
		(width/2)-(TERMINAL_SIZE_WIDTH/2),
		(height/2)-(TERMINAL_SIZE_HEIGHT/2),
		TERMINAL_SIZE_WIDTH,
		TERMINAL_SIZE_HEIGHT,
	)

	defStyle := tcell.StyleDefault.Background(DEFAULT_STYLE_BACKGROUND).Foreground(DEFAULT_STYLE_FOREGROUND)

	events := make(chan tcell.Event)
	quit := make(chan struct{})

	go s.ChannelEvents(events, quit)

	context := new(renderContext)

	context.screen = s
	context.defaultStyle = defStyle
	context.events = events
	context.quit = quit
	context.view = view

	return context, nil
}

func (c *renderContext) Stop() {
	c.screen.Fini()
}

func (c *renderContext) HandleRender(renderHandler func(view views.View, deltaTime int64)) {
	c.renderHandler = renderHandler
}

func (c *renderContext) HandleInput(inputHandler func(ev *tcell.EventKey)) {
	c.inputHandler = inputHandler
}

func (c *renderContext) onResize(ev *tcell.EventResize) {
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

func (c *renderContext) BeginRendering() {
	c.lastRenderTime = time.Now()

	for {
		deltaTime := 1 + time.Since(c.lastRenderTime).Microseconds()
		c.lastRenderTime = time.Now()

		c.screen.Clear()

		c.renderHandler(c.view, deltaTime)

		c.screen.Show()

		select {
		case ev, ok := <-c.events:

			if !ok {
				break
			}

			switch ev := ev.(type) {
			case *tcell.EventResize:
				c.onResize(ev)
			case *tcell.EventKey:
				c.inputHandler(ev)
			}
		default:
		}
	}
}
