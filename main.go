package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"./obj"
)

type blob struct {
	img *image.RGBA
	col color.Color
}

// swap receives two pointers and swaps the contents between them
func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

// abs returns the absolute value of an integer
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (b *blob) line(x0, y0, x1, y1 int) {
	// If abs(slope) > 1, switch axes
	steep := false
	if abs(x1-x0) < abs(y1-y0) {
		swap(&x0, &y0)
		swap(&x1, &y1)
		steep = true
	}
	// If line goes right to left, switch start/end points
	if x0 > x1 {
		swap(&x0, &x1)
		swap(&y0, &y1)
	}
	// Find the delta
	dx := x1 - x0
	dy := y1 - y0
	// And the error rate, to know when to increment y
	derror := abs(dy) * 2
	error := 0
	y := y0
	for x := x0; x <= x1; x++ {
		// Remember if the axes are switched, revert
		if steep {
			b.img.Set(y, x, b.col)
		} else {
			b.img.Set(x, y, b.col)
		}
		// If error has accumulated enough, increment/decrement y
		error += derror
		if error > dx {
			if y0 > y1 {
				y--
			} else {
				y++
			}
			error -= dx * 2
		}
	}
	return
}

// Save the contents to a file as PNG
func toFile(img *image.RGBA, name string) {
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

func main() {
	width := 600
	height := 600
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	b := &blob{img: img, col: color.RGBA{0, 0, 0, 255}}
	verts, faces := obj.Parse("head.obj")

	for i := 0; i < len(faces); i++ {
		face := faces[i]
		for j := 0; j < 3; j++ {
			v0 := verts[face[j]]
			v1 := verts[face[(j+1)%3]]
			x0 := int((v0[0] + 1) * float32(width) / 2)
			y0 := height - int((v0[1]+1)*float32(height)/2)
			x1 := int((v1[0] + 1) * float32(width) / 2)
			y1 := height - int((v1[1]+1)*float32(height)/2)
			b.line(x0, y0, x1, y1)
		}
	}
	toFile(img, "face.png")
}
