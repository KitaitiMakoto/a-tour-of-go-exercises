package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	rows := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		rows[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			rows[i][j] = uint8((i + j) / 2)
		}
	}
	return rows
}

func main() {
	pic.Show(Pic)
}
