package task3

import (
	"fmt"

	"github.com/andlabs/ui"
	t3 "github.com/clientGUI/src/controller/task3"
	res "github.com/clientGUI/src/view/resultmsg"
)

var Task3entries []*ui.Entry
var BoxT3controls []*ui.Box

func collectTriangles(entries []*ui.Entry) (names, a, b, c []string) {
	var values []string
	for i := 0; i < 4; i++ {
		for j := i; j < len(entries); j += 4 {
			values = append(values, entries[j].Text())
		}
	}
	step := len(values) / 4
	names = values[:step]
	a = values[step : step*2]
	b = values[step*2 : step*3]
	c = values[step*3 : step*4]

	return names, a, b, c
}

func addTriangle(boxes []*ui.Box, entries []*ui.Entry) ([]*ui.Box, []*ui.Entry, *ui.Box) {
	verts := ui.NewEntry()
	verts.SetText("ABC")
	a := ui.NewEntry()
	a.SetText("10")
	b := ui.NewEntry()
	b.SetText("15")
	c := ui.NewEntry()
	c.SetText("18")
	box := ui.NewHorizontalBox()
	box.Append(verts, true)
	box.Append(a, true)
	box.Append(b, true)
	box.Append(c, true)
	boxes = append(boxes, box)
	entries = append(entries, verts, a, b, c)
	return boxes, entries, box

}

func Task3(boxes []*ui.Box, entries []*ui.Entry) *ui.Group {
	task3 := ui.NewGroup("Task3")
	boxT3 := ui.NewVerticalBox()
	boxT3controls := ui.NewVerticalBox()
	boxT3.Append(boxT3controls, true)
	vert := ui.NewEntry()
	vert.SetText("ABC")
	a := ui.NewEntry()
	a.SetText("10")
	b := ui.NewEntry()
	b.SetText("15")
	c := ui.NewEntry()
	c.SetText("18")
	boxesCount := 1

	entries = append(entries, vert, a, b, c)
	contr := ui.NewHorizontalBox()
	contr.Append(vert, true)
	contr.Append(a, true)
	contr.Append(b, true)
	contr.Append(c, true)
	boxes = append(boxes, contr)

	boxT3controls.Append(contr, true)

	task3.SetChild(boxT3)

	btnAdd := ui.NewButton("+ MORE +")
	btnT3 := ui.NewButton("Run Task3")
	btnRmv := ui.NewButton("- LESS -")
	btnBox := ui.NewHorizontalBox()
	btnBox.Append(btnAdd, false)
	btnBox.Append(btnT3, true)
	btnBox.Append(btnRmv, false)
	var bo *ui.Box

	btnAdd.OnClicked(func(*ui.Button) {
		boxes, entries, bo = addTriangle(boxes, entries)
		boxT3controls.Append(bo, true)
		boxesCount = boxesCount + 1

	})

	btnRmv.OnClicked(func(*ui.Button) {
		if boxesCount > 1 {
			boxT3controls.Delete(boxesCount - 1)
			boxes = boxes[:boxesCount]
			entries = entries[:boxesCount*4-4]
			boxesCount = boxesCount - 1
		}
	})
	btnT3.OnClicked(func(*ui.Button) {
		for _, v := range entries {
			fmt.Println(v.Text() + "\n---\n")
		}
		fmt.Println("END")
		fmt.Println(collectTriangles(entries))
		res.Res(t3.T3(collectTriangles(entries)), "Task3")
	})

	boxT3.Append(btnBox, false)

	return task3
}
