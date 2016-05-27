     

package gif    

import (
"image"
"image/color"
"image/gif"
"io"
"math"
"math/rand"
)


var palette = []color.Color{color.White, color.RGBA{30, 30, 30, 11},color.RGBA{ 0xA3, 90, 90, 11} }

const (
        whiteIndex = 0 // first color in palette
        colorA = 1 // next color in palette
        colorB = 2
        )


func Lissajous(out io.Writer, cycle float64) {
   const (
    // number of complete x oscillator revolutions
    res     = 0.001 // angular resolution
    size    = 100
    nframes = 64
    delay   = 8
    )

    cycles  := cycle 
    // image canvas covers [-size..+size]
    // number of animation frames
    // delay between frames in 10ms units

    freq := rand.Float64() * 3.0 // relative frequency of y oscillator
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0 // phase difference
    
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < cycles*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
                uint8(i%2+1))
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

