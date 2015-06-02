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

	n, matched := ps.ExecFetch(kw, F_VERSION_V1, F_EVENTTYPE_PLAYER)
	if n > 0 {
		ps.ExecFetch(kw, F_VERSION_NOW, F_EVENTTYPE_PLAYER)
	}

	if !matched {
		n, matched = ps.ExecFetch(kw, F_VERSION_V1, F_EVENTTYPE_PRO)
		if n > 0 {
			ps.ExecFetch(kw, F_VERSION_NOW, F_EVENTTYPE_PRO)
		}
	}
	// var scores1 map[int]PlayerScore
	scores1 := ps.GetScores()

	ms := NewPlayerScoreSorter(scores1)
	sort.Sort(ms)

	// P(ms)
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
		case <-time.After(time.Second * 40):
			return
		}
	}
}

func (this *Crawler) FetchHonors() {
	if len(this.Scores) == 1 {
		ps := new(PlayerScores)
		this.Scores[0].Honors = ps.GetHonors(this.Scores[0].Mname)
	}
}
