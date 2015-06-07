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

const (
	URL_SCORES         = "http://kaiqiu.cc/home/space.php?searchmember=%s&province=&do=score&sex=&version=%s&bg=&score=&eventtype=%s&asso=&age=&searchscorefrom=&searchscoreto="
	URL_SPACE          = "http://kaiqiu.cc/home/%d"
	URL_HONORS         = "http://kaiqiu.cc/home/space.php?do=honor&username=%s"
	URL_DEFAULT_AVATAR = "http://116.255.247.74/ucenter/images/noavatar_big.gif"

	F_EVENTTYPE_PLAYER = "0"   //业余
	F_EVENTTYPE_PRO    = "1"   //专业
	F_VERSION_NOW      = "now" //即时
	F_VERSION_V1       = "v1"  //镜像
)

type PlayerScore struct {
	Spaceid      int
	Mname        string
	Avatar       string
	Cardno       string
	Siteorder    int
	Sex          string
	Region       string
	Name         string
	Score        int
	Level        string
	Score_year   int
	Score_high   int //历史最高
	Score_mirror int
	Mtype        uint8 //0 业余 1 专业
	Honors       []string
	matched      bool //是否完全匹配比赛用名
}

type SpaceInfo struct {
	Spaceid int
	Avatar  string
	Cardno  string
}

func NewPlayerScore() PlayerScore {
	return PlayerScore{0, "", URL_DEFAULT_AVATAR, "", 0, "", "", "", 0, "", 0, 0, 0, 0, []string{}, false}
}

// func (p PlayerScore) func getPlayerScore() {
// 	return
// }

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

func (ps *PlayerScores) ExecFetch(kw string, exact int, ver, etype string) (nFind int, bMatched bool) {
	P("exec find: ", kw, exact, ver, etype)
	url := fmt.Sprintf(URL_SCORES, url.QueryEscape(kw), ver, etype)
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
	nFind = doc.Find(".scoretab tr").Length() - 1

	findExact := false

	doc.Find(".scoretab tr").Each(func(i int, s *goquery.Selection) {

		if i == 0 {
			return
		}

		mname_raw, _ := s.Find("td").Eq(2).Html()

		reg := regexp.MustCompile(`space\-(\d+)\.html`)

		spaceida := reg.FindAllStringSubmatch(mname_raw, -1)
		spaceid, err := strconv.Atoi(spaceida[0][1])

		mnameFind := strings.TrimSpace(s.Find("td").Eq(2).Text())

		if err != nil {
			panic(err)
			return
		}

		pf, ok := ps.GetScore(spaceid)
		if !ok {
			// pf = PlayerScore{}
			pf = NewPlayerScore()
		}

		pf.Mname = mnameFind
		if pf.Mname == kw {
			pf.matched = true
			bMatched = true
		}else{
			if exact == 1 || findExact {
				findExact = true
				return
			}
		}

		pf.Spaceid = spaceid
		P(spaceid)

		pf.Siteorder, _ = strconv.Atoi(strings.TrimSpace(s.Find("td").Eq(1).Text()))
		pf.Sex = strings.TrimSpace(s.Find("td").Eq(3).Text())
		pf.Region = strings.TrimSpace(s.Find("td").Eq(4).Text())
		pf.Name = strings.TrimSpace(s.Find("td").Eq(5).Text())
		pf.Level = strings.TrimSpace(s.Find("td").Eq(7).Text())

		switch ver {
		case F_VERSION_V1:
			pf.Score_mirror, _ = strconv.Atoi(strings.TrimSpace(s.Find("td").Eq(6).Text()))
			pf.Score_year, _ = strconv.Atoi(strings.TrimSpace(s.Find("td").Eq(8).Text()))
		case F_VERSION_NOW:
			pf.Score, _ = strconv.Atoi(strings.TrimSpace(s.Find("td").Eq(6).Text()))
			pf.Score_high, _ = strconv.Atoi(strings.TrimSpace(s.Find("td").Eq(8).Text()))
		}

		ps.SetScore(spaceid, pf)
		// ps.scores[spaceid] = pf

		// pfs[spaceid] = pf

		// fmt.Printf("%v\n", pf)

		// PlayerFinds[]

	})

	P("find over ", ver, etype)
	// fmt.Printf("%v\n", pfs)

	//*/
	return
}

func (this *PlayerScores) GetSpaceInfo(spaceid int) (s SpaceInfo) {
	url := fmt.Sprintf(URL_SPACE, spaceid)
	P("get spaceinfo", url)
	// url := "http://kaiqiu.cc/home/space.php?searchmember=%E7%A7%A6%E6%98%8E&province=&do=score&sex=&version=v1&bg=&score=&eventtype=0&asso=&age=&searchscorefrom=&searchscoreto="
	c := 0
GOQUERYSTART:
	doc, err := goquery.NewDocument(url)
	if err != nil {
		if c < 3 {
			time.Sleep(time.Millisecond * 100)
			c++
			goto GOQUERYSTART
		}else{
			panic(err)
		}

		return
		// log.Fatal("goquery error:" , err)
	}

	s = SpaceInfo{Spaceid: spaceid, Avatar: URL_DEFAULT_AVATAR, Cardno: ""}

	avatar, exists := doc.Find("#space_avatar img").First().Attr("src")
	if exists == false {
		P("not exist")
		avatar, exists = doc.Find("td.image img").First().Attr("src")
		if exists == false {
			P("not exist")
			return
		}
	}

	// P(avatar)
	s.Avatar = GetRealAvatar(avatar)

	return s
}

/**
$title = pq($row)->find('td:eq(0)')->text();
           $title_h = pq($row)->find('td:eq(1) img')->attr('title');

           if($title && $title_h){
               $honors[] = $title.'['.$title_h.']';
           }else{
               break;
           }

*/
func (this *PlayerScores) GetHonors(mname string) (honors []string) {
	url := fmt.Sprintf(URL_HONORS, url.QueryEscape(mname))

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

	honors = []string{}

	doc.Find("table.spacegame tr").Each(func(i int, s *goquery.Selection) {

		title := s.Find("td").Eq(0).Text()
		title_h, exists := s.Find("img").Eq(0).Attr("title")
		if exists == false {
			return
		}

		// P(title, title_h)

		honors = append(honors, fmt.Sprintf("%s[%s]", title, title_h))
	})

	return
}
