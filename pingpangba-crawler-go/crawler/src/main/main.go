package main

import (
	"fmt"
	// "html"
	"log"
	// "sync"
	// "os"
	kq "kqcrawler"
	"encoding/json"
	// "os"
	// "sort"
	"net/http"
	"strconv"
)

func main() {
	
	http.HandleFunc("/scores", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		r.ParseForm()

		// fmt.Fprintf(w, "kw", r.Form["kw"])
		if r.Method == "GET" {
			// fmt.Println("kw", r.Form["kw"])
			kw := r.Form["kw"][0]
			n,_ := strconv.Atoi(r.Form["n"][0])

			kq.Debug(kw, n)

			crawler := new(kq.Crawler)
			scores := crawler.FetchUserScores(kw, n)

			jsonScores, err := json.Marshal(scores)
			if err != nil {
			    kq.Debug("error:", err)
			}

			fmt.Fprintf(w, "%s", jsonScores)
			// kq.P(scores)
		}
	})

	log.Fatal(http.ListenAndServe(":9001", nil))
}
