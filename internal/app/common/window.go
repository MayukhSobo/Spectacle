package common

type Window struct {
	Height int
	Width  int
}

func NewWindow(height, width int) *Window {
	return &Window{
		Height: height,
		Width:  width,
	}
}
