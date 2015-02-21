package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var toSay []TTS
var VALID_LANGS = [...]string{"af", "de", "en", "es", "fi", "fr", "is", "la", "no", "ru", "sv"}

// func start() {
// 	for range time.Tick(100 * time.Millisecond) {
// 		if len(toSay) > 0 {
// 			toSay[0].Play()
// 			toSay = toSay[1:]
// 		}
// 	}
// }

// func queue(tl, q string) {
// 	t := TTS{Tl: tl, Q: q}
// 	toSay = append(toSay, t)
// }

func handler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		rw.Header().Set("Content-Type", "application/json")
		q := req.FormValue("q")
		tl := req.FormValue("tl")

		if q != "" {
			for i := 0; i < len(VALID_LANGS); i++ {
				if tl == VALID_LANGS[i] {
					// go queue(tl, q)
					t := TTS{Tl: tl, Q: q}
					fmt.Fprint(rw, Response{Success: true, Message: "Here it is", Speech: t.Q, Link: t.Link()})
					return
				}
			}
			fmt.Fprint(rw, Response{Success: false, Message: "You know what you did... I need q and tl."})
		} else {
			fmt.Fprint(rw, Response{Success: false, Message: "You know what you did... I need q and tl."})
		}
	} else {
		http.NotFound(rw, req)
	}
}

func main() {
	port := flag.Int("port", 8000, "port number")
	flag.Parse()
	// go start()
	http.HandleFunc("/", handler)

	message := fmt.Sprintf("Starting server on :%v", *port)
	log.Println(message)
	address := fmt.Sprintf(":%v", *port)
	log.Fatal(http.ListenAndServe(address, nil))
}
