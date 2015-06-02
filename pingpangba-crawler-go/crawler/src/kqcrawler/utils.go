package kqcrawler

import (
	"fmt"
	"net/http"
)

const env = "prod"

func CheckResExists(url string) bool {
	res, err := http.Head(url)
	if err != nil {
		//
		return false
	}

	switch res.StatusCode {
	case 304:
	case 200:
		return true
	default:
		return false
	}

	return false
	// fmt.Println(res.StatusCode)
}

func GetRealAvatar(url string) string {
	if CheckResExists(url) {
		return url
	}
	return "http://116.255.247.74/ucenter/images/noavatar_big.gif"
}

func Debug(v ...interface{}) {
	fmt.Printf("[debug]%v \n", v)
}

func P(v ...interface{}) {
	if env != "prod" {
		fmt.Printf("%v \n", v)
	}
}
