package main

import "github.com/fogleman/gg"

type Room struct {
	x, y    int
	visited bool
	doors   direction // 1<< iota: up 1 right 2 down 4 left 8
	pathLen int
}

// func (room Room) hasOnlyOneDoor() bool {
// 	return room.doors != 0 && (room.doors&(room.doors-1)) == 0
// }

func (room Room) RemoveWall(dc *gg.Context, direction direction) {
	SetBackgroundColor(dc)
	roomX := room.x * roomWidth
	roomY := room.y * roomHeight

	x := float64(roomX) + WALL_WIDTH/2
	y := float64(roomY) + WALL_WIDTH/2
	w := float64(roomWidth) - WALL_WIDTH
	h := float64(roomHeight) - WALL_WIDTH

	if direction == Top {
		y -= WALL_WIDTH / 2
		h += WALL_WIDTH / 2
	} else if direction == Down {
		h += WALL_WIDTH / 2
	} else if direction == Right {
		w += WALL_WIDTH / 2
	} else if direction == Left {
		x -= WALL_WIDTH / 2
		w += WALL_WIDTH / 2
	}

	dc.DrawRectangle(
		x+float64(roomOffsetX(room)),
		y+float64(roomOffsetY(room)),
		w, h,
	)
	dc.Fill()
	SetMainColor(dc)
}

func (room Room) HasDoor(dir direction) bool {
	if room.doors&dir != 0 {
		return true
	} else {
		return false
	}
}

func (room Room) GetDoorsAsArray() []direction {
	var doors []direction

	for _, d := range defaultDirections {
		if room.HasDoor(d) {
			doors = append(doors, d)
		}
	}

	return doors
}

func (room Room) DrawRoom() {
	roomX := room.x * roomWidth
	roomY := room.y * roomHeight

	dc.DrawRectangle(
		float64(roomX+roomOffsetX(room)),
		float64(roomY+roomOffsetY(room)),
		float64(roomWidth),
		float64(roomHeight),
	)

	dc.Fill()

	SetMainColor(dc)

	if room.doors != 0 {
		doorsDirs := room.GetDoorsAsArray()

		for _, d := range doorsDirs {
			room.RemoveWall(dc, d)
		}
	}
}

func roomOffsetX(room Room) int {
	return room.x * 0
}

func roomOffsetY(room Room) int {
	return room.y * 0
}
