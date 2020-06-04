package internal

/**
* @Author: Jam Wong
* @Date: 2020/6/3
 */

import (
	"fmt"
	"log"
	"sort"
)

type ts []*Tile

type Player struct {
	IsBanker     bool // dealer/banker
	Name         string
	Index        int
	DiscardTiles ts // 出牌牌堆
	HoldTiles    ts // 手牌
}

func InitPlayer(index int, isBanker bool) *Player {
	return &Player{
		IsBanker:     isBanker,
		Name:         fmt.Sprintf("player_%d", index),
		Index:        index,
		DiscardTiles: make([]*Tile, 0),
		HoldTiles:    make([]*Tile, 0),
	}
}

// Draw
func (pl *Player) Draw(ch chan *Tile) *Tile {
	t := <-ch
	pl.HoldTiles = append(pl.HoldTiles, t)
	return t
}

// Discard
func (pl *Player) Discard(n int) {
	// Remove the tile at index n from pl.HoldTiles.
	if len(pl.HoldTiles) < 2 {
		fmt.Printf("%s\n", pl.Show())
		fmt.Println(pl.HoldTiles)
		log.Fatal("at least 1 tile!")
	}
	pl.HoldTiles[n] = pl.HoldTiles[len(pl.HoldTiles)-1]
	pl.HoldTiles[len(pl.HoldTiles)-1] = nil
	pl.HoldTiles = pl.HoldTiles[:len(pl.HoldTiles)-1]
	pl.SortTiles()
}

// Show
func (pl *Player) Show() string {
	var s string
	s += "|"
	for idx, i := range pl.HoldTiles {
		s += i.Print()
		s += "|"
		if idx%4 == 3 {
			s += "  |"
		}
	}
	return s
}

func (t ts) Show(from, to int) string {
	var s string
	s += "|"
	for _, i := range t[from:to] {
		s += i.Print()
		s += "|"
	}
	return s
}

// Order 使用快排
func (pl *Player) SortTiles() {
	sort.Sort(pl.HoldTiles)
}

func (pl *Player) CopySortTiles() ts {
	ttt := make(ts, 0)
	ttt = append(ttt, pl.HoldTiles...)
	sort.Sort(ttt)
	return ttt
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
