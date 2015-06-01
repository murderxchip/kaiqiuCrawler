package main

import (
	// "fmt"
	// "log"
	//
	// "sync"
	// "os"
	kq "../kqcrawler"
	// "encoding/json"
	// "os"
	// "sort"
)

// var PlayerFinds = make(map[uint16]PlayerScore{})

func main() {
	// url := `http://116.255.247.74/ucenter/data/avatar/000/03/84/98_avatar_big.jpg`
	// kq.P(kq.GetRealAvatar(url), "test")
	// kq.P(kq.CheckResExists(url))
	// fmt.Println(r);
	// kq.P("test")

	/*
	   	wg := sync.WaitGroup{}
	   	ch := make(chan map[int]kq.PlayerScore, 4)

	   	// kw := "黑杰克"
	   	kw := "秦"

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

	   	// kq.P(<-ch)


	   	for i := 0;i<4;i++{
	   		m := <-ch //map[int]PlayerScore
	   		kq.P("m length:", len(m))


	   	}

	   	scores_merge := make(map[int]PlayerScore)

	   	for {
	   		select {

	   		}
	   	}
	   //*
	   	for v := range ch {
	   		kq.P(v)


	   	}

	   //*/

	// os.Exit(0)
	/*
		scores := kq.NewPlayerScores()
		scores.ExecFindList("秦明", kq.F_VERSION_V1, kq.F_EVENTTYPE_PRO)
		kq.P("get scores")
		kq.P(scores.GetScores())
		//*/
	// fmt.Println(scores.GetScores())

	//*

	/*
		kw := "贝壳"

		scores := kq.NewPlayerScores()
		scores.ExecFindList(kw, kq.F_VERSION_V1, kq.F_EVENTTYPE_PLAYER)
		if scores.Count() > 0 {
			scores.ExecFindList(kw, kq.F_VERSION_NOW, kq.F_EVENTTYPE_PLAYER)
		}

		// scores.ExecFindList(kw, kq.F_VERSION_V1, kq.F_EVENTTYPE_PRO)
		// scores.ExecFindList(kw, kq.F_VERSION_NOW, kq.F_EVENTTYPE_PRO)

		kq.P(scores.Count())
		scores1 := scores.GetScores()

		ms := kq.NewPlayerScoreSorter(scores1)
		sort.Sort(ms)

		var ms1 []kq.PlayerScore

		if len(ms) > 8 {
			ms1 = ms[0:9]
		} else {
			ms1 = ms
		}
		// ms1 := ms[0:9]

		//todo get avatar

		b, err := json.Marshal(ms1)
		if err != nil {
			kq.P("error:", err)
		}

		// kq.P(ms1)

		os.Stdout.Write(b)

		// kq.P(ms1)
		// for _, item := range ms {
		//     kq.P(item)
		// }

		//*/

	//* getspaceinfo

	//11328 6896
	scores := kq.NewPlayerScores()
	// var spaceinfo kq.SpaceInfo
	// spaceinfo = scores.GetSpaceInfo(11328)


	// fmt.Println(spaceinfo.Avatar)

	scores.GetHonors("创意")

	//*/

	//*


	//*/

}
