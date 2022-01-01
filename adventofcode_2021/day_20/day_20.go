package day_20

import (
	"adventofcode_2021/errutil"
	"adventofcode_2021/readutil"
	"fmt"
	gimage "image"
	"image/color"
	"image/png"
	"os"

	"github.com/pkg/errors"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	res, err := numLitPixels(inputAlgo, inputImg, 2)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := numLitPixels(inputAlgo, inputImg, 50)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

func parseAlgorithm(in string) (*algorithm, error) {
	a := &algorithm{}
	in = readutil.ReadString(in)
	for _, r := range in {
		switch r {
		case '#':
			a.values = append(a.values, true)
		case '.':
			a.values = append(a.values, false)
		default:
			return nil, errors.Errorf("invalid pixel %q", string(r))
		}
	}
	if len(a.values) == 0 {
		return nil, errors.Errorf("no entries")
	}
	return a, nil
}

func parseImage(in string) (*image, error) {
	img := newImage(false)
	lines := readutil.ReadLines(in)
	for y, line := range lines {
		for x, r := range line {
			switch r {
			case '#':
				img.set(p(x, y), true)
			case '.':
				img.set(p(x, y), false)
			default:
				return nil, errors.Errorf("invalid pixel %q", string(r))
			}
		}
	}
	if img.isEmpty() {
		return nil, errors.Errorf("image is empty")
	}

	return img, nil
}

func bitsToNumber(bs []bool) int {
	if len(bs) == 0 {
		return 0
	}
	var n int
	// this is big endian, p.e. 100 => 4
	for i := 0; i < len(bs); i++ {
		b := bs[len(bs)-1-i]
		if !b {
			continue
		}
		n |= 1 << i
	}
	return n
}

type algorithm struct {
	values []bool
}

func (a *algorithm) get(idx int) bool {
	if idx < 0 || idx >= len(a.values) {
		return false
	}
	return a.values[idx]
}

type point struct {
	x, y int
}

func p(x, y int) point {
	return point{x: x, y: y}
}

func (pt point) pointsAround() []point {
	var ps []point
	for y := pt.y - 1; y <= pt.y+1; y++ {
		for x := pt.x - 1; x <= pt.x+1; x++ {
			ps = append(ps, p(x, y))
		}
	}
	return ps
}

type image struct {
	pixels     map[point]bool
	minX, maxX int
	minY, maxY int
	outsideVal bool
}

func (img *image) isEmpty() bool {
	return len(img.pixels) == 0
}

func newImage(outsideVal bool) *image {
	return &image{
		pixels:     map[point]bool{},
		outsideVal: outsideVal,
	}
}

func (img *image) set(pt point, val bool) {
	if len(img.pixels) == 0 {
		img.minX, img.maxX, img.minY, img.maxY = pt.x, pt.x, pt.y, pt.y
	} else {
		if pt.x < img.minX {
			img.minX = pt.x
		}
		if pt.x > img.maxX {
			img.maxX = pt.x
		}
		if pt.y < img.minY {
			img.minY = pt.y
		}
		if pt.y > img.maxY {
			img.maxY = pt.y
		}
	}

	img.pixels[pt] = val
}

func (img *image) get(pt point) bool {
	if pt.x < img.minX || pt.x > img.maxX ||
		pt.y < img.minY || pt.y > img.maxY {
		return img.outsideVal
	}

	if v, ok := img.pixels[pt]; ok {
		return v
	}
	return false
}

func (img *image) windowNum(pt point) int {
	psa := pt.pointsAround()
	var numB []bool
	for _, p := range psa {
		numB = append(numB, img.get(p))
	}
	return bitsToNumber(numB)
}

func (img *image) numLitPixels() int {
	var numLit int
	for _, v := range img.pixels {
		if v {
			numLit++
		}
	}
	return numLit
}

func (img *image) applied(algo *algorithm) *image {
	aimg := newImage(img.outsideVal)
	pad := 1
	for x := img.minX - pad; x <= img.maxX+pad; x++ {
		for y := img.minY - pad; y <= img.maxY+pad; y++ {
			wn := img.windowNum(p(x, y))
			aval := algo.get(wn)
			aimg.set(p(x, y), aval)
		}
	}

	return aimg
}

func numLitPixels(inAlgo string, inImg string, steps int) (int, error) {
	algo, err := parseAlgorithm(inAlgo)
	if err != nil {
		return 0, errors.Wrap(err, "parse-algo")
	}
	img, err := parseImage(inImg)
	if err != nil {
		return 0, errors.Wrap(err, "parse-image")
	}

	flipOutside := false
	if algo.values[0] {
		flipOutside = true
	}

	//log("*** initial ***\n%s", img.format())
	for i := 0; i < steps; i++ {
		img = img.applied(algo)
		if flipOutside {
			img.outsideVal = !img.outsideVal
		}
		//log("\n*** after %d ***\n%s", i+1, img.format())
	}

	num := img.numLitPixels()

	pngImg := gimage.NewGray(gimage.Rect(img.minX, img.minY, img.maxX, img.maxY))
	for p, v := range img.pixels {
		if v {
			pngImg.Set(p.x, p.y, color.White)
		} else {
			pngImg.Set(p.x, p.y, color.Black)
		}
	}
	f, _ := os.Create("trench.png")
	defer f.Close()
	png.Encode(f, pngImg)

	return num, nil
}
