package kqcrawler

import (
	// "log"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
	// "sort"
	"github.com/PuerkitoBio/goquery"
	// "errors"
)

type PlayerScore struct {
	spaceid      int
	mname        string
	avatar       string
	siteorder    int
	sex          string
	region       string
	name         string
	score        int
	level        string
	score_year   int
	score_high   int //历史最高
	score_mirror int
	mtype        uint8 //0 业余 1 专业
	matched      bool  //是否完全匹配比赛用名
}

func NewPlayerScore() PlayerScore {
	return PlayerScore{0, "", "http://116.255.247.74/ucenter/images/noavatar_big.gif", 0, "", "", "", 0, "", 0, 0, 0, 0, false}
}

// func (p PlayerScore) func getPlayerScore() {
// 	return
// }

const (
	SCORES_URL = "http://kaiqiu.cc/home/space.php?searchmember=%s&province=&do=score&sex=&version=%s&bg=&score=&eventtype=%s&asso=&age=&searchscorefrom=&searchscoreto="

	F_EVENTTYPE_PLAYER = "0"   //业余
	F_EVENTTYPE_PRO    = "1"   //专业
	F_VERSION_NOW      = "now" //即时
	F_VERSION_V1       = "v1"  //镜像
)

type PlayerScores struct {
	scores map[int]PlayerScore
	mux    sync.Mutex
}

func NewPlayerScores() *PlayerScores {
	return &PlayerScores{scores: make(map[int]PlayerScore)}
}

func (ps *PlayerScores) Count() int {
	return len(ps.scores)
}

func (ps *PlayerScores) GetScores() map[int]PlayerScore {
	return ps.scores
}

func (ps *PlayerScores) GetScore(spaceid int) (p PlayerScore, ok bool) {
	p, ok = ps.scores[spaceid]
	return
}

func (ps *PlayerScores) SetScore(spaceid int, score PlayerScore) {
	ps.mux.Lock()
	defer ps.mux.Unlock()
	ps.scores[spaceid] = score
}

func (ps *PlayerScores) ExecFindList(kw string, ver, etype string) {
	P("exec find: ", kw, ver, etype)
	url := fmt.Sprintf(SCORES_URL, url.QueryEscape(kw), ver, etype)
	P(url)
	// url := "http://kaiqiu.cc/home/space.php?searchmember=%E7%A7%A6%E6%98%8E&province=&do=score&sex=&version=v1&bg=&score=&eventtype=0&asso=&age=&searchscorefrom=&searchscoreto="
	c := 0
GOQUERYSTART:
	doc, err := goquery.NewDocument(url)
	if err != nil {
		if c < 3 {
			time.Sleep(time.Millisecond * 100)
			c++
			goto GOQUERYSTART
		}

		return
		// log.Fatal("goquery error:" , err)
	}
	//*
	// pfs := PlayerScores{}

	doc.Find(".scoretab tr").Each(func(i int, s *goquery.Selection) {

		if i == 0 {
			return
		}

		mname_raw, _ := s.Find("td").Eq(2).Html()

		reg := regexp.MustCompile(`space\-(\d+)\.html`)

		spaceida := reg.FindAllStringSubmatch(mname_raw, -1)
		spaceid, err := strconv.Atoi(spaceida[0][1])

		if err != nil {
			// panic(err)
			return
		}

		pf, ok := ps.GetScore(spaceid)
		if !ok {
			// pf = PlayerScore{}
			pf = NewPlayerScore()
		}

		pf.mname = strings.TrimSpace(s.Find("td").Eq(2).Text())
		if pf.mname == kw {
			pf.matched = true
		}

		pf.spaceid = spaceid
		P(spaceid)

		pf.siteorder, _ = strconv.Atoi(strings.TrimSpace(s.Find("td").Eq(1).Text()))
		pf.sex = strings.TrimSpace(s.Find("td").Eq(3).Text())
		pf.region = strings.TrimSpace(s.Find("td").Eq(4).Text())
		pf.name = strings.TrimSpace(s.Find("td").Eq(5).Text())
		pf.level = strings.TrimSpace(s.Find("td").Eq(7).Text())

		switch ver {
		case F_VERSION_V1:
			pf.score_mirror, _ = strconv.Atoi(strings.TrimSpace(s.Find("td").Eq(6).Text()))
			pf.score_year, _ = strconv.Atoi(strings.TrimSpace(s.Find("td").Eq(8).Text()))
		case F_VERSION_NOW:
			pf.score, _ = strconv.Atoi(strings.TrimSpace(s.Find("td").Eq(6).Text()))
			pf.score_high, _ = strconv.Atoi(strings.TrimSpace(s.Find("td").Eq(8).Text()))
		}

		ps.SetScore(spaceid, pf)
		// ps.scores[spaceid] = pf

		// pfs[spaceid] = pf

		fmt.Printf("%v\n", pf)

		// PlayerFinds[]
	})

	P("find over ", ver, etype)
	// fmt.Printf("%v\n", pfs)

	//*/
}

//****** map sorter
type PlayerScoreSorter []PlayerScore

func NewPlayerScoreSorter(m map[int]PlayerScore) PlayerScoreSorter {
	ms := make(PlayerScoreSorter, 0, len(m))
	for _, v := range m {
		ms = append(ms, v)
	}

	return ms
}

func (ms PlayerScoreSorter) Len() int {
	return len(ms)
}

func (ms PlayerScoreSorter) Less(i, j int) bool {
	if ms[i].matched {
		return true
	}

	return ms[i].score > ms[j].score
}

func (ms PlayerScoreSorter) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}
