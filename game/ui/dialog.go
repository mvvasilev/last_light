package ui

import (
	"mvvasilev/last_light/engine"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type UIDialog struct {
	window *UIWindow
	prompt *UILabel

	yesBtn *UISimpleButton
	noBtn  *UISimpleButton
}

func CreateYesNoDialog(title string, prompt string, yesText string, noText string, lineWidth int, yesAction func(), noAction func()) *UIDialog {
	d := new(UIDialog)

	numLines := len(prompt) / lineWidth
	winWidth := lineWidth + 2
	winHeight := numLines + 2
	winX := engine.TERMINAL_SIZE_WIDTH - winWidth/2
	winY := engine.TERMINAL_SIZE_HEIGHT - winHeight/2

	d.window = CreateWindow(winX, winY, winWidth, winHeight, title, tcell.StyleDefault)
	d.prompt = CreateUILabel(winX+1, winY+1, lineWidth, numLines, prompt, tcell.StyleDefault)

	yesBtnLength := len(yesText) + 4
	noBtnLength := len(noText) + 4

	yesBtnPosX := winX + winWidth/4 - yesBtnLength/2

	d.yesBtn = CreateSimpleButton(yesBtnPosX, winY+winHeight-1, yesText, tcell.StyleDefault, tcell.StyleDefault.Attributes(tcell.AttrBold), yesAction)

	noBtnPosX := winX + 3*winWidth/4 - noBtnLength/2

	d.noBtn = CreateSimpleButton(noBtnPosX, winY+winHeight-2, noText, tcell.StyleDefault, tcell.StyleDefault.Attributes(tcell.AttrBold), noAction)

	d.yesBtn.Highlight()

	return d
}

func CreateOkDialog(title string, prompt string, okText string, lineWidth int, okAction func()) *UIDialog {
	d := new(UIDialog)

	numLines := len(prompt) / lineWidth
	winWidth := lineWidth + 2
	winHeight := numLines + 5
	winX := engine.TERMINAL_SIZE_WIDTH/2 - winWidth/2
	winY := engine.TERMINAL_SIZE_HEIGHT/2 - winHeight/2

	d.window = CreateWindow(winX, winY, winWidth, winHeight, title, tcell.StyleDefault)
	d.prompt = CreateUILabel(winX+1, winY+1, lineWidth, numLines, prompt, tcell.StyleDefault)

	yesBtnLength := len(okText) + 4

	yesBtnPosX := winX + winWidth/2 - yesBtnLength/2

	d.yesBtn = CreateSimpleButton(yesBtnPosX, winY+winHeight-2, okText, tcell.StyleDefault, tcell.StyleDefault.Attributes(tcell.AttrBold), okAction)
	d.yesBtn.Highlight()

	return d
}

func (d *UIDialog) Select() {
	if d.yesBtn.IsHighlighted() {
		d.yesBtn.Select()
	} else if d.noBtn != nil && d.noBtn.IsHighlighted() {
		d.noBtn.Select()
	}
}

func (d *UIDialog) OnSelect(f func()) {
	d.yesBtn.OnSelect(f)
}

func (d *UIDialog) MoveTo(x int, y int) {

}

func (d *UIDialog) Position() engine.Position {
	return d.window.Position()
}

func (d *UIDialog) Size() engine.Size {
	return d.window.Size()
}

func (d *UIDialog) Input(e *tcell.EventKey) {
	if e.Key() == tcell.KeyLeft {
		if !d.yesBtn.IsHighlighted() {
			d.noBtn.Unhighlight()
			d.yesBtn.Highlight()
		}
	} else if e.Key() == tcell.KeyRight {
		if d.noBtn == nil {
			return
		}

		if !d.noBtn.IsHighlighted() {
			d.noBtn.Highlight()
			d.yesBtn.Unhighlight()
		}
	}
}

func (d *UIDialog) UniqueId() uuid.UUID {
	return d.window.UniqueId()
}

func (d *UIDialog) Draw(v views.View) {
	d.window.Draw(v)
	d.prompt.Draw(v)
	d.yesBtn.Draw(v)

	if d.noBtn != nil {
		d.noBtn.Draw(v)
	}
}
