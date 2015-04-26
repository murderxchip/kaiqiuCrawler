package kqcrawler

import (
	"log"
	"fmt"
	"strings"
	"regexp"
	"strconv"
	nu "net/url"
	"github.com/PuerkitoBio/goquery"
)

type PlayerScore struct {
	mname string
	siteorder uint32
	sex uint8
	region string
	name string
	score uint16
	level string
	score_year uint16
	score_high uint16
	mtype uint8 //0 业余 1 专业 
}

// func (p PlayerScore) func getPlayerScore() {
// 	return 
// }

const (
	SCORES_URL = "http://kaiqiu.cc/home/space.php?searchmember=%s&province=&do=score&sex=&version=v1&bg=&score=&eventtype=0&asso=&age=&searchscorefrom=&searchscoreto="
)

type PlayerScores struct {
	scores map[int]PlayerScore
}

func NewPlayerScores() *PlayerScores{
	ps := &PlayerScores{scores: make(map[int]PlayerScore) }
	return ps
}

func (ps *PlayerScores) GetScores() map[int]PlayerScore {
	return ps.scores
	// return map[int]PlayerScore{}
}

func (ps *PlayerScores) GetScore(spaceid int) PlayerScore {
	return ps.scores[spaceid]
}

func (ps *PlayerScores) SetScore(spaceid int, score PlayerScore) {
	ps.scores[spaceid] = score
}

func (ps *PlayerScores) ExecFind(kw string) {
	P("exec find: " + kw)
	url := fmt.Sprintf(SCORES_URL, nu.QueryEscape(kw))
	P(url)
	// url := "http://kaiqiu.cc/home/space.php?searchmember=%E7%A7%A6%E6%98%8E&province=&do=score&sex=&version=v1&bg=&score=&eventtype=0&asso=&age=&searchscorefrom=&searchscoreto="
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	//*
	// pfs := PlayerScores{}

	doc.Find(".scoretab tr").Each(func(i int, s *goquery.Selection) {
		
		if i == 0 {
			return
		}

		pf := PlayerScore{}

		pf.mname = strings.TrimSpace(s.Find("td").Eq(2).Text())
		mname_raw,_ := s.Find("td").Eq(2).Html()
		reg := regexp.MustCompile(`space\-(\d+)\.html`)

		spaceida := reg.FindAllStringSubmatch(mname_raw, -1)
		spaceid, err := strconv.Atoi(spaceida[0][1])

		if err != nil {
	        panic(err)
	    }

		P(spaceid)

		// pfs[spaceid] = pf

		// fmt.Printf("%v\n", pf.mname)

		// PlayerFinds[]
	})

	// fmt.Printf("%v\n", pfs)

	//*/
}
