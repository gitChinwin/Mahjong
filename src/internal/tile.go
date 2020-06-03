package internal

import (
	"log"
)

/**
* @Author: Jam Wong
* @Date: 2020/6/3 3:01 下午
 */

type tileType int

var (
	suitWords   = []rune("一二三四五六七八九")
	windWords   = []rune("东南西北")
	dragonWords = []rune("中发白")
)

const (
	TileDragon    tileType = 1
	TileWind      tileType = 2
	TileDot       tileType = 3
	TileBamboo    tileType = 4
	TileCharacter tileType = 5
)

type Tile struct {
	Type tileType
	Rank int
}

func (t *Tile) Print() string {
	if t == nil {
		log.Fatal("tile is not init")
	}

	var unit []rune

	if t.Rank < 1 || t.Rank > 9 {
		log.Fatal("num out of range")
	}

	switch t.Type {
	case TileDot:
		unit = append(unit, suitWords[t.Rank-1], '筒')
	case TileBamboo:
		unit = append(unit, suitWords[t.Rank-1], '条')
	case TileCharacter:
		unit = append(unit, suitWords[t.Rank-1], '万')
	case TileDragon:
		unit = append(unit, dragonWords[t.Rank-1])
	case TileWind:
		unit = append(unit, windWords[t.Rank-1])
	default:
		log.Fatal("type not found")
	}
	return string(unit)
}

func AllTiles() []*Tile {
	a := make([]*Tile, 0)

	for i := 1; i <= 4; i++ {
		a = append(a,
			&Tile{TileDragon, 1},
			&Tile{TileDragon, 2},
			&Tile{TileDragon, 3},
			&Tile{TileWind, 1},
			&Tile{TileWind, 2},
			&Tile{TileWind, 3},
			&Tile{TileWind, 4},
			&Tile{TileDot, 1},
			&Tile{TileDot, 2},
			&Tile{TileDot, 3},
			&Tile{TileDot, 4},
			&Tile{TileDot, 5},
			&Tile{TileDot, 6},
			&Tile{TileDot, 7},
			&Tile{TileDot, 8},
			&Tile{TileDot, 9},
			&Tile{TileBamboo, 1},
			&Tile{TileBamboo, 2},
			&Tile{TileBamboo, 3},
			&Tile{TileBamboo, 4},
			&Tile{TileBamboo, 5},
			&Tile{TileBamboo, 6},
			&Tile{TileBamboo, 7},
			&Tile{TileBamboo, 8},
			&Tile{TileBamboo, 9},
			&Tile{TileCharacter, 1},
			&Tile{TileCharacter, 2},
			&Tile{TileCharacter, 3},
			&Tile{TileCharacter, 4},
			&Tile{TileCharacter, 5},
			&Tile{TileCharacter, 6},
			&Tile{TileCharacter, 7},
			&Tile{TileCharacter, 8},
			&Tile{TileCharacter, 9},
		)
	}
	return a
}
