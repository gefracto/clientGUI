package task7

import (
	"github.com/andlabs/ui"
	t7 "github.com/clientGUI/src/controller/task7"
	res "github.com/clientGUI/src/view/resultmsg"
)

func Task7() *ui.Group {
	task7 := ui.NewGroup("Task7")
	boxT7 := ui.NewVerticalBox()
	boxT7controls := ui.NewHorizontalBox()
	boxT7.Append(boxT7controls, false)
	task7.SetChild(boxT7)

	fileT7 := ui.NewEntry()
	fileT7.SetText("context")
	fileT7g := ui.NewGroup("File name")
	fileT7g.SetChild(fileT7)
	boxT7controls.Append(fileT7g, true)
	btnT7 := ui.NewButton("Run Task7")
	btnT7.OnClicked(func(*ui.Button) {
		res.Res(t7.T7(fileT7.Text()), "Task7")
	})

	boxT7.Append(btnT7, true)
	return task7
}
