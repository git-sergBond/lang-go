// Package lissajous
// Generate GIF lissajous and save it to file
package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.Black, // background 	- Black
	color.RGBA{R: 0x7E, G: 0xDA, B: 0xEE, A: 0xFF}, // color 1 		- #7EDAEE
	color.RGBA{R: 0x00, G: 0x92, B: 0xFA, A: 0xFF}, // color 2 		- #0092FA
	color.RGBA{R: 0x14, G: 0x24, B: 0xD3, A: 0xFF}, // color 3 		- #1424D3
	color.RGBA{R: 0x00, G: 0xA5, B: 0xDE, A: 0xFF}, // color 4 		- #00A5DE
	color.RGBA{R: 0x9D, G: 0xE1, B: 0x7D, A: 0xFF}, // color 5 		- #9DE17D
	color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF}, // color 6 		- Green
}

func Lissajous() {
	file, err := os.Create("lissajous.gif")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		errClose := file.Close()
		if errClose != nil {
			log.Fatal(errClose)
		}
	}(file)
	LissajousGif(file, nil)
}

// LissajousGif
// settings can be nil or redefine some variables
func LissajousGif(out io.Writer, settings map[string]any) {
	// constants for GIF
	const (
		_cycles  = 5     // Количество полных колебаний
		_res     = 0.001 // Угловое разрешение
		_size    = 100   // Канва изображения охватывает [_size..+_size]
		_nFrames = 64    // Количество кадров анимации
		_delay   = 8     // Задержка между кадрами
	)

	// override settings
	var cycles = float64(_cycles)
	if len(settings) > 0 {
		if val, ok := settings["cycles"].(int); ok {
			cycles = float64(val)
		}
	}
	log.Printf("DEBUG: cycles=%v", cycles)

	// variables for gradient
	var lenPalette = uint8(len(palette)) - 1
	var lineColorIndex uint8 = 1
	gradientSwitchAfterFrames := _nFrames / lenPalette

	// variables for figure
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	freq := r.Float64() * 3.0 // Относительная частота колебаний y
	anim := gif.GIF{LoopCount: _nFrames}
	phase := 0.0 // разность фаз

	// algorithm draw frames for GIF
	for frame := 0; frame < _nFrames; frame++ {

		// create new empty frame
		rect := image.Rect(0, 0, 2*_size+1, 2*_size+1)
		img := image.NewPaletted(rect, palette)

		// algorithm draw for gradient
		if frame%int(gradientSwitchAfterFrames) == 0 {
			if lineColorIndex < lenPalette {
				lineColorIndex++
			} else {
				lineColorIndex = 1
			}
		}

		// algorithm draw for figure
		for t := 0.0; t < cycles*2*math.Pi; t += _res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(_size+int(x*_size+0.5), _size+int(y*_size+0.5), lineColorIndex)
		}

		// create animation
		phase += 0.1
		anim.Delay = append(anim.Delay, _delay)
		anim.Image = append(anim.Image, img)
	}

	// save animation
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		log.Fatal(err)
	}
}
