// +build !js

package main

import (
	"fmt"
	"github.com/shibukawa/glfw"
	"github.com/gianpaolog/nanogui-go"
	"github.com/shibukawa/nanovgo"
	"io/ioutil"
	"path"
	_ "math"
	"strconv"
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
	window := nanogui.NewWindow(a.screen, "Vertical Scroll Panel")
	window.SetPosition(100,50)
	window.SetLayout(nanogui.NewBoxLayout(nanogui.Vertical, nanogui.Middle, 10, 20))

	dataListView := nanogui.NewWidget(window)
	dataListView.SetLayout(nanogui.NewBoxLayout(nanogui.Vertical, nanogui.Fill))

	//nanogui.NewButton(dataListView, "New Button!")

	table := nanogui.NewGridLayout(nanogui.Horizontal, 4, nanogui.Fill, 3, 6)

	scrollPanel := nanogui.NewVScrollPanel(dataListView)
	scrollPanel.SetFixedHeight(400)
	scrollPanel.SetLayout(nanogui.NewBoxLayout(nanogui.Vertical, nanogui.Fill))


	container := nanogui.NewWidget(scrollPanel)
	//container.SetFixedSize(200,400)
	container.SetLayout(table)

	for i := 0; i < 16; i++ {
		label := nanogui.NewLabel(container, "A label " + strconv.Itoa(i+1))
		label.SetFixedWidth(100)
		label.SetFont("sans-bold")

		combo := nanogui.NewComboBox(container,[]string{"Combo 1 Item", "Combo Item 2", "Combo Item 3"})
		combo.SetEnabled(true)
		combo.SetFixedWidth(100)
		combo.SetFontSize(12)

		tb := nanogui.NewTextBox(container, "A text box")
		tb.SetEditable(true)
		tb.SetFixedWidth(100)
		tb.SetFont("sans-bold")

		ck := nanogui.NewCheckBox(container,"CheckBox ctrl")
		ck.SetChecked(i%2 != 0)
		ck.SetFixedWidth(100)
		ck.SetFontSize(12)

	}


	//nanogui.SetDebug(true)
	a.screen.PerformLayout()
	a.screen.DebugPrint()
}

func main() {
	nanogui.Init()
	//nanogui.SetDebug(true)
	app := Application{}
	app.init()
	app.screen.DrawAll()
	app.screen.SetVisible(true)
	nanogui.MainLoop()
}

func loadImageDirectory(ctx *nanovgo.Context, dir string) []nanogui.Image {
	var images []nanogui.Image
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(fmt.Sprintf("loadImageDirectory: read error %v\n", err))
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		ext := path.Ext(file.Name())
		if ext != ".png" {
			continue
		}
		fullPath := path.Join(dir, file.Name())
		img := ctx.CreateImage(fullPath, 0)
		if img == 0 {
			panic("Could not open image data!")
		}
		images = append(images, nanogui.Image{
			ImageID: img,
			Name:    fullPath[:len(fullPath)-4],
		})
	}
	return images
}
