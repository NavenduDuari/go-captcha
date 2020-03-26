package captcha

import (
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"log"

	"github.com/golang/freetype"
)

type Image struct {
	*image.RGBA
}

func newImage(w, h int) *Image {
	img := &Image{image.NewRGBA(image.Rect(0, 0, w, h))}
	return img
}

func (img *Image) fillBackground(srcImg image.Image) {
	draw.Draw(img, img.Bounds(), srcImg, image.ZP, draw.Over)
}

func (img *Image) drawString(str string) {
	ctx := freetype.NewContext()
	ctx.SetDst(img)
	ctx.SetClip(img.Bounds())
	ctx.SetSrc(image.NewUniform(color.Black))
	fontsize := float64(img.Bounds().Size().Y) * 0.4
	ctx.SetFontSize(fontsize)
	fontdata, err := ioutil.ReadFile("./comic.ttf")
	if err != nil {
		log.Fatal(err)
	}
	font, err := freetype.ParseFont(fontdata)
	if err != nil {
		log.Fatal(err)
	}
	ctx.SetFont(font)
	pt := freetype.Pt(0, int(fontsize))
	ctx.DrawString(str, pt)
}

func (img *Image) drawLine(x1, y1, x2, y2 int) {
	var dx, dy, e, slope int

	// Because drawing p1 -> p2 is equivalent to draw p2 -> p1,
	// I sort points in x-axis order to handle only half of possible cases.
	if x1 > x2 {
		x1, y1, x2, y2 = x2, y2, x1, y1
	}

	dx, dy = x2-x1, y2-y1
	// Because point is x-axis ordered, dx cannot be negative
	if dy < 0 {
		dy = -dy
	}

	switch {

	// Is line a point ?
	case x1 == x2 && y1 == y2:
		img.Set(x1, y1, color.Black)

	// Is line an horizontal ?
	case y1 == y2:
		for ; dx != 0; dx-- {
			img.Set(x1, y1, color.Black)
			x1++
		}
		img.Set(x1, y1, color.Black)

	// Is line a vertical ?
	case x1 == x2:
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for ; dy != 0; dy-- {
			img.Set(x1, y1, color.Black)
			y1++
		}
		img.Set(x1, y1, color.Black)

	// Is line a diagonal ?
	case dx == dy:
		if y1 < y2 {
			for ; dx != 0; dx-- {
				img.Set(x1, y1, color.Black)
				x1++
				y1++
			}
		} else {
			for ; dx != 0; dx-- {
				img.Set(x1, y1, color.Black)
				x1++
				y1--
			}
		}
		img.Set(x1, y1, color.Black)

	// wider than high ?
	case dx > dy:
		if y1 < y2 {
			// BresenhamDxXRYD(img, x1, y1, x2, y2, color.Black)
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				img.Set(x1, y1, color.Black)
				x1++
				e -= dy
				if e < 0 {
					y1++
					e += slope
				}
			}
		} else {
			// BresenhamDxXRYU(img, x1, y1, x2, y2, color.Black)
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				img.Set(x1, y1, color.Black)
				x1++
				e -= dy
				if e < 0 {
					y1--
					e += slope
				}
			}
		}
		img.Set(x2, y2, color.Black)

	// higher than wide.
	default:
		if y1 < y2 {
			// BresenhamDyXRYD(img, x1, y1, x2, y2, color.Black)
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				img.Set(x1, y1, color.Black)
				y1++
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		} else {
			// BresenhamDyXRYU(img, x1, y1, x2, y2, color.Black)
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				img.Set(x1, y1, color.Black)
				y1--
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		}
		img.Set(x2, y2, color.Black)
	}
}
