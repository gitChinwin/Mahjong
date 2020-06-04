package internal

/**
* @Author: Jam Wong
* @Date: 2020/6/3
 */

func (pl *Player) Win() bool {
	cc := pl.CopySortTiles()

	// 牌数量 3n+2
	if cc.Len() < 2 || (cc.Len()-2)%3 != 0 {
		return false
	}

	for i := 0; i < len(cc)-1; i++ {
		if isPair(cc[i], cc[i+1]) {
			tmp := make([]*Tile, 0)

			tmp = append(tmp, cc[:i]...)
			tmp = append(tmp, cc[i+2:]...)
			if canWinLoop(tmp) {
				return true
			}
		}
	}

	return false
}

func canWinLoop(t []*Tile) bool {
	if len(t) == 0 {
		return true
	}
	//for _, i := range t {
	//	fmt.Printf("%s,", i.Print())
	//}

	//fmt.Printf("\n%s %s %s\n", t[0].Print(), t[1].Print(), t[2].Print())
	if isSequence(t[0], t[1], t[2]) {
		//fmt.Println("tttttt")
		return canWinLoop(t[3:])
	}

	if isTriplet(t[0], t[1], t[2]) {
		//fmt.Println("sssss")
		return canWinLoop(t[3:])
	}

	if len(t) >= 6 {
		if isTwoSequence(t[0], t[1], t[2], t[3], t[4], t[5]) {
			return canWinLoop(t[6:])
		}
		if isTwoSequence2(t[0], t[1], t[2], t[3], t[4], t[5]) {
			return canWinLoop(t[6:])
		}
	}

	return false
}

// 将牌(雀头、对子)
func isPair(a, b *Tile) bool {
	return a.Type == b.Type && a.Rank == b.Rank
}

// 顺子	Sequence
func isSequence(a, b, c *Tile) bool {
	if !(a.Type == b.Type && a.Type == c.Type) {
		return false
	}
	if !(a.Rank == b.Rank-1 && b.Rank == c.Rank-1) {
		return false
	}
	return true
}

// 顺子	Sequence
func isTwoSequence(a, a2, b, b2, c, c2 *Tile) bool {
	return isSequence(a, b, c) && isSequence(a2, b2, c2)
}

func isTwoSequence2(a, b, b2, c, c2, d *Tile) bool {
	return isSequence(a, b, c) && isSequence(b2, c2, d)
}

// 刻子	Triplet
func isTriplet(a, b, c *Tile) bool {
	return (a.Type == b.Type && a.Type == c.Type) && (a.Rank == b.Rank && a.Rank == c.Rank)
}
