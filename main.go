package main

import (
	"gioui.org/ui/app"
)

func main() {
	go func() {
		w := app.NewWindow(nil)
		for range w.Events() {
		}
	}()
	app.Main()
}
