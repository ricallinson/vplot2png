package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
)

type plot struct {
	cmd string
	x   int
	y   int
}

func readPlotFile(f *os.File) (plots []*plot) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		cmd := line[0]
		if len(line) == 3 {
			x, _ := strconv.Atoi(line[1])
			y, _ := strconv.Atoi(line[2])
			plots = append(plots, &plot{cmd, x, y})
		}
	}
	return plots
}

func getCanvasSize(plots []*plot) (x, y int) {
	for _, p := range plots {
		if x < p.x {
			x = p.x
		}
		if y < p.y {
			y = p.y
		}
	}
	return x, y
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func drawline(x0, y0, x1, y1 int, img *image.RGBA, penSize int) {
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	sx, sy := 1, 1
	if x0 >= x1 {
		sx = -1
	}
	if y0 >= y1 {
		sy = -1
	}
	err := dx - dy
	for {
		for s := 0; s <= penSize; s++ {
			img.Set(x0+s, y0+s, color.Black)
		}
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func draw(plots []*plot, penSize int) *image.RGBA {
	w, h := getCanvasSize(plots)
	img := image.NewRGBA(image.Rect(0, 0, w+1, h+1))
	cx, cy := 0, 0
	// fmt.Printf("x: %v, y: %v\n", cx, cy)
	for _, p := range plots {
		if p.cmd == "l" {
			drawline(cx, cy, cx+p.x, cy+p.y, img, penSize)
			cx, cy = cx+p.x, cy+p.y
		} else if p.cmd == "L" {
			drawline(cx, cy, p.x, p.y, img, penSize)
			cx, cy = p.x, p.y
		} else if p.cmd == "m" {
			cx, cy = cx+p.x, cy+p.y
		} else if p.cmd == "M" {
			cx, cy = p.x, p.y
		}
		// fmt.Printf("x: %v, y: %v\n", cx, cy)
	}
	return img
}

func main() {

	var penSize = flag.Int("p", 5, "the size of the pen tip used for the drawing")

	flag.Parse()
	source := flag.Arg(0)
	dest := flag.Arg(1)

	if source == "" {
		fmt.Println("A source vplot file must be provide as the first argument.")
		return
	}

	if dest == "" {
		dest = strings.Replace(source, ".vplot", "-vplot.png", 1)
	}

	vplotData, err := os.Open(source)

	if err != nil {
		fmt.Println("Could not open the source vplot file.")
		return
	}

	destData := draw(readPlotFile(vplotData), *penSize)

	file, _ := os.Create(dest)
	defer file.Close()
	png.Encode(file, destData)
}
