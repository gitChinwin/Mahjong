package internal

/**
* @Author: Jam Wong
* @Date: 2020/6/3
 */

import (
	"fmt"
	"log"
	"os"
	"sort"
)

type Player struct {
	IsBanker     bool // dealer/banker
	Name         string
	Index        int
	DiscardTiles ts // 出牌牌堆
	HoldTiles    ts // 手牌
	Melds        ts // 组牌
}

func InitPlayer(index int, isBanker bool) *Player {
	return &Player{
		IsBanker:     isBanker,
		Name:         fmt.Sprintf("player_%d", index),
		Index:        index,
		DiscardTiles: make([]*Tile, 0),
		HoldTiles:    make([]*Tile, 0),
		Melds:        make([]*Tile, 0),
	}
}

// Draw
func (pl *Player) Draw(ch chan *Tile) *Tile {
	t := <-ch
	pl.HoldTiles = append(pl.HoldTiles, t)
	if len(ch) == 0 {
		fmt.Println("over!!!!")
		os.Exit(0)
	}
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
	pl.DiscardTiles = append(pl.DiscardTiles, pl.HoldTiles[n])

	pl.HoldTiles[n] = pl.HoldTiles[len(pl.HoldTiles)-1]
	pl.HoldTiles[len(pl.HoldTiles)-1] = nil
	pl.HoldTiles = pl.HoldTiles[:len(pl.HoldTiles)-1]
	pl.SortTiles()
}

// Pong 碰
func (pl *Player) Pong() {

}

// Kong 杠
func (pl *Player) Kong() {

}

// Chow 吃
func (pl *Player) Chow() {

}

// Exposed 明杠
func (pl *Player) Exposed() {

}

// Concealed Kong 暗杠
func (pl *Player) ConcealedKong() {

}

// Show
func (pl *Player) Show() string {
	return pl.HoldTiles.Show(0, len(pl.HoldTiles))
}

// Order 使用快排
func (pl *Player) SortTiles() {
	sort.Sort(pl.HoldTiles)
}

func (pl *Player) CopySortTiles() ts {
	return pl.HoldTiles.SortTiles()
}
