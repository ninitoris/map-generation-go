package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fogleman/gg"
)

const (
	roomsNumX = 30
	roomsNumY = 20
)

const (
	zalip   = false
	timeout = 200 * time.Millisecond
)

const (
	roomWidth  = 64
	roomHeight = 64
)

const (
	WALL_WIDTH = 10
	W          = roomsNumX*roomWidth + WALL_WIDTH
	H          = roomsNumY*roomHeight + WALL_WIDTH
)

type direction int

const (
	Top direction = 1 << iota
	Right
	Down
	Left
)

var defaultDirections = [4]direction{Top, Right, Down, Left}

var dc *gg.Context

var seed int64
var random *rand.Rand

func getSeedFromString(s string) int64 {
	if s == "" {
		return rand.Int63()
	} else {
		h := fnv.New64a()       // Создаем хешер FNV-1a 64-битный
		h.Write([]byte(s))      // Хешируем строку
		return int64(h.Sum64()) // Возвращаем результат как int64
	}
}

func SetBackgroundColor(dc *gg.Context) {
	dc.SetRGBA(1, 1, 1, 1)
}

func SetMainColor(dc *gg.Context) {
	dc.SetRGBA(0, 0, 0, 1)
}

func main() {
	dc = gg.NewContext(W, H)
	SetBackgroundColor(dc)
	dc.Clear()
	SetMainColor(dc)
	dc.SetLineWidth(WALL_WIDTH)

	fmt.Print("Maze seed (empty for random seed): ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error!", err)
		return
	}
	input = strings.TrimSuffix(input, "\n")

	seed = getSeedFromString(input)
	random = rand.New(rand.NewSource(int64(seed)))

	var mazeGrid MazeGrid = make(MazeGrid, roomsNumX)
	for i := range mazeGrid {
		mazeGrid[i] = make([]*Room, roomsNumY)
	}

	for x := 0; x < roomsNumX; x++ {
		for y := 0; y < roomsNumY; y++ {
			room := &Room{
				x:       x,
				y:       y,
				doors:   0,
				visited: false,
				pathLen: 0,
			}
			mazeGrid[x][y] = room
		}
	}

	if zalip {
		for i := range mazeGrid {
			for j := range mazeGrid[i] {
				room := mazeGrid[i][j]
				room.DrawRoom()
			}
		}
	}

	generateMaze(nil, mazeGrid[0][0], mazeGrid)

	if !zalip {
		for i := range mazeGrid {
			for j := range mazeGrid[i] {
				room := mazeGrid[i][j]
				// fmt.Printf("Room [%d,%d]: pathlen %d\n", room.x, room.y, room.pathLen)
				room.DrawRoom()
			}
		}
	}
	createExit(mazeGrid)

	dc.SavePNG("out.png")
}
