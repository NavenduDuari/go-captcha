package captcha

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand"
	"time"
)

type Captcha struct {
	size image.Point
}

func New() *Captcha {
	c := &Captcha{
		size: image.Point{128, 64},
	}
	return c
}

func (c *Captcha) Create() (*Image, string) {
	dst := newImage(c.size.X, c.size.Y)
	str := string(c.randStr(6))
	c.drawString(dst, str)
	c.drawNoises(dst)
	return dst, str
}

var letters = []byte("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (c *Captcha) randStr(size int) []byte {
	result := make([]byte, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return result
}

func (c *Captcha) drawString(img *Image, str string) {

	tmpImg := newImage(c.size.X, c.size.Y)

	strImg := newImage(c.size.X, c.size.Y)

	strImg.fillBackground(image.NewUniform(color.White))
	tmpImg.fillBackground(image.NewUniform(color.White))

	strImg.drawString(str)

	draw.Draw(tmpImg, image.Rect(10, 15, strImg.Bounds().Size().X, strImg.Bounds().Size().Y), strImg, image.ZP, draw.Over)
	draw.Draw(img, tmpImg.Bounds(), tmpImg, image.ZP, draw.Over)
}

func (c *Captcha) drawNoises(img *Image) {
	rand.Seed(time.Now().UnixNano())
	size := img.Bounds().Size()
	for i := 1; i <= 5; i++ {
		x := i * rand.Intn(size.X/5)
		y := i * rand.Intn(size.Y/5)
		dx := rand.Intn(size.X)
		dy := rand.Intn(size.Y)
		img.drawLine(x, y, x+dx, y+dy)
	}
}
