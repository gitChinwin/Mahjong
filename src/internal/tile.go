package internal

import (
	"log"
	"sort"
)

/********************
* @Author: Jam Wong *
* @Date: 2020/6/3   *
 ********************/

type tileType int

type ts []*Tile

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

func (tile *Tile) SameAs(t *Tile) bool {
	return tile.Type == t.Type && tile.Rank == t.Rank
}

func (tile *Tile) Print() string {
	if tile == nil {
		log.Fatal("tile is not init")
	}

	var unit []rune

	if tile.Rank < 1 || tile.Rank > 9 {
		log.Fatal("num out of range")
	}

	switch tile.Type {
	case TileDot:
		unit = append(unit, suitWords[tile.Rank-1], '筒')
	case TileBamboo:
		unit = append(unit, suitWords[tile.Rank-1], '条')
	case TileCharacter:
		unit = append(unit, suitWords[tile.Rank-1], '万')
	case TileDragon:
		unit = append(unit, dragonWords[tile.Rank-1])
	case TileWind:
		unit = append(unit, windWords[tile.Rank-1])
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

func (t ts) Len() int {
	return len(t)
}

// 如果 i 索引的数据小于 j 索引的数据，返回 true，且不会调用下面的 Swap()，即数据升序排序。
func (t ts) Less(i, j int) bool {
	iObj := t[i]
	jObj := t[j]
	if iObj.Type < jObj.Type {
		return true
	}
	if iObj.Type == jObj.Type {
		return iObj.Rank <= jObj.Rank
	}
	return false
}

func (t ts) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t ts) SortTiles() ts {
	ttt := make(ts, 0)
	ttt = append(ttt, t...)
	sort.Sort(ttt)
	return ttt
}

func (t ts) Show(from, to int) string {
	var s string
	s += "|"
	for idx, i := range t[from:to] {
		s += i.Print()
		s += "|"
		if idx%4 == 3 {
			s += "  |"
		}
	}
	return s
}

func (t ts) HaveConcealedKong() bool {
	ttt := t.SortTiles()
	if len(ttt) < 4 {
		return false
	}

	for i := 0; i < len(ttt)-3; i++ {
		if ttt[i].SameAs(ttt[i+1]) && ttt[i].SameAs(ttt[i+2]) && ttt[i].SameAs(ttt[i+3]) {
			return true
		}
	}
	return false
}

func (t ts) CanKong(tile *Tile) bool {
	tmp := make(ts, 0)
	tmp = append(tmp, t...)
	tmp = append(tmp, tile)

	ttt := tmp.SortTiles()
	if len(ttt) < 4 {
		return false
	}

	for i := 0; i < len(ttt)-3; i++ {
		if ttt[i].SameAs(tile) && ttt[i].SameAs(ttt[i+1]) && ttt[i].SameAs(ttt[i+2]) && ttt[i].SameAs(ttt[i+3]) {
			return true
		}
	}
	return false
}

func (t ts) CanPong(tile *Tile) bool {
	tmp := make(ts, 0)
	tmp = append(tmp, t...)
	tmp = append(tmp, tile)

	ttt := tmp.SortTiles()
	if len(ttt) < 3 {
		return false
	}

	for i := 0; i < len(ttt)-2; i++ {
		if ttt[i].SameAs(tile) && ttt[i].SameAs(ttt[i+1]) && ttt[i].SameAs(ttt[i+2]) {
			return true
		}
	}
	return false
}
