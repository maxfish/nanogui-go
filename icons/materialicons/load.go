package materialicons

import (
	"github.com/maxfish/vg4go-gl4"
)

func LoadFont(ctx *nanovgo.Context) {
	ctx.CreateFontFromMemory("materialicons", MustAsset("font/MaterialIcons-Regular.ttf"), 0)
}

func LoadFontAs(ctx *nanovgo.Context, name string) {
	ctx.CreateFontFromMemory(name, MustAsset("font/MaterialIcons-Regular.ttf"), 0)
}