package main

import (
	"Mahjong/src/internal"
	"goutils/choke"
	"log"
)

/**
* @Author: Jam Wong
* @Date: 2020/6/3 2:30 下午
 */

func main() {
	game := internal.NewGame()
	game.Deal()

	log.Print("hello <Mahjong>")
	choke.Choke()
}
