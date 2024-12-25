package main

import (
	"bufio"
	"fmt"
	"image/color"
	"os"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	width  = 800
	height = 450
)

const (
	fontHeight = 30
)

type Shape interface {
	Draw()
	Update()
}

type Circle struct {
	X, Y, R float32
	vx, vy  float32
	n       string
	color   rl.Color
}

func (c *Circle) Draw() {
	rl.DrawCircle(int32(c.X), int32(c.Y), c.R, c.color)
	length := rl.MeasureText(c.n, fontHeight)
	rl.DrawText(
		c.n,
		int32(c.X-float32(length/2)),
		int32(c.Y-fontHeight/2),
		30,
		rl.Black,
	)
}

func (c *Circle) Update() {
	c.X += c.vx
	c.Y += c.vy
	if c.X < c.R || c.X > float32(width)-c.R {
		c.vx = -c.vx
	}
	if c.Y < c.R || c.Y > float32(height)-c.R {
		c.vy = -c.vy
	}
}

type Rectangle struct {
	X, Y, W, H float32
	vx, vy     float32
	n          string
	color      rl.Color
}

func (r *Rectangle) Draw() {
	rl.DrawRectangle(int32(r.X-r.W/2), int32(r.Y-r.H/2), int32(r.W), int32(r.H), r.color)
	length := rl.MeasureText(r.n, fontHeight)
	rl.DrawText(
		r.n,
		int32(r.X-float32(length/2)),
		int32(r.Y-fontHeight/2),
		30,
		rl.Black,
	)
}

func (r *Rectangle) Update() {
	r.X += r.vx
	r.Y += r.vy
	if r.X < r.W/2 || r.X > float32(width)-r.W/2 {
		r.vx = -r.vx
	}
	if r.Y < r.H/2 || r.Y > float32(height)-r.H/2 {
		r.vy = -r.vy
	}
}

func main() {
	// width, height := 800, 450
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	shapes := []Shape{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	// col := rl.NewColor(255, 255, 255, 255)
	bkColor := rl.Gray
	for scanner.Scan() {
		if scanner.Text() == "W" {
			bkColor = readColor(scanner)
			scanner.Scan()
			width, _ = strconv.Atoi(scanner.Text())
			scanner.Scan()
			height, _ = strconv.Atoi(scanner.Text())
			fmt.Println("Window", width, height)
		}
		if scanner.Text() == "C" {
			c1 := readCircle(scanner)
			shapes = append(shapes, &c1)
		}
		if scanner.Text() == "R" {
			r1 := readRectangle(scanner)
			shapes = append(shapes, &r1)
		}
	}
	fmt.Println("shapes", shapes)

	rl.InitWindow(int32(width), int32(height), "Shapes")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	// c := Circle{X: 128, Y: 256, R: 20, color: rl.Red}
	// r := Rectangle{X: 360, Y: 256, W: 50, H: 20, color: rl.Yellow}
	// shapes = append(shapes, c, r)

	for !rl.WindowShouldClose() {
		for _, shape := range shapes {
			shape.Update()
		}
		rl.BeginDrawing()
		rl.ClearBackground(bkColor)
		for _, shape := range shapes {
			shape.Draw()
		}
		rl.EndDrawing()
	}
}

func readCircle(scanner *bufio.Scanner) Circle {
	col := readColor(scanner)
	scanner.Scan()
	n := scanner.Text()
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	y, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	vx, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	vy, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	r, _ := strconv.Atoi(scanner.Text())
	c1 := Circle{X: float32(x), Y: float32(y), R: float32(r), color: col, n: n, vx: float32(vx), vy: float32(vy)}
	return c1
}

func readRectangle(scanner *bufio.Scanner) Rectangle {
	col := readColor(scanner)
	scanner.Scan()
	n := scanner.Text()
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	y, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	vx, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	vy, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	w, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	h, _ := strconv.Atoi(scanner.Text())
	r1 := Rectangle{X: float32(x), Y: float32(y), W: float32(w), H: float32(h), color: col, n: n, vx: float32(vx), vy: float32(vy)}
	return r1
}

func readColor(scanner *bufio.Scanner) color.RGBA {
	scanner.Scan()
	cr, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	cg, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	cb, _ := strconv.Atoi(scanner.Text())
	col := rl.NewColor(uint8(cr), uint8(cg), uint8(cb), 255)
	return col
}
