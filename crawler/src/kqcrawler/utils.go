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

func p(v interface{}){
	fmt.Println(v)
}