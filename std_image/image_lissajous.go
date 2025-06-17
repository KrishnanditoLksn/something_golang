package std_image

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

const (
	whiteIndex = 0
	colorIdx   = 2
)

// define color palette
var palette = []color.Color{color.White, color.RGBA{0, 255, 0, 1}, color.RGBA{255, 165, 0, 1}}

func Lissajous(out io.Writer) {
	const (
		cycles  = 5     // complete x revolutions
		res     = 0.001 // angular res
		size    = 100   // image canvas covers
		nFrames = 64    // number animation frames
		delay   = 8     // delay between frames
	)

	freq := rand.Float64() * 3.0
	animation := gif.GIF{
		LoopCount: nFrames,
	}

	phase := 0.0 // phase diff

	for range nFrames {
		rectangle := image.Rect(0, 0, 2*size+1, 2*size+1)
		image := image.NewPaletted(rectangle, palette)
		for t := 0.0; t < cycles*2*3.14; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			image.SetColorIndex(size+int(x*size+1), size+int(y*size+0.5),
				colorIdx)
		}
		phase += 0.1
		//individual field of a struct can be accessed using dot notation
		animation.Delay = append(animation.Delay, delay)
		animation.Image = append(animation.Image, image)
	}
	gif.EncodeAll(out, &animation) // NOTE: ignoring encoding errors
}
