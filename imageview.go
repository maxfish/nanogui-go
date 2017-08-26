package nanogui

import (
	"github.com/gianpaolog/nanovgo"
	"github.com/goxjs/glfw"
	"math"
)

type ImageStretchMode int

const (
	StretchNone ImageStretchMode = iota
	StretchFit
)

type ImageView struct {
	WidgetImplement
	imgPosY float32
	imgPosX float32
	scale float32
	image  int
	stretch ImageStretchMode
}

func NewImageView(parent Widget, images ...int) *ImageView {
	var image int
	switch len(images) {
	case 0:
	case 1:
		image = images[0]
	default:
		panic("NewImageView can accept only one extra parameter (image)")
	}

	imageView := &ImageView{
		image:   image,
		stretch: StretchNone,
		scale:   1.0,
	}
	InitWidget(imageView, parent)

	return imageView
}

func (i *ImageView) Image() int {
	return i.image
}

func (i *ImageView) SetImage(image int) {
	i.image = image
	i.fit()
}

func (i *ImageView) StretchMode() ImageStretchMode {
	return i.stretch
}

func (i *ImageView) SetStretchMode(mode ImageStretchMode) {
	i.stretch = mode

	if mode == StretchFit {
		i.fit()
	} else {
		i.reset()
	}
}

func (i *ImageView) PreferredSize(self Widget, ctx *nanovgo.Context) (int, int) {
	if i.image == 0 {
		return 0, 0
	}
	w, h, _ := ctx.ImageSize(i.image)
	return w, h
}

func (i *ImageView) imageCoordinateAt(posX, posY int) (imgX, imgY float32){
	px := float32(posX)
	py := float32(posY)
	return (px-i.imgPosX)/i.scale, (py-i.imgPosY)/i.scale
}

func (i *ImageView) Draw(self Widget, ctx *nanovgo.Context) {
	if i.image == 0 {
		return
	}
	x := float32(i.x)
	y := float32(i.y)
	ow := float32(i.w)
	oh := float32(i.h)

	var w, h float32
	{
		iw, ih, _ := ctx.ImageSize(i.image)
		w = float32(iw)
		h = float32(ih)
	}

	ctx.Save()
	ctx.IntersectScissor(x,y,ow,oh)

	imgX := i.imgPosX + x
	imgY := i.imgPosY + y
	w *= i.scale
	h *= i.scale

	imgPaint := nanovgo.ImagePattern(imgX, imgY, w, h, 0, i.image, 1.0)

	ctx.BeginPath()
	ctx.Rect(imgX, imgY, w, h)
	ctx.SetFillPaint(imgPaint)
	ctx.Fill()

	i.drawImageBorder(ctx,imgX, imgY, w, h)
	if i.scale>20 {
		i.drawPixelGrid(ctx,imgX, imgY, w, h)
	}

	i.drawWidgetBorder(ctx, x, y, ow, oh)


	ctx.Restore()
}

func (i *ImageView) drawPixelGrid(ctx *nanovgo.Context, imgX, imgY, w, h float32) {
	scale := i.scale
	ctx.BeginPath()
	for cx := imgX; cx < w+imgX; cx +=scale {
		ctx.MoveTo(cx,imgY)
		ctx.LineTo(cx, imgY+h)
	}

	for cy := imgY; cy < h+imgY; cy +=scale {
		ctx.MoveTo(imgX,cy)
		ctx.LineTo(imgX+w, cy)
	}

	ctx.SetStrokeWidth(1.0)
	ctx.SetStrokeColor(nanovgo.MONO(255, 50))
	ctx.Stroke()
}

func (i *ImageView) drawImageBorder(ctx *nanovgo.Context, imgX, imgY, w, h float32) {
	ctx.BeginPath()
	ctx.SetStrokeWidth(1.0)
	ctx.Rect(imgX - 0.5, imgY - 0.5, w+1, h+1)
	ctx.SetStrokeColor(nanovgo.MONO(255, 255))
	ctx.Stroke()
}

func (i *ImageView) drawWidgetBorder(ctx *nanovgo.Context, x, y, w, h float32) {
	ctx.BeginPath()
	ctx.SetStrokeWidth(1.0)
	ctx.RoundedRect(x +0.5, y + 0.5, w - 1, h -1, 0 )
	ctx.SetStrokeColor(i.theme.WindowPopup)
	ctx.Stroke()

	ctx.BeginPath()
	ctx.RoundedRect(x +0.5, y + 0.5, w - 1, h -1, 0 )
	ctx.SetStrokeColor(i.theme.BorderDark)
	ctx.Stroke()
}

func (i *ImageView) String() string {
	return i.StringHelper("ImageView", "")
}

func (i *ImageView) ScrollEvent(self Widget, x, y, relX, relY int) bool {
	if i.stretch == StretchNone {
		return false
	}
	i.zoom(-relY, x, y)
	return true
}

func (i *ImageView) MouseDragEvent(self Widget, x, y, relX, relY, button int, modifier glfw.ModifierKey) bool{
	i.imgPosX += float32(relX)
	i.imgPosY += float32(relY)
	return true
}

func (i *ImageView) zoom (amount int, focusX, focusY int) {
	cx, cy := i.imageCoordinateAt(focusX, focusY)
	scaleFactor := math.Pow(1.1, float64(amount))
	i.scale = maxF(0.01, float32(scaleFactor) * i.scale)
	i.SetImageCoordinateAt(float32(focusX), float32(focusY), cx, cy)
}

func (i *ImageView) SetImageCoordinateAt(posX, posY, coordX, coordY float32) {
	i.imgPosX = posX - (coordX * i.scale)
	i.imgPosY = posY - (coordY * i.scale)
}

func (i *ImageView) center() {
	w := float32(i.w)
	h := float32(i.h)
	screen := i.FindWindow().Parent().(*Screen)
	iw, ih, _ := screen.context.ImageSize(i.image)
	sw := float32(iw) * i.scale
	sh := float32(ih) * i.scale
	i.imgPosX = (w-sw) / 2
	i.imgPosY = (h-sh) / 2
}

func (i *ImageView) fit() {
	w := float32(i.w)
	h := float32(i.h)
	if w==0||h==0 {
		return
	}
	screen := i.FindWindow().Parent().(*Screen)
	iw, ih, _ := screen.context.ImageSize(i.image)
	if iw==-1 && ih==-1 {
		return
	}
	i.scale = minF(w/float32(iw), h/float32(ih))
	i.center()
}

func (i *ImageView) reset() {
	i.scale = 1.0
	i.imgPosX = 0.0
	i.imgPosY = 0.0
}

func (i *ImageView) SetSize(w, h int) {
	i.w = w
	i.h = h

	if i.stretch == StretchFit {
		i.fit()
	}
}