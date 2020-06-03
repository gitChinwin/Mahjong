package internal

import (
	"fmt"
	"log"
)

/**
* @Author: Jam Wong
* @Date: 2020/6/3
 */

func (pl *Player) Win2() bool {
	pl.SortTiles()

	// 牌数量 3n+2
	if (pl.HoldTiles.Len()-2)%3 != 0 {
		return false
	}

	tmp := make([]*Tile, 0)
	copy(tmp, pl.HoldTiles)

	for i := 0; i < len(pl.HoldTiles)-1; i++ {
		if isPair(pl.HoldTiles[i], pl.HoldTiles[i+1]) {
			tmp := make([]*Tile, 0)

			tmp = pl.HoldTiles
			fmt.Println(pl.HoldTiles)
			remove(tmp, i+1)
			remove(tmp, i)

			if hu(tmp) {
				return true
			}
		}
	}

	return false
}

func remove(t []*Tile, n int) {
	if len(t) < 2 {
		log.Fatal("at least 2 tile.")
	}
	t[n] = t[len(t)-1]
	t[len(t)-1] = nil
	t = t[:len(t)-1]
}

func hu(t []*Tile) bool {
	if len(t) == 0 {
		return true
	}
	if isSequence(t[0], t[1], t[2]) {
		return hu(t[3:])
	}

	if isTriplet(t[0], t[1], t[2]) {
		return hu(t[3:])
	}
	return false
}
