package task4

import (
	"github.com/andlabs/ui"
	t4 "github.com/clientGUI/src/controller/task4"
	res "github.com/clientGUI/src/view/resultmsg"
)

func Task4() *ui.Group {
	task4 := ui.NewGroup("Task4")
	boxT4 := ui.NewVerticalBox()
	boxT4controls := ui.NewHorizontalBox()
	boxT4.Append(boxT4controls, false)
	task4.SetChild(boxT4)

	numberT4 := ui.NewEntry()
	numberT4.SetText("1424156513")
	numberT4g := ui.NewGroup("Number")
	numberT4g.SetChild(numberT4)
	boxT4controls.Append(numberT4g, true)
	btnT4 := ui.NewButton("Run Task4")

	btnT4.OnClicked(func(*ui.Button) {
		res.Res(t4.T4(numberT4.Text()), "Task4")
	})
	boxT4.Append(btnT4, true)

	return task4
}
