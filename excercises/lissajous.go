package excercises

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.Black, color.RGBA{60, 255, 0, 255},
	color.RGBA{255, 0, 255, 255}, color.RGBA{238, 255, 0, 255}}

func Lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for range nframes {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Tan(t*freq + phase)
			rndColor := uint8(rand.Intn(4))

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				rndColor)
		}
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Cos(t)
			y := math.Tan(t*freq + phase)
			rndColor := uint8(rand.Intn(4))

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				rndColor)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
}
