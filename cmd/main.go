package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/NavenduDuari/go-captcha"
)

func main() {
	cap := captcha.New()

	img, str := cap.Create()
	captchaImage, _ := os.Create("test.png")
	png.Encode(captchaImage, img)
	fmt.Println(str)

}
