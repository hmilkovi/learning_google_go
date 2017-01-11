 // Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"time"
	"os"
)

// Exercises 1.5 and 1.6
var palette = []color.Color{color.RGBA{85, 165, 34, 50}, color.RGBA{230, 255, 230, 0}, color.RGBA{0, 102, 0, 0}, color.RGBA{255, 255, 204, 0}}

func main() {
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
    const (
        cycles  = 10     // number of complete x oscillator revolutions
        res     = 0.001 // angular resolution
		size    = 200
		nframes = 64
		delay   = 8
		// image canvas covers [-size..+size]
		// number of animation frames
		// delay between frames in 10ms units
    )
    freq := rand.Float64() * 3.0 // relative frequency of y oscillator
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0 // phase difference
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < cycles*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), random(1,4))
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
	}
    gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func random(min, max int) uint8 {
    rand.Seed(time.Now().Unix())
    return uint8(rand.Intn(max - min) + min)
}