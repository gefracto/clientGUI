package resultmsg

import "github.com/andlabs/ui"

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
