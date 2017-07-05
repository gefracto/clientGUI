package gui

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/andlabs/ui"
)

func Res(msg string, title string) {
	results := ui.NewWindow("RESULTS", 220, 220, false)
	results.OnClosing(func(*ui.Window) bool {
		results.Destroy()
		return false
	})
	results.SetTitle(title)

	root := ui.NewLabel("")
	root.SetText(msg)
	results.SetChild(root)
	results.Show()

}

func Draw() {
	err := ui.Main(func() {
		//RootElems
		boxROOT := ui.NewVerticalBox()
		btnAll := ui.NewButton("Run All")
		//Task1
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

		btnT1 := ui.NewButton("Run Task1")

		// Appendings Task1
		boxT1controls.Append(widthT1g, true)
		boxT1controls.Append(heightT1g, true)
		boxT1controls.Append(symbolT1g, true)
		boxT1.Append(btnT1, false)

		//Task2
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

		// Appendings Root
		boxROOT.Append(task1, true)
		boxROOT.Append(task2, true)
		boxROOT.Append(btnAll, true)

		window := ui.NewWindow("Client GUI", 400, 100, false)
		window.SetChild(boxROOT)

		// rules for buttons

		btnT1.OnClicked(func(*ui.Button) {
			w := widthT1.Text()
			h := heightT1.Text()
			s := symbolT1.Text()
			body := []byte("{\"width\":" + w + ",\"height\":" + h + ",\"symbol\":\"" + s + "\"}")
			resp, _ := http.Post("http://localhost:1111/task/1", "application/json", bytes.NewBuffer(body))
			r, _ := ioutil.ReadAll(resp.Body)

			Res(resp.Status+". Result1: \n-------\n"+string(r), "Task1")

		})

		btnT2.OnClicked(func(*ui.Button) {
			w1, h1 := widthE1T2.Text(), heightE1T2.Text()
			w2, h2 := widthE2T2.Text(), heightE2T2.Text()
			body := []byte("[{\"width\":" + w1 + ",\"height\":" + h1 + "},{\"width\":" + w2 + ",\"height\":" + h2 + "}]")
			resp, _ := http.Post("http://localhost:1111/task/2", "application/json", bytes.NewBuffer(body))
			r, _ := ioutil.ReadAll(resp.Body)
			Res(resp.Status+". Result2: \n-------\n"+string(r), "Task2")
		})

		// some cool shit

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
