package task1

import (
	"fmt"

	"github.com/andlabs/ui"
	t1 "github.com/clientGUI/src/controller/task1"
	res "github.com/clientGUI/src/view/resultmsg"
)

var Width string

func Task1() *ui.Group {
	task1 := ui.NewGroup("Task1")

	boxT1 := ui.NewVerticalBox()
	boxT1controls := ui.NewHorizontalBox()
	boxT1.Append(boxT1controls, false)

	task1.SetChild(boxT1)

	widthT1 := ui.NewEntry()
	widthT1.SetText("8")
	widthT1g := ui.NewGroup("width")
	widthT1g.SetChild(widthT1)

	heightT1 := ui.NewEntry()
	heightT1.SetText("8")
	heightT1g := ui.NewGroup("height")
	heightT1g.SetChild(heightT1)

	symbolT1 := ui.NewEntry()
	symbolT1.SetText(" * ")
	symbolT1g := ui.NewGroup("symbol")
	symbolT1g.SetChild(symbolT1)

	BtnT1 := ui.NewButton("Run Task1")

	// Appendings Task1
	boxT1controls.Append(widthT1g, true)
	boxT1controls.Append(heightT1g, true)
	boxT1controls.Append(symbolT1g, true)
	boxT1.Append(BtnT1, false)
	//	Width, Height, Symbol = widthT1.Text(), heightT1.Text(), symbolT1.Text()
	widthT1.OnChanged(func(*ui.Entry) {
		Width = widthT1.Text()
		fmt.Println(Width)
	})
	BtnT1.OnClicked(func(*ui.Button) {
		res.Res(t1.T1(widthT1.Text(), heightT1.Text(), symbolT1.Text()), "Task1")
	})

	return task1
}
