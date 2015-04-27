package kqcrawler

import (
	"net/http"
	"fmt"
	)

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
	return "http://116.255.247.74/ucenter/images/noavatar_big.gif";
}

func P(v ...interface{}){
	fmt.Printf("[debug print] %v \n", v)
}