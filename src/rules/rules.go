package rules

type T1 struct {
	Width  int
	Height int
	Symbol string
}

func (t *T1) Check() bool {
	return t.Width > 0 && t.Height > 0
}

type T2 struct {
	Width  int
	Height int
}

func (t *T2) Check() bool {
	return t.Width > 0 && t.Height > 0
}
