package task6

import (
	"github.com/andlabs/ui"
	t6 "github.com/clientGUI/src/controller/task6"
	res "github.com/clientGUI/src/view/resultmsg"
)

func Task6() *ui.Group {
	task6 := ui.NewGroup("Task6")
	boxT6 := ui.NewVerticalBox()
	boxT6controls := ui.NewHorizontalBox()
	boxT6.Append(boxT6controls, false)
	task6.SetChild(boxT6)

	lengthT6 := ui.NewEntry()
	lengthT6.SetText("10")
	lengthT6g := ui.NewGroup("length")
	lengthT6g.SetChild(lengthT6)
	boxT6controls.Append(lengthT6g, true)

	valueT6 := ui.NewEntry()
	valueT6.SetText("625")
	valueT6g := ui.NewGroup("value")
	valueT6g.SetChild(valueT6)
	boxT6controls.Append(valueT6g, true)

	btnT6 := ui.NewButton("Run Task6")
	btnT6.OnClicked(func(*ui.Button) {
		res.Res(t6.T6(lengthT6.Text(), valueT6.Text()), "Task6")
	})

	boxT6.Append(btnT6, true)
	return task6
}
