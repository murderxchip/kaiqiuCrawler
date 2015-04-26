package kqcrawler

import (
	"fmt"
	"log"
	"strings"
	"regexp"
	"strconv"

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

type PlayerScores map[int]PlayerScore

// var pfs = PlayerScores{}

func Scores() PlayerScores {
	url := "http://kaiqiu.cc/home/space.php?searchmember=%E7%A7%A6%E6%98%8E&province=&do=score&sex=&version=v1&bg=&score=&eventtype=0&asso=&age=&searchscorefrom=&searchscoreto="
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	pfs := PlayerScores{}

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

		p(spaceid)

		pfs[spaceid] = pf

		// fmt.Printf("%v\n", pf.mname)

		// PlayerFinds[]
	})

	fmt.Printf("%v\n", pfs)

	return pfs
	
}
