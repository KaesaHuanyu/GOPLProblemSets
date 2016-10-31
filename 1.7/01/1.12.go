package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black} //调色板 切片

var (
	cycles  = 5
	res     = 0.001
	size    = 100 // image canvas covers [-size, +size]
	nframes = 64  // 帧数
	delay   = 8   // 帧间时
)

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			switch k {
			case "cycles":
				cycles, _ = strconv.Atoi(v[0])
			case "res":
				res, _ = strconv.ParseFloat(v[0], 64)
			case "size":
				size, _ = strconv.Atoi(v[0])
			case "nframes":
				nframes, _ = strconv.Atoi(v[0])
			case "delay":
				delay, _ = strconv.Atoi(v[0])
			}
		}
		lissajous(w) //都实现了io.Writer接口
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5),
				size+int(y*float64(size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
