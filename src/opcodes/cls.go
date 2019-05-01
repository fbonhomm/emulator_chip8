package CLS

// CLS 00E0 - Clear the display.
func CLS(s *screen) {
	s.RemoveScreen()
}
