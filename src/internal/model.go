package internal

import (
	"fmt"
	"log"
)

/**
* @Author: Jam Wong
* @Date: 2020/6/3 2:53 下午
 */

type Player struct {
	IsBanker     bool // dealer/banker
	Name         string
	Index        int
	Win          bool
	DiscardTiles []*Tile // 出牌牌堆
	HoldTiles    []*Tile // 手牌
}

func InitPlayer(index int, isBanker bool) *Player {
	return &Player{
		IsBanker:     isBanker,
		Name:         fmt.Sprintf("player_%d", index),
		Index:        index,
		Win:          false,
		DiscardTiles: make([]*Tile, 0),
		HoldTiles:    make([]*Tile, 0),
	}
}

// Draw
func (pl *Player) Draw(ch chan *Tile) {
	pl.HoldTiles = append(pl.HoldTiles, <-ch)
}

// Discard
func (pl *Player) Discard(n int) {
	// Remove the tile at index n from pl.HoldTiles.
	if len(pl.HoldTiles) < 2 {
		log.Fatal("at least 1 tile.")
	}
	pl.HoldTiles[n] = pl.HoldTiles[len(pl.HoldTiles)-1]
	pl.HoldTiles[len(pl.HoldTiles)-1] = nil
	pl.HoldTiles = pl.HoldTiles[:len(pl.HoldTiles)-1]
}

// Show
func (pl *Player) Show() string {
	var s string
	s += "|"
	for _, i := range pl.HoldTiles {
		s += i.Print()
		s += "|"
	}
	return s
}

// Order 使用快排
//func (pl *Player) SortTiles()  {
//	newHold := make([]*Tile, 0)
//	for _, i := range pl.HoldTiles {
//		// 先排 type
//
//	}
//
//
//}
