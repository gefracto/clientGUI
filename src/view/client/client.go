package client

import (
	"github.com/andlabs/ui"
	t1 "github.com/clientGUI/src/view/task1"
	t2 "github.com/clientGUI/src/view/task2"
	t3 "github.com/clientGUI/src/view/task3"
	t4 "github.com/clientGUI/src/view/task4"
	t5 "github.com/clientGUI/src/view/task5"
	t6 "github.com/clientGUI/src/view/task6"
	t7 "github.com/clientGUI/src/view/task7"
)

func ShowClient() {
	err := ui.Main(func() {
		//RootElems
		boxROOT := ui.NewVerticalBox()
		btnAll := ui.NewButton("Run All")

		btnAll.OnClicked(func(*ui.Button) {
			//			runall.Runall()
		})

		// Appendings Root
		boxROOT.Append(t1.Task1(), false)
		boxROOT.Append(t2.Task2(), false)
		boxROOT.Append(t3.Task3(t3.BoxT3controls, t3.Task3entries), false)
		boxROOT.Append(t4.Task4(), false)
		boxROOT.Append(t5.Task5(), false)
		boxROOT.Append(t6.Task6(), false)
		boxROOT.Append(t7.Task7(), false)
		boxROOT.Append(btnAll, false)
		window := ui.NewWindow("Client GUI", 300, 100, false)
		window.SetChild(boxROOT)

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
