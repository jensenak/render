/*
The MIT License (MIT)

Copyright (c) 2015 Angus Gibson

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// This file has been heavily modified by Adam Jensen
// Original source can be found here: https://github.com/angus-g/go-obj/blob/master/obj/obj.go

// Package obj allows simple parsing of Wavefront OBJ files
// into slices of vertices, normals and elements.
package obj

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Parse a vertex (normal) string, a list of whitespace-separated
// floating point numbers.
func parseVertex(t []string) []float32 {
	x, _ := strconv.ParseFloat(t[0], 32)
	y, _ := strconv.ParseFloat(t[1], 32)
	z, _ := strconv.ParseFloat(t[2], 32)

	return []float32{float32(x), float32(y), float32(z)}
}

// Parse an element string, a list of whitespace-separated elements.
// Elements are of the form "<vi>/<ti>/<ni>" where indices are the
// vertex, texture coordinate and normal, respectively.
func parseElement(t []string) [3]int {
	var e [3]int
	var err error

	for i := 0; i < 3; i++ {
		e[i], err = strconv.Atoi(strings.Split(t[i], "/")[0])
		if err != nil {
			panic(err)
		}
		e[i]--
	}

	return e
}

// Parse a wavefront object file, returning a list of vertices and faces
func Parse(filename string) ([][]float32, [][3]int) {
	fp, _ := os.Open(filename)
	scanner := bufio.NewScanner(fp)

	vertices := [][]float32{}
	normals := [][]float32{}
	elements := [][3]int{}

	for scanner.Scan() {
		toks := strings.Fields(strings.TrimSpace(scanner.Text()))

		// Skip empty lines
		if len(toks) == 0 {
			continue
		}
		switch toks[0] {
		case "v":
			vertices = append(vertices, parseVertex(toks[1:]))
		case "vn":
			normals = append(normals, parseVertex(toks[1:]))
		case "f":
			elements = append(elements, parseElement(toks[1:]))
		}
	}

	return vertices, elements
}
