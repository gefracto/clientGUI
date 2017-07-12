package task5

import (
	"github.com/andlabs/ui"
	t5 "github.com/clientGUI/src/controller/task5"
	res "github.com/clientGUI/src/view/resultmsg"
)

func Task5() *ui.Group {
	task5 := ui.NewGroup("Task5")
	boxT5 := ui.NewVerticalBox()
	boxT5controls := ui.NewHorizontalBox()
	boxT5.Append(boxT5controls, false)
	task5.SetChild(boxT5)

	minT5 := ui.NewEntry()
	minT5.SetText("0")
	minT5g := ui.NewGroup("Min")
	minT5g.SetChild(minT5)
	boxT5controls.Append(minT5g, true)

	maxT5 := ui.NewEntry()
	maxT5.SetText("999999")
	maxT5g := ui.NewGroup("Max")
	maxT5g.SetChild(maxT5)
	boxT5controls.Append(maxT5g, true)

	btnT5 := ui.NewButton("Run Task5")
	btnT5.OnClicked(func(*ui.Button) {
		res.Res(t5.T5(minT5.Text(), maxT5.Text()), "Task5")
	})

	boxT5.Append(btnT5, true)
	return task5
}
