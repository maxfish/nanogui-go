// +build !js

package main

import (
	"github.com/goxjs/glfw"
	"github.com/gianpaolog/nanogui-go"
)

type Application struct {
	screen   *nanogui.Screen
	progress *nanogui.ProgressBar
	shader   *nanogui.GLShader
}

func (a *Application) init() {
	glfw.WindowHint(glfw.Samples, 4)
	a.screen = nanogui.NewScreen(1024, 768, "NanoGUI.Go Test", true, false)

	//images := loadImageDirectory(a.screen.NVGContext(), "icons")
	window := nanogui.NewWindow(a.screen, "Tab Widget")
	window.SetPosition(100,50)
	window.SetFixedSize(400, 300)
	window.SetLayout(nanogui.NewBoxLayout(nanogui.Vertical, nanogui.Middle, 10, 20))

	tab := nanogui.NewTabHeader(window)

	tab.AddTab(0,"First")
	tab.AddTab(0,"Second")

	//nanogui.SetDebug(true)
	a.screen.PerformLayout()
	a.screen.DebugPrint()
}

func main() {
	nanogui.Init()
	nanogui.SetDebug(true)
	app := Application{}
	app.init()
	app.screen.DrawAll()
	app.screen.SetVisible(true)
	nanogui.MainLoop()
}


