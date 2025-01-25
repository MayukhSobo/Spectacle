package common

// Window represents terminal dimensions for UI layout calculations.
// Tracks height and width to enable responsive component rendering.
type Window struct {
	Height int
	Width  int
}

// NewWindow creates a new Window instance with given dimensions.
// Used to initialize window size tracking for UI components.
func NewWindow(height, width int) *Window {
	return &Window{
		Height: height,
		Width:  width,
	}
}
