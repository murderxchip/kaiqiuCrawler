package main

import (
	// "fmt"
	// "log"
	// 

	kq "kqcrawler"
)






// var PlayerFinds = make(map[uint16]PlayerScore{})

func main() {
	// url := `http://116.255.247.74/ucenter/data/avatar/000/03/84/98_avatar_big.jpg`
	// kq.P(kq.CheckResExists(url))
	// fmt.Println(r);
	// kq.P("test")
	scores := kq.NewPlayerScores()
	scores.ExecFind("秦明")
}
