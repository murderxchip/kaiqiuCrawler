package kqcrawler

import (
	"sort"
	"time"
)

type Crawler struct {
	Scores []PlayerScore
}

func (this *Crawler) FetchUserScores(kw string, limit int) (scores []PlayerScore) {

	ps := NewPlayerScores()

	var n,m int
	var matched bool
	n=0
	m = 0
	matched = false

	n, matched = ps.ExecFetch(kw, limit, F_VERSION_V1, F_EVENTTYPE_PLAYER)
	P(n,m)
	if n == 1 {
		ps.ExecFetch(kw, limit, F_VERSION_NOW, F_EVENTTYPE_PLAYER)
	}

	if !matched || n == 0 {
		m, matched = ps.ExecFetch(kw, limit, F_VERSION_V1, F_EVENTTYPE_PRO)
		// P(m, matched)
		if m == 1 && n == 0 && matched {
			ps.ExecFetch(kw, limit, F_VERSION_NOW, F_EVENTTYPE_PRO)
		}
	}
	// var scores1 map[int]PlayerScore
	scores1 := ps.GetScores()
	// P(scores1)
	ms := NewPlayerScoreSorter(scores1)
	sort.Sort(ms)

	P(ms)
	// P(scores1)
	// var scores_1 []PlayerScore
	if len(ms) > limit {
		this.Scores = ms[:limit]
	} else {
		this.Scores = ms
	}

	this.FetchSpaceInfo()
	this.FetchHonors()

	scores = this.Scores

	return
}

func (this *Crawler) FetchSpaceInfo() {
	if len(this.Scores)== 0 {
		return
	}
	// ps := NewPlayerScores()
	ch := make(chan SpaceInfo, len(this.Scores))
	stop := make(chan int)

	for _, v := range this.Scores {
		// P(v)
		go func(spaceid int) {
			ps := new(PlayerScores)
			ch <- ps.GetSpaceInfo(spaceid)
		}(v.Spaceid)
	}

	wait := 0

	for {
		select {
		case <-stop:
			P("stoped")
			return
		case si := <-ch:
			P(si)
			for i := 0; i < len(this.Scores); i++ {
				// P()
				if si.Spaceid == this.Scores[i].Spaceid {
					this.Scores[i].Avatar = si.Avatar
					break
				}
			}
			wait++
			P("wait:", wait, len(this.Scores))
			if wait >= len(this.Scores) {
				P("trigger stop")
				go func() { stop <- 1 }()
			}
		case <-time.After(time.Second * 15):
			return
		}
	}
}

func (this *Crawler) FetchHonors() {
	if len(this.Scores) == 1 {
		ps := new(PlayerScores)
		this.Scores[0].Honors = ps.GetHonors(this.Scores[0].Mname)
	}
	return
}
