package internal

import (
	"bufio"
	"fmt"
	"github.com/wzyonggege/goutils/convert"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

/********************
* @Author: Jam Wong *
* @Date: 2020/6/3   *
 ********************/

type Game struct {
	Tiles       chan *Tile
	Players     []*Player
	DealerIndex int
}

func NewGame() *Game {
	tiles := shuffle()

	g := &Game{
		Players:     make([]*Player, 0),
		Tiles:       make(chan *Tile, len(tiles)),
		DealerIndex: 0,
	}

	for _, i := range tiles {
		g.Tiles <- i
	}

	// TODO Dice 骰子确定index
	fmt.Println("init players")
	// init player
	for i := 0; i < 4; i++ {
		p := InitPlayer(i, g.DealerIndex == 0)
		g.Players = append(g.Players, p)
		fmt.Printf("player: %s\n", p.Name)
	}

	return g
}

// shuffle 洗牌
// TODO shuffle 洗牌 Fisher-Yates 高纳德置乱算法
// 从1~8中随机抽取一个数，例如随机数是3，那么交换第8位和第三位的数字。
// 此时数组顺序为12456783，重复第一步，从1~7中抽取一个数字，假设为4，那么交换第7位和第4位的数字
// 依次类推，直到第一个位置也被替代。
func shuffle() []*Tile {
	tiles := AllTiles()
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(tiles), func(i, j int) {
		tiles[i], tiles[j] = tiles[j], tiles[i]
	})

	return tiles
}

func (game *Game) DealNToPlayer(index int, n int) {
	pl := game.Players[index]
	if n < 1 || n > 4 {
		log.Fatal("n require [1, 4]")
	}
	for i := 1; i <= n; i++ {
		pl.DealTo(game.Tiles)
	}
}

func (game *Game) genOptions() {

}

// Deal 发牌
func (game *Game) Deal() {
	fmt.Println("start deal")

	// 三轮：一人四张， 从index 开始
	for i := 0; i < 4; i++ {
		for _, j := range game.Players {
			if i == 3 {
				game.DealNToPlayer(j.Index, 1)
			} else {
				game.DealNToPlayer(j.Index, 4)
			}
		}
	}

	// debug
	for _, j := range game.Players {
		j.SortTiles()
		fmt.Printf("%s: %s\n\n", j.Name, j.Show())
	}

	// 庄 1 张
	_, isWin := game.Players[game.DealerIndex].Draw(game.Tiles)
	if isWin {
		fmt.Printf("winwin %s\n", game.Players[game.DealerIndex].Show())
		return
	}
	fmt.Printf("round 1: %s \n", game.Players[game.DealerIndex].Show())

	discardIndex := readIndex()
	_dis := game.Players[game.DealerIndex].HoldTiles[discardIndex]
	game.Players[game.DealerIndex].Discard(discardIndex)
	fmt.Printf("%s ====> %s\n", game.Players[game.DealerIndex].Show(), _dis.Print())

	i := 1
	for {
		for u := 0; u < 4; u++ {
			if u == game.DealerIndex {
				if i == 1 {
					continue
				}
				fmt.Println("==========================================================")
				fmt.Printf("round %d: %s \n", i, game.Players[game.DealerIndex].Show())

				// draw
				draw, isWin := game.Players[game.DealerIndex].Draw(game.Tiles)

				fmt.Printf("<<<======= %s \n", draw.Print())

				if isWin {
					fmt.Printf("winwin %s\n", game.Players[game.DealerIndex].Show())
					return
				}

				discardIndex := readIndex()
				_dis := game.Players[game.DealerIndex].HoldTiles[discardIndex]
				game.Players[game.DealerIndex].Discard(discardIndex)
				fmt.Printf("%s ====> %s\n", game.Players[game.DealerIndex].Show(), _dis.Print())
			} else {
				if game.BotRound(u) {
					fmt.Printf("Bot winwin %s\n", game.Players[u].Show())
					return
				}
			}
		}
		i++

	}
}

func readIndex() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	i, _ := convert.StringToInt(text)
	return i - 1
}

func readOptions() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	i, _ := convert.StringToInt(text)
	return i - 1
}

func (game *Game) BotRound(playerIndex int) bool {
	player := game.Players[playerIndex]
	// draw
	_, isWin := player.Draw(game.Tiles)
	// discard
	_dis := player.HoldTiles[13]
	player.Discard(13)
	fmt.Printf("%s ====> %s\n", player.Name, _dis.Print())
	return isWin
}
