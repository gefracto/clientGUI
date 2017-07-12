package task2

import (
	"github.com/andlabs/ui"
	t2 "github.com/clientGUI/src/controller/task2"
	res "github.com/clientGUI/src/view/resultmsg"
)

func Task2() *ui.Group {
	task2 := ui.NewGroup("Task2")
	boxT2 := ui.NewVerticalBox()
	boxT2controls := ui.NewHorizontalBox()
	boxT2.Append(boxT2controls, false)
	task2.SetChild(boxT2)

	// Entries for Task2
	widthE1T2 := ui.NewEntry()
	widthE1T2.SetText("10")
	widthE1T2g := ui.NewGroup("width env1")
	widthE1T2g.SetChild(widthE1T2)

	heightE1T2 := ui.NewEntry()
	heightE1T2.SetText("15")
	heightE1T2g := ui.NewGroup("height env1")
	heightE1T2g.SetChild(heightE1T2)

	widthE2T2 := ui.NewEntry()
	widthE2T2.SetText("5")
	widthE2T2g := ui.NewGroup("width env2")
	widthE2T2g.SetChild(widthE2T2)

	heightE2T2 := ui.NewEntry()
	heightE2T2.SetText("40")
	heightE2T2g := ui.NewGroup("height env2")
	heightE2T2g.SetChild(heightE2T2)

	// Run button for Task2
	btnT2 := ui.NewButton("Run Task2")

	// Appendings Task2
	boxT2controls.Append(widthE1T2g, true)
	boxT2controls.Append(heightE1T2g, true)
	boxT2controls.Append(widthE2T2g, true)
	boxT2controls.Append(heightE2T2g, true)
	boxT2.Append(btnT2, false)

	btnT2.OnClicked(func(*ui.Button) {
		a, b := widthE1T2.Text(), heightE1T2.Text()
		c, d := widthE2T2.Text(), heightE2T2.Text()

		res.Res(t2.T2(a, b, c, d), "Task2")

	})

	return task2
}
