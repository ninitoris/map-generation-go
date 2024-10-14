package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

type MazeGrid = [][]*Room

func shuffleArrayWithSeed(arr []*Room, r *rand.Rand) {
	sort.Slice(arr, func(i, j int) bool {
		return r.Intn(2) == 0 // Псевдослучайное сравнение
	})
}

func getAvailableRooms(curRoom *Room, mazeGrid MazeGrid) []*Room {
	var availableRooms []*Room
	if curRoom.x > 0 {
		checkRoom := mazeGrid[curRoom.x-1][curRoom.y]
		if !checkRoom.visited {
			availableRooms = append(availableRooms, checkRoom)
		}
	}

	if curRoom.x < roomsNumX-1 {
		checkRoom := mazeGrid[curRoom.x+1][curRoom.y]
		if !checkRoom.visited {
			availableRooms = append(availableRooms, checkRoom)
		}
	}

	if curRoom.y > 0 {
		checkRoom := mazeGrid[curRoom.x][curRoom.y-1]
		if !checkRoom.visited {
			availableRooms = append(availableRooms, checkRoom)
		}
	}
	if curRoom.y < roomsNumY-1 {
		checkRoom := mazeGrid[curRoom.x][curRoom.y+1]
		if !checkRoom.visited {
			availableRooms = append(availableRooms, checkRoom)
		}
	}

	return availableRooms
}

func generateMaze(prevRoom *Room, curRoom *Room, mazeGrid MazeGrid) {
	// fmt.Printf("Current room [%d,%d]\n", curRoom.x, curRoom.y)
	if curRoom.visited {
		return
	}
	curRoom.visited = true

	if prevRoom == nil {
		curRoom.doors = Top
	} else {
		curRoom.pathLen = prevRoom.pathLen + 1

		if prevRoom.x < curRoom.x {
			prevRoom.doors = prevRoom.doors | Right
			curRoom.doors = curRoom.doors | Left
		} else if prevRoom.x > curRoom.x {
			prevRoom.doors = prevRoom.doors | Left
			curRoom.doors = curRoom.doors | Right
		} else if prevRoom.y < curRoom.y {
			prevRoom.doors = prevRoom.doors | Down
			curRoom.doors = curRoom.doors | Top
		} else if prevRoom.y > curRoom.y {
			prevRoom.doors = prevRoom.doors | Top
			curRoom.doors = curRoom.doors | Down
		}
	}

	if zalip {
		curRoom.DrawRoom()

		if prevRoom != nil {
			prevRoom.DrawRoom()
		}

		dc.SavePNG("out.png")
		time.Sleep(timeout)
	}

	availableRooms := getAvailableRooms(curRoom, mazeGrid)
	// sort.Slice(availableRooms, func(i, j int) bool {
	// 	return 0.5-rand.Float32() > 0
	// })

	shuffleArrayWithSeed(availableRooms, random)

	// fmt.Printf("Room [%d,%d] has available rooms:\n", curRoom.x, curRoom.y)
	for _, r := range availableRooms {
		// fmt.Printf("[%d,%d]\n", r.x, r.y)

		generateMaze(curRoom, r, mazeGrid)
	}
}

func findExitRoom(mazeGrid MazeGrid) *Room {
	max := 0
	var exitRoom *Room
	for i := range mazeGrid {
		for j := range mazeGrid[i] {
			room := mazeGrid[i][j]
			if room.pathLen > max {
				max = room.pathLen
				exitRoom = room
			}
		}
	}

	return exitRoom
}

func createExit(mazeGrid MazeGrid) {
	var exitRoom *Room = findExitRoom(mazeGrid)
	if exitRoom == nil {
		fmt.Println("Could not find an exit room")
		return
	}
	// fmt.Println("Exit room", exitRoom.x, exitRoom.y)

	im, err := gg.LoadPNG("goool.png")
	if err != nil {
		panic(err)
	}
	im = resize.Resize(roomWidth, roomHeight, im, resize.Lanczos3)

	// dc.DrawImage(im, 100, 100)
	dc.DrawImage(im, exitRoom.x*roomWidth, exitRoom.y*roomHeight)
	// dc.SavePNG("out.png")
}
