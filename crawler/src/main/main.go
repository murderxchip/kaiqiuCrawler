package main

import (
	// "fmt"
	// "log"
	// 
	"sync"
	"os"
	kq "../kqcrawler"
)






// var PlayerFinds = make(map[uint16]PlayerScore{})

func main() {
	// url := `http://116.255.247.74/ucenter/data/avatar/000/03/84/98_avatar_big.jpg`
	// kq.P(kq.GetRealAvatar(url), "test")
	// kq.P(kq.CheckResExists(url))
	// fmt.Println(r);
	// kq.P("test")

	wg := sync.WaitGroup{}
	ch := make(chan map[int]kq.PlayerScore, 4)

	kw := "秦明"

	Finder := func(kw, ver, etype string) {
		defer wg.Done()
		scores := kq.NewPlayerScores()
		scores.ExecFindList(kw, ver, etype)
		kq.P("get scores ",ver,etype)
		v, _ := scores.GetScores()
		ch <- v
		kq.P("get scores finished ",ver,etype)
	}

	wg.Add(4)
	go Finder(kw, kq.F_VERSION_V1, kq.F_EVENTTYPE_PLAYER)
	go Finder(kw, kq.F_VERSION_NOW, kq.F_EVENTTYPE_PLAYER)
	go Finder(kw, kq.F_VERSION_V1, kq.F_EVENTTYPE_PRO)
	go Finder(kw, kq.F_VERSION_NOW, kq.F_EVENTTYPE_PRO)
	wg.Wait()


//*
	for v := range ch {
		kq.P(v)



	}

//*/

	os.Exit(0)
	/*
	scores := kq.NewPlayerScores()
	scores.ExecFindList("秦明", kq.F_VERSION_V1, kq.F_EVENTTYPE_PRO)
	kq.P("get scores")
	kq.P(scores.GetScores())
	//*/
	// fmt.Println(scores.GetScores())

	//*
	
	//*/
}