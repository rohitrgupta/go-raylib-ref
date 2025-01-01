package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	Width  = 80
	Height = 35
)

type Game struct {
	World [][]int
	Next  [][]int
}

func NewGame() *Game {
	world := make([][]int, Height)
	for i := range world {
		world[i] = make([]int, Width)
	}
	next := make([][]int, Height)
	for i := range next {
		next[i] = make([]int, Width)
	}
	return &Game{
		World: world,
		Next:  next,
	}
}

func (g *Game) ClearNext() {
	for i := range g.Next {
		for j := range g.Next[i] {
			g.Next[i][j] = 0
		}
	}
}

func (g *Game) CopyNextToWorld() {
	for i := range g.World {
		for j := range g.World[i] {
			g.World[i][j] = g.Next[i][j]
		}
	}
}

func (g *Game) Print() {
	for i := range g.World {
		for j := range g.World[i] {
			if g.World[i][j] == 1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (g *Game) CountNeighbors(i, j int) int {
	count := 0
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if x == i && y == j {
				continue
			}
			if x < 0 || y < 0 || x >= Height || y >= Width {
				continue
			}
			if g.World[x][y] == 1 {
				count++
			}
		}
	}
	return count
}

func (g *Game) Update() {
	for i := range g.World {
		for j := range g.World[i] {
			neighbors := g.CountNeighbors(i, j)
			if g.World[i][j] == 1 {
				if neighbors < 2 || neighbors > 3 {
					g.Next[i][j] = 0
				} else {
					g.Next[i][j] = 1
				}
			} else {
				if neighbors == 3 {
					g.Next[i][j] = 1
				} else {
					g.Next[i][j] = 0
				}
			}
		}
	}
	g.CopyNextToWorld()
}

func (g *Game) ReadFile() {
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Iterate through the records and print them
	for _, record := range records {
		fmt.Println(record)
		y, _ := strconv.Atoi(record[0])
		x, _ := strconv.Atoi(record[1])
		g.World[y][x] = 1
	}
}

func main() {
	game := NewGame()
	game.ReadFile()
	// game.World[20][40] = 1
	// game.World[20][41] = 1
	// game.World[20][42] = 1
	// game.World[20][43] = 1
	// game.World[20][44] = 1
	// game.World[20][45] = 1

	for i := 0; i < 1000; i++ {
		game.Update()
		game.Print()
		time.Sleep(500 * time.Millisecond)
	}
}
